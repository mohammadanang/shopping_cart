package service

import (
	"context"
	"log"
	"math"

	"github.com/mohammadanang/shopping-cart/internal/modules/payment/domain"
	"github.com/mohammadanang/shopping-cart/internal/modules/payment/repository"
	"github.com/mohammadanang/shopping-cart/pkg/wrapper"
)

type PaymentService struct {
	repo repository.Repository
}

func NewPaymentService(repo repository.Repository) Service {
	return &PaymentService{
		repo: repo,
	}
}

func (s *PaymentService) Paginate(ctx context.Context, payload *domain.PaginateRequest) (*domain.PaginateResponse, error) {
	page := int32(1)
	size := int32(10)
	if payload.Size != nil {
		size = *payload.Size
	}

	if payload.Page != nil {
		page = *payload.Page
	}

	params := domain.PaginateParam{
		Offset:        (page - 1) * size,
		Limit:         size,
		PaymentNumber: payload.PaymentNumber,
	}
	if payload.PaymentNumber != nil {
		params.PaymentNumber = payload.PaymentNumber
	}

	result := make(chan domain.Result)
	go func() {
		defer close(result)

		var inResult domain.Result
		payments, err := s.repo.Paginate(ctx, params)
		if err != nil {
			inResult.Error = err
			result <- inResult
			return
		}

		count, err := s.repo.Count(ctx, params.PaymentNumber)
		if err != nil {
			inResult.Error = err
			result <- inResult
			return
		}

		var data []domain.PaymentWithOrder
		for _, payment := range payments {
			order, err := s.repo.ShowOrder(ctx, payment.OrderID)
			if err != nil {
				inResult.Error = err
				result <- inResult
				return
			}

			createdDate := payment.CreatedAt.Format("2006-01-02 15:04:05")
			updatedDate := payment.UpdatedAt.Format("2006-01-02 15:04:05")
			paymentItem := domain.PaymentWithOrder{
				OrderNumber: order.OrderNumber,
				Discount:    order.Discount,
				TotalOrder:  order.Total,
				Payment: domain.Payment{
					ID:            payment.ID,
					OrderID:       payment.OrderID,
					PaymentNumber: payment.PaymentNumber,
					Method:        payment.Method,
					Total:         payment.Total,
					CreatedAt:     createdDate,
					UpdatedAt:     updatedDate,
				},
			}
			if !payment.PaidAt.IsZero() {
				paidDate := payment.PaidAt.Format("2006-01-02 15:04:05")
				paymentItem.PaidAt = &paidDate
			}

			data = append(data, paymentItem)
		}

		totalPages := int32(math.Ceil(float64(count) / float64(size)))
		meta := wrapper.Meta{
			Page:       page,
			Limit:      size,
			TotalData:  int32(count),
			TotalPages: totalPages,
		}

		paginate := domain.PaginateResponse{
			Items: data,
			Meta:  meta,
		}

		inResult.Value = paginate
		result <- inResult
	}()

	res := <-result
	if res.Error != nil {
		log.Println("paginate payment failed", res.Error.Error())
		return nil, res.Error
	}

	data := res.Value.(domain.PaginateResponse)

	return &data, nil
}

func (s *PaymentService) CompletePayment(ctx context.Context, orderId int64, payload domain.CompleteRequest) (*domain.CompleteResponse, error) {
	return nil, nil
}
