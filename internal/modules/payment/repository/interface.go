package repository

import (
	"context"

	"github.com/mohammadanang/shopping-cart/db/dbgen"
	"github.com/mohammadanang/shopping-cart/internal/modules/payment/domain"
)

type Repository interface {
	Paginate(ctx context.Context, payload domain.PaginateParam) ([]*dbgen.Payment, error)
	Count(ctx context.Context, paymentNumber *string) (int64, error)
	ShowOrder(ctx context.Context, orderId int64) (*dbgen.Order, error)
	TxCompletePayment(ctx context.Context, orderId int64, payload domain.CompleteRequest) (*domain.CompleteResponse, error)
}
