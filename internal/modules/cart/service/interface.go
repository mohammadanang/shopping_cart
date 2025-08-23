package service

import (
	"context"

	"github.com/mohammadanang/shopping-cart/pkg/api"
)

type Service interface {
	Add(ctx context.Context, payload api.AddCartJSONRequestBody) (*api.CartSuccessResponse, error)
	List(ctx context.Context) (*api.CartListSuccessResponse, error)
	Remove(ctx context.Context, id int64) (*api.CartSuccessResponse, error)
}
