package repository

import (
	"context"

	"github.com/mohammadanang/shopping-cart/db/dbgen"
	"github.com/mohammadanang/shopping-cart/internal/modules/cart/domain"
	"github.com/mohammadanang/shopping-cart/pkg/util"
)

type CartRepository struct {
	store dbgen.Store
}

func NewCartRepository(store dbgen.Store) Repository {
	return &CartRepository{
		store: store,
	}
}

func (r *CartRepository) TxCreateOrUpdate(ctx context.Context, payload domain.AddRequest) (*domain.AddResponse, error) {
	var result domain.AddResponse
	err := r.store.ExecTx(ctx, func(q *dbgen.Queries) error {
		var orderData *dbgen.Order
		if payload.OrderId != nil {
			getOrder, err := q.GetOrder(ctx, *payload.OrderId)
			if err != nil {
				return err
			}

			orderData = getOrder
			result.Type = "edit"
		} else {
			orderNumber := util.GenerateKeyNumber("ORD")
			addedOrder, err := q.AddOrder(ctx, &dbgen.AddOrderParams{
				Status:      "pending",
				Total:       0,
				Discount:    0,
				OrderNumber: orderNumber,
				Buyer:       payload.Buyer,
			})
			if err != nil {
				return err
			}

			orderData = addedOrder
			result.Type = "add"
		}

		var cartSlices []domain.Cart
		var orderTotal float64 = 0
		err := q.RemoveCartsByOrder(ctx, orderData.ID)
		if err != nil {
			return err
		}

		for _, item := range payload.Data {
			addedCrt, err := q.AddCart(ctx, &dbgen.AddCartParams{
				ProductName: item.ProductName,
				Qty:         item.Qty,
				Price:       item.Price,
				OrderID:     orderData.ID,
			})
			if err != nil {
				return err
			}

			cart := domain.Cart{
				Id:          addedCrt.ID,
				ProductName: addedCrt.ProductName,
				Qty:         addedCrt.Qty,
				Price:       addedCrt.Price,
				OrderId:     addedCrt.OrderID,
				CreatedAt:   &addedCrt.CreatedAt,
				UpdatedAt:   &addedCrt.UpdatedAt,
			}

			cartSlices = append(cartSlices, cart)
			cartTotal := cart.Price * float64(cart.Qty)
			orderTotal += cartTotal
		}

		editedOrd, err := q.EditOrder(ctx, &dbgen.EditOrderParams{
			ID:       orderData.ID,
			Discount: orderData.Discount,
			Status:   orderData.Status,
			Total:    orderTotal,
		})
		if err != nil {
			return err
		}

		result.OrderId = editedOrd.ID
		result.OrderNumber = editedOrd.OrderNumber
		result.Items = cartSlices
		result.Buyer = editedOrd.Buyer

		return nil
	})

	return &result, err
}

func (r *CartRepository) ShowOrder(ctx context.Context, id int64) (*dbgen.Order, error) {
	order, err := r.store.GetOrder(ctx, id)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (r *CartRepository) List(ctx context.Context, orderId int64) ([]*dbgen.Cart, error) {
	carts, err := r.store.ListCarts(ctx, orderId)
	if err != nil {
		return nil, err
	}

	return carts, nil
}
