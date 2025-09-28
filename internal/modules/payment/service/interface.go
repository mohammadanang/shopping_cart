package service

import (
	"context"

	"github.com/mohammadanang/shopping-cart/internal/modules/payment/domain"
)

type Service interface {
	Paginate(ctx context.Context, payload *domain.PaginateRequest) (*domain.PaginateResponse, error)
	CompletePayment(ctx context.Context, orderId int64, payload domain.CompleteRequest) (*domain.CompleteResponse, error)
	WebHookOfXendit(ctx context.Context, payload domain.XenditWebhook, cbToken string) (*domain.PaymentWithOrder, error)
}
