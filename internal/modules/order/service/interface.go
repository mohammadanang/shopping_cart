package service

import (
	"context"

	"github.com/mohammadanang/shopping-cart/internal/modules/order/domain"
)

type Service interface {
	EditOrderAndPayment(ctx context.Context, orderId int64, payload domain.UpdateRequest) (*domain.UpdateResponse, error)
	Show(ctx context.Context, id int64) (*domain.ShowResponse, error)
	Paginate(ctx context.Context, payload *domain.PaginateRequest) (*domain.PaginateResponse, error)
}
