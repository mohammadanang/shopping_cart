package repository

import (
	"context"

	"github.com/mohammadanang/shopping-cart/db/dbgen"
	"github.com/mohammadanang/shopping-cart/internal/modules/order/domain"
)

type OrderRepository struct {
	q *dbgen.Queries
}

func NewOrderRepository(queries *dbgen.Queries) Repository {
	return &OrderRepository{
		q: queries,
	}
}

func (r *OrderRepository) Create(ctx context.Context, params dbgen.AddOrderParams) (*dbgen.Order, error) {
	order, err := r.q.AddOrder(ctx, &params)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (r *OrderRepository) Update(ctx context.Context, params dbgen.EditOrderParams) (*dbgen.Order, error) {
	order, err := r.q.EditOrder(ctx, &params)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (r *OrderRepository) Show(ctx context.Context, id int64) (*dbgen.Order, error) {
	order, err := r.q.GetOrder(ctx, id)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (r *OrderRepository) Paginate(ctx context.Context, params domain.PaginateRequest) ([]*dbgen.Order, error) {
	if params.OrderNumber != nil {
		payload := dbgen.PaginateOrdersWithParamsParams{
			Limit:       params.Limit,
			Offset:      params.Offset,
			OrderNumber: *params.OrderNumber,
		}
		orders, err := r.q.PaginateOrdersWithParams(ctx, &payload)
		if err != nil {
			return nil, err
		}

		return orders, nil
	}

	payload := dbgen.PaginateOrdersParams{
		Limit:  params.Limit,
		Offset: params.Offset,
	}
	orders, err := r.q.PaginateOrders(ctx, &payload)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *OrderRepository) Count(ctx context.Context, orderNumber *string) (int64, error) {
	if orderNumber != nil {
		count, err := r.q.CountOrdersByOrderNumber(ctx, *orderNumber)
		if err != nil {
			return 0, err
		}

		return count, nil
	}

	count, err := r.q.CountOrders(ctx)
	if err != nil {
		return 0, err
	}

	return count, nil
}
