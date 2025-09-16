package repository

import (
	"context"

	"github.com/mohammadanang/shopping-cart/db/dbgen"
	"github.com/mohammadanang/shopping-cart/internal/modules/order/domain"
)

type OrderRepository struct {
	store dbgen.Store
}

func NewOrderRepository(store dbgen.Store) Repository {
	return &OrderRepository{
		store: store,
	}
}

func (r *OrderRepository) TxUpdateOrderAndPayment(ctx context.Context, payload domain.UpdateRequest) (*domain.UpdateResponse, error) {
	return nil, nil
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
