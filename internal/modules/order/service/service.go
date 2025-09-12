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

// func (s *OrderService) Show(ctx context.Context, id api.IdParam) (*api.OrderSuccessResponse, error) {
// 	show, err := s.repo.Show(ctx, int64(id))
// 	if err != nil {
// 		return nil, err
// 	}

// 	data := api.Order{
// 		Id:          show.ID,
// 		OrderNumber: show.OrderNumber,
// 		Discount:    float32(show.Discount),
// 		Status:      show.Status,
// 		Total:       float32(show.Total),
// 		CreatedAt:   show.CreatedAt.Format("2006-01-02 15:04:05"),
// 		UpdatedAt:   show.UpdatedAt.Format("2006-01-02 15:04:05"),
// 	}

// 	return &api.OrderSuccessResponse{
// 		Code:    200,
// 		Data:    data,
// 		Message: "Order retrieved successfully",
// 		Status:  "success",
// 	}, nil
// }

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
			return
		}

		count, err := s.repo.Count(ctx, params.OrderNumber)
		if err != nil {
			inResult.Error = err
			return
		}

		var data []domain.Order
		for _, order := range orders {
			createdDate := order.CreatedAt.Format("2006-01-02 15:04:05")
			updatedDate := order.UpdatedAt.Format("2006-01-02 15:04:05")

			data = append(data, domain.Order{
				ID:          order.ID,
				OrderNumber: order.OrderNumber,
				Discount:    order.Discount,
				Status:      order.Status,
				Total:       order.Total,
				CreatedAt:   &createdDate,
				UpdatedAt:   &updatedDate,
			})
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
