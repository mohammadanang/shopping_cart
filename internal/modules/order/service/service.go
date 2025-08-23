package service

import (
	"context"

	"github.com/mohammadanang/shopping-cart/db/dbgen"
	"github.com/mohammadanang/shopping-cart/internal/modules/order/domain"
	"github.com/mohammadanang/shopping-cart/internal/modules/order/repository"
	"github.com/mohammadanang/shopping-cart/pkg/api"
)

type OrderService struct {
	repo repository.Repository
}

func NewOrderService(repo repository.Repository) Service {
	return &OrderService{
		repo: repo,
	}
}

func (s *OrderService) Add(ctx context.Context, payload api.AddOrderJSONRequestBody) (*api.OrderSuccessResponse, error) {
	order := dbgen.AddOrderParams{
		OrderNumber: payload.OrderNumber,
		Discount:    float64(payload.Discount),
		Status:      payload.Status,
		Total:       float64(payload.Total),
	}
	created, err := s.repo.Create(ctx, order)
	if err != nil {
		return nil, err
	}

	data := api.Order{
		Id:          created.ID,
		OrderNumber: created.OrderNumber,
		Discount:    float32(created.Discount),
		Status:      created.Status,
		Total:       float32(created.Total),
		CreatedAt:   created.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   created.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	return &api.OrderSuccessResponse{
		Code:    201,
		Data:    data,
		Message: "Order created successfully",
		Status:  "success",
	}, nil
}

func (s *OrderService) Edit(ctx context.Context, id api.IdParam, payload api.EditOrderJSONRequestBody) (*api.OrderSuccessResponse, error) {
	order := dbgen.EditOrderParams{
		ID:       int64(id),
		Discount: float64(payload.Discount),
		Status:   payload.Status,
		Total:    float64(payload.Total),
	}
	updated, err := s.repo.Update(ctx, order)
	if err != nil {
		return nil, err
	}

	data := api.Order{
		Id:          updated.ID,
		OrderNumber: updated.OrderNumber,
		Discount:    float32(updated.Discount),
		Status:      updated.Status,
		Total:       float32(updated.Total),
		CreatedAt:   updated.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   updated.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	return &api.OrderSuccessResponse{
		Code:    200,
		Data:    data,
		Message: "Order updated successfully",
		Status:  "success",
	}, nil
}

func (s *OrderService) Show(ctx context.Context, id api.IdParam) (*api.OrderSuccessResponse, error) {
	show, err := s.repo.Show(ctx, int64(id))
	if err != nil {
		return nil, err
	}

	data := api.Order{
		Id:          show.ID,
		OrderNumber: show.OrderNumber,
		Discount:    float32(show.Discount),
		Status:      show.Status,
		Total:       float32(show.Total),
		CreatedAt:   show.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   show.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	return &api.OrderSuccessResponse{
		Code:    200,
		Data:    data,
		Message: "Order retrieved successfully",
		Status:  "success",
	}, nil
}

func (s *OrderService) Paginate(ctx context.Context, page api.PageParam, size api.SizeParam, orderNumber *api.OrderNumberParam) (*api.OrderPaginateSuccessResponse, error) {
	params := domain.PaginateRequest{
		Limit:  size,
		Offset: (page - 1) * size,
	}
	if string(*orderNumber) != "" {
		params.OrderNumber = orderNumber
	}

	orders, err := s.repo.Paginate(ctx, params)
	if err != nil {
		return nil, err
	}

	count, err := s.repo.Count(ctx, orderNumber)
	if err != nil {
		return nil, err
	}

	var data []api.Order
	for _, order := range orders {
		data = append(data, api.Order{
			Id:          order.ID,
			OrderNumber: order.OrderNumber,
			Discount:    float32(order.Discount),
			Status:      order.Status,
			Total:       float32(order.Total),
			CreatedAt:   order.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   order.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	totalPages := int32(count) / size
	meta := api.MetaData{
		Page:       &page,
		Size:       &size,
		TotalData:  &count,
		TotalPages: &totalPages,
	}

	return &api.OrderPaginateSuccessResponse{
		Code:    200,
		Data:    data,
		Meta:    meta,
		Message: "Orders retrieved successfully",
		Status:  "success",
	}, nil
}
