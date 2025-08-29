package service

import (
	"context"

	"github.com/mohammadanang/shopping-cart/internal/modules/cart/domain"
)

type Service interface {
	AddAndOrEdit(ctx context.Context, payload domain.AddRequest) (*domain.AddResponse, error)
	// List(ctx context.Context) ([]*domain.Cart, error)
	// Remove(ctx context.Context, id int64) (*domain.Cart, error)
}
