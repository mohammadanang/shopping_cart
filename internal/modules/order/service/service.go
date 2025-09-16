package service

import (
	"context"
	"log"
	"math"

	"github.com/mohammadanang/shopping-cart/internal/modules/order/domain"
	"github.com/mohammadanang/shopping-cart/internal/modules/order/repository"
	"github.com/mohammadanang/shopping-cart/pkg/wrapper"
)

type OrderService struct {
	repo repository.Repository
}

func NewOrderService(repo repository.Repository) Service {
	return &OrderService{
		repo: repo,
	}
}

// func (s *OrderService) Edit(ctx context.Context, id api.IdParam, payload api.EditOrderJSONRequestBody) (*api.OrderSuccessResponse, error) {
// 	order := dbgen.EditOrderParams{
// 		ID:       int64(id),
// 		Discount: float64(payload.Discount),
// 		Status:   payload.Status,
// 		Total:    float64(payload.Total),
// 	}
// 	updated, err := s.repo.Update(ctx, order)
// 	if err != nil {
// 		return nil, err
// 	}

// 	data := api.Order{
// 		Id:          updated.ID,
// 		OrderNumber: updated.OrderNumber,
// 		Discount:    float32(updated.Discount),
// 		Status:      updated.Status,
// 		Total:       float32(updated.Total),
// 		CreatedAt:   updated.CreatedAt.Format("2006-01-02 15:04:05"),
// 		UpdatedAt:   updated.UpdatedAt.Format("2006-01-02 15:04:05"),
// 	}

// 	return &api.OrderSuccessResponse{
// 		Code:    200,
// 		Data:    data,
// 		Message: "Order updated successfully",
// 		Status:  "success",
// 	}, nil
// }

func (s *OrderService) Show(ctx context.Context, id int64) (*domain.ShowResponse, error) {
	result := make(chan domain.Result)
	go func() {
		defer close(result)

		var inResult domain.Result
		order, err := s.repo.Show(ctx, id)
		if err != nil {
			inResult.Error = err
			result <- inResult
			return
		}

		carts, err := s.repo.ListCart(ctx, order.ID)
		if err != nil {
			inResult.Error = err
			result <- inResult
			return
		}

		var cartItems []domain.Cart
		for _, item := range carts {
			cartItems = append(cartItems, domain.Cart{
				ProductName: item.ProductName,
				Qty:         item.Qty,
				Price:       item.Price,
			})
		}

		var paymentData *domain.Payment
		payment, err := s.repo.ShowPayment(ctx, order.ID)
		if err == nil && payment != nil {
			paymentData = &domain.Payment{
				Method:        payment.Method,
				Total:         payment.Total,
				PaymentNumber: payment.PaymentNumber,
			}
		} else {
			paymentData = nil
		}

		createdDate := order.CreatedAt.Format("2006-01-02 15:04:05")
		updatedDate := order.UpdatedAt.Format("2006-01-02 15:04:05")

		detailResp := domain.ShowResponse{
			Order: domain.Order{
				ID:          order.ID,
				OrderNumber: order.OrderNumber,
				Discount:    order.Discount,
				Status:      order.Status,
				Total:       order.Total,
				Buyer:       order.Buyer,
				CreatedAt:   &createdDate,
				UpdatedAt:   &updatedDate,
			},
			Carts: cartItems,
		}
		if paymentData != nil {
			detailResp.Payment = paymentData
		}

		inResult.Value = detailResp
		result <- inResult
	}()
	res := <-result
	if res.Error != nil {
		log.Println("show order failed", res.Error.Error())
		return nil, res.Error
	}

	data := res.Value.(domain.ShowResponse)

	return &data, nil
}

func (s *OrderService) Paginate(ctx context.Context, payload *domain.PaginateRequest) (*domain.PaginateResponse, error) {
	page := int32(1)
	size := int32(10)
	if payload.Size != nil {
		size = *payload.Size
	}

	if payload.Page != nil {
		page = *payload.Page
	}

	params := domain.PaginateParam{
		Offset:      (page - 1) * size,
		Limit:       size,
		OrderNumber: payload.OrderNumber,
	}
	if payload.OrderNumber != nil {
		params.OrderNumber = payload.OrderNumber
	}

	result := make(chan domain.Result)
	go func() {
		defer close(result)

		var inResult domain.Result
		orders, err := s.repo.Paginate(ctx, params)
		if err != nil {
			inResult.Error = err
			result <- inResult
			return
		}

		count, err := s.repo.Count(ctx, params.OrderNumber)
		if err != nil {
			inResult.Error = err
			result <- inResult
			return
		}

		var data []domain.OrderWithCarts
		for _, order := range orders {
			carts, err := s.repo.ListCart(ctx, order.ID)
			if err != nil {
				inResult.Error = err
				result <- inResult
				break
			}

			var cartItems []domain.Cart
			for _, cart := range carts {
				cartItems = append(cartItems, domain.Cart{
					ProductName: cart.ProductName,
					Qty:         cart.Qty,
					Price:       cart.Price,
				})
			}

			createdDate := order.CreatedAt.Format("2006-01-02 15:04:05")
			updatedDate := order.UpdatedAt.Format("2006-01-02 15:04:05")

			orderItem := domain.OrderWithCarts{
				Order: domain.Order{
					ID:          order.ID,
					OrderNumber: order.OrderNumber,
					Discount:    order.Discount,
					Status:      order.Status,
					Total:       order.Total,
					CreatedAt:   &createdDate,
					UpdatedAt:   &updatedDate,
				},
				Carts:        cartItems,
				TotalProduct: int32(len(carts)),
			}

			data = append(data, orderItem)
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
		log.Println("update order & payment failed", res.Error.Error())
		return nil, res.Error
	}

	data := res.Value.(domain.PaginateResponse)

	return &data, nil
}
