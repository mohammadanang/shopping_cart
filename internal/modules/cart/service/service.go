package service

import (
	"context"
	"log"

	"github.com/mohammadanang/shopping-cart/db/dbgen"
	"github.com/mohammadanang/shopping-cart/internal/modules/cart/domain"
	cartR "github.com/mohammadanang/shopping-cart/internal/modules/cart/repository"
	"github.com/mohammadanang/shopping-cart/pkg/util"
)

type CartService struct {
	store    dbgen.Store
	cartRepo cartR.Repository
}

func NewCartService(store dbgen.Store, cartRepo cartR.Repository) Service {
	return &CartService{
		cartRepo: cartRepo,
	}
}

func (s *CartService) AddAndOrEdit(ctx context.Context, payload domain.AddRequest) (*domain.AddResponse, error) {
	result := make(chan domain.Result)
	go func() {
		defer close(result)

		var cartResult domain.Result
		err := s.store.ExecTx(ctx, func(q *dbgen.Queries) error {
			var orderData *dbgen.Order
			if payload.OrderId != nil {
				getOrder, err := q.GetOrder(ctx, *payload.OrderId)
				if err != nil {
					return err
				}

				orderData = getOrder
			} else {
				orderNumber := util.GenerateKeyNumber("ORD")
				addedOrder, err := q.AddOrder(ctx, &dbgen.AddOrderParams{
					Status:      "pending",
					Total:       0,
					Discount:    0,
					OrderNumber: orderNumber,
				})
				if err != nil {
					return err
				}

				orderData = addedOrder
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
					Buyer:       item.Buyer,
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
					Buyer:       addedCrt.Buyer,
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

			response := domain.AddResponse{
				OrderId:     editedOrd.ID,
				OrderNumber: editedOrd.OrderNumber,
				Items:       cartSlices,
			}

			cartResult.Value = response
			result <- cartResult

			return nil
		})
		if err != nil {
			cartResult.Error = err
			result <- cartResult
			return
		}
	}()

	res := <-result
	if res.Error != nil {
		log.Println("transaction failed", res.Error.Error())
	}

	data := res.Value.(domain.AddResponse)

	return &data, nil
}

// func (s *CartService) List(ctx context.Context) (*api.CartListSuccessResponse, error) {
// 	items, err := s.repo.List(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// relate order data
// 	var carts []api.Cart
// 	for _, item := range items {
// 		carts = append(carts, api.Cart{
// 			Id:          item.ID,
// 			ProductName: item.ProductName,
// 			Qty:         item.Qty,
// 			Buyer:       item.Buyer,
// 			OrderId:     item.OrderID,
// 			Price:       int(item.Price),
// 			CreatedAt:   item.CreatedAt.Format("2006-01-02 15:04:05"),
// 			UpdatedAt:   item.UpdatedAt.Format("2006-01-02 15:04:05"),
// 		})
// 	}

// 	return &api.CartListSuccessResponse{
// 		Code:    200,
// 		Data:    carts,
// 		Message: "Cart items retrieved successfully",
// 		Status:  "success",
// 	}, nil
// }
