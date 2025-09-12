package repository

import (
	"context"

	"github.com/mohammadanang/shopping-cart/db/dbgen"
	"github.com/mohammadanang/shopping-cart/internal/modules/order/domain"
)

type Repository interface {
	Show(ctx context.Context, id int64) (*dbgen.Order, error)
	Paginate(ctx context.Context, payload domain.PaginateParam) ([]*dbgen.Order, error)
	Count(ctx context.Context, orderNumber *string) (int64, error)
	TxUpdateOrderAndPayment(ctx context.Context, payload domain.UpdateRequest) (*domain.UpdateResponse, error)
}
