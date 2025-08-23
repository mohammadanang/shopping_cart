package repository

import (
	"context"

	"github.com/mohammadanang/shopping-cart/db/dbgen"
	"github.com/mohammadanang/shopping-cart/internal/modules/order/domain"
)

type Repository interface {
	Create(ctx context.Context, params dbgen.AddOrderParams) (*dbgen.Order, error)
	Update(ctx context.Context, params dbgen.EditOrderParams) (*dbgen.Order, error)
	Show(ctx context.Context, id int64) (*dbgen.Order, error)
	Paginate(ctx context.Context, params domain.PaginateRequest) ([]*dbgen.Order, error)
	Count(ctx context.Context, orderNumber *string) (int64, error)
}
