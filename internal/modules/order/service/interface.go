package service

import (
	"context"

	"github.com/mohammadanang/shopping-cart/internal/modules/order/domain"
)

type Service interface {
	// Add(ctx context.Context, payload api.AddOrderJSONRequestBody) (*api.OrderSuccessResponse, error)
	// Edit(ctx context.Context, id api.IdParam, payload api.EditOrderJSONRequestBody) (*api.OrderSuccessResponse, error)
	// Show(ctx context.Context, id api.IdParam) (*api.OrderSuccessResponse, error)
	Paginate(ctx context.Context, payload *domain.PaginateRequest) (*domain.PaginateResponse, error)
}
