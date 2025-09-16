package repository

import (
	"context"

	"github.com/mohammadanang/shopping-cart/db/dbgen"
	"github.com/mohammadanang/shopping-cart/internal/modules/order/domain"
	"github.com/mohammadanang/shopping-cart/pkg/util"
)

type OrderRepository struct {
	store dbgen.Store
}

func NewOrderRepository(store dbgen.Store) Repository {
	return &OrderRepository{
		store: store,
	}
}

func (r *OrderRepository) TxUpdateOrderAndPayment(ctx context.Context, orderId int64, payload domain.UpdateRequest) (*domain.UpdateResponse, error) {
	var result domain.UpdateResponse
	err := r.store.ExecTx(ctx, func(q *dbgen.Queries) error {
		var orderData dbgen.Order
		getOrder, err := q.GetOrder(ctx, orderId)
		if err != nil {
			return err
		}

		orderData = *getOrder
		updatePayload := dbgen.EditOrderParams{
			ID:       orderData.ID,
			Status:   orderData.Status,
			Discount: payload.Discount,
			Total:    orderData.Total,
		}
		updated, err := q.EditOrder(ctx, &updatePayload)
		if err != nil {
			return err
		}

		var paymentData dbgen.Payment
		getPayment, err := q.GetPaymentByOrder(ctx, orderId)
		if err != nil {
			paymentPayload := dbgen.AddPaymentParams{
				PaymentNumber: util.GenerateKeyNumber("PAY"),
				OrderID:       updated.ID,
				Method:        payload.PaymentMethod,
				Total:         updated.Total - updated.Discount,
			}
			newPayment, err := q.AddPayment(ctx, &paymentPayload)
			if err != nil {
				return err
			}

			paymentData = *newPayment
		} else {
			paymentData = *getPayment
		}

		result.Buyer = updated.Buyer
		result.OrderId = updated.ID
		result.Discount = updated.Discount
		result.OrderNumber = updated.OrderNumber
		result.PaymentNumber = paymentData.PaymentNumber
		result.Status = updated.Status
		result.TotalOrder = updated.Total
		result.TotalPaid = paymentData.Total

		return nil
	})

	return &result, err
}

func (r *OrderRepository) Show(ctx context.Context, id int64) (*dbgen.Order, error) {
	order, err := r.store.GetOrder(ctx, id)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (r *OrderRepository) Paginate(ctx context.Context, payload domain.PaginateParam) ([]*dbgen.Order, error) {
	if payload.OrderNumber != nil {
		payload := dbgen.PaginateOrdersWithParamsParams{
			Limit:       payload.Limit,
			Offset:      payload.Offset,
			OrderNumber: *payload.OrderNumber,
		}
		orders, err := r.store.PaginateOrdersWithParams(ctx, &payload)
		if err != nil {
			return nil, err
		}

		return orders, nil
	}

	params := dbgen.PaginateOrdersParams{
		Limit:  payload.Limit,
		Offset: payload.Offset,
	}
	orders, err := r.store.PaginateOrders(ctx, &params)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *OrderRepository) Count(ctx context.Context, orderNumber *string) (int64, error) {
	if orderNumber != nil {
		count, err := r.store.CountOrdersByOrderNumber(ctx, *orderNumber)
		if err != nil {
			return 0, err
		}

		return count, nil
	}

	count, err := r.store.CountOrders(ctx)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *OrderRepository) ListCart(ctx context.Context, orderId int64) ([]*dbgen.Cart, error) {
	carts, err := r.store.ListCarts(ctx, orderId)
	if err != nil {
		return nil, err
	}

	return carts, nil
}

func (r *OrderRepository) ShowPayment(ctx context.Context, id int64) (*dbgen.Payment, error) {
	payment, err := r.store.GetPaymentByOrder(ctx, id)
	if err != nil {
		return nil, err
	}

	return payment, nil
}
