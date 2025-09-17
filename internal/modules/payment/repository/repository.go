package repository

import (
	"context"

	"github.com/mohammadanang/shopping-cart/db/dbgen"
	"github.com/mohammadanang/shopping-cart/internal/modules/payment/domain"
)

type PaymentRepository struct {
	store dbgen.Store
}

func NewPaymentRepository(store dbgen.Store) Repository {
	return &PaymentRepository{
		store: store,
	}
}

func (r *PaymentRepository) Paginate(ctx context.Context, payload domain.PaginateParam) ([]*dbgen.Payment, error) {
	if payload.PaymentNumber != nil {
		payload := dbgen.PaginatePaymentsWithParamsParams{
			Limit:         payload.Limit,
			Offset:        payload.Offset,
			PaymentNumber: *payload.PaymentNumber,
		}
		payments, err := r.store.PaginatePaymentsWithParams(ctx, &payload)
		if err != nil {
			return nil, err
		}

		return payments, nil
	}

	params := dbgen.PaginatePaymentsParams{
		Limit:  payload.Limit,
		Offset: payload.Offset,
	}
	payments, err := r.store.PaginatePayments(ctx, &params)
	if err != nil {
		return nil, err
	}

	return payments, nil
}

func (r *PaymentRepository) Count(ctx context.Context, paymentNumber *string) (int64, error) {
	if paymentNumber != nil {
		count, err := r.store.CountPaymentsByPaymentNumber(ctx, *paymentNumber)
		if err != nil {
			return 0, err
		}

		return count, nil
	}

	count, err := r.store.CountPayments(ctx)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *PaymentRepository) ShowOrder(ctx context.Context, orderId int64) (*dbgen.Order, error) {
	order, err := r.store.GetOrder(ctx, orderId)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (r *PaymentRepository) TxCompletePayment(ctx context.Context, orderId int64, payload domain.CompleteRequest) (*domain.CompleteResponse, error) {
	var result domain.CompleteResponse
	err := r.store.ExecTx(ctx, func(q *dbgen.Queries) error {
		getOrder, err := q.GetOrder(ctx, orderId)
		if err != nil {
			return err
		}

		updateOrder, err := q.EditOrder(ctx, &dbgen.EditOrderParams{
			ID:       getOrder.ID,
			Discount: getOrder.Discount,
			Status:   "completed",
			Total:    getOrder.Total,
		})
		if err != nil {
			return err
		}

		result.OrderId = updateOrder.ID
		result.OrderNumber = updateOrder.OrderNumber
		result.Buyer = updateOrder.Buyer
		result.Status = updateOrder.Status
		result.Discount = updateOrder.Discount
		result.TotalOrder = updateOrder.Total

		getPayment, err := q.GetPaymentByOrder(ctx, orderId)
		if err != nil {
			return err
		}

		result.PaymentNumber = getPayment.PaymentNumber
		result.Method = getPayment.Method

		approvePayment, err := q.ApprovePayment(ctx, &dbgen.ApprovePaymentParams{
			ID:    getPayment.ID,
			Total: getPayment.Total,
		})
		if err != nil {
			return err
		}

		paidDate := approvePayment.PaidAt.Format("2006-01-02 15:04:05")
		result.TotalPaid = approvePayment.Total
		result.PaidAt = paidDate

		return nil
	})

	return &result, err
}
