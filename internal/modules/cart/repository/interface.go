package repository

import (
	"context"

	"github.com/mohammadanang/shopping-cart/db/dbgen"
	"github.com/mohammadanang/shopping-cart/internal/modules/cart/domain"
)

type Repository interface {
	TxCreateOrUpdate(ctx context.Context, payload domain.AddRequest) (*domain.AddResponse, error)
	ShowOrder(ctx context.Context, id int64) (*dbgen.Order, error)
	List(ctx context.Context, orderId int64) ([]*dbgen.Cart, error)
}
