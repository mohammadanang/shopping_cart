package service

import (
	"context"

	"github.com/mohammadanang/shopping-cart/db/dbgen"
	"github.com/mohammadanang/shopping-cart/internal/modules/cart/repository"
	"github.com/mohammadanang/shopping-cart/pkg/api"
)

type CartService struct {
	repo repository.Repository
}

func NewCartService(repo repository.Repository) Service {
	return &CartService{
		repo: repo,
	}
}

func (s *CartService) Add(ctx context.Context, payload api.AddCartJSONRequestBody) (*api.CartSuccessResponse, error) {
	// create order first with status = 'pending', total = 0, discount = 0, generated order_number

	item := dbgen.AddCartParams{
		ProductName: payload.ProductName,
		Qty:         int32(payload.Qty),
		Buyer:       payload.Buyer,
		Price:       float64(payload.Price),
		OrderID:     int64(payload.OrderId),
	}
	created, err := s.repo.Create(ctx, item)
	if err != nil {
		return nil, err
	}

	// update order total

	data := api.Cart{
		Id:          created.ID,
		ProductName: created.ProductName,
		Qty:         created.Qty,
		Buyer:       created.Buyer,
		Price:       int(created.Price),
		OrderId:     created.OrderID,
		CreatedAt:   created.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   created.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	return &api.CartSuccessResponse{
		Code:    201,
		Data:    data,
		Message: "Cart item added successfully",
		Status:  "success",
	}, nil
}

func (s *CartService) List(ctx context.Context) (*api.CartListSuccessResponse, error) {
	items, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}

	// relate order data
	var carts []api.Cart
	for _, item := range items {
		carts = append(carts, api.Cart{
			Id:          item.ID,
			ProductName: item.ProductName,
			Qty:         item.Qty,
			Buyer:       item.Buyer,
			OrderId:     item.OrderID,
			Price:       int(item.Price),
			CreatedAt:   item.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   item.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return &api.CartListSuccessResponse{
		Code:    200,
		Data:    carts,
		Message: "Cart items retrieved successfully",
		Status:  "success",
	}, nil
}

func (s *CartService) Remove(ctx context.Context, id int64) (*api.CartSuccessResponse, error) {
	find, err := s.repo.Show(ctx, id)
	if err != nil {
		return nil, err
	}

	err = s.repo.Delete(ctx, id)
	if err != nil {
		return nil, err
	}

	// count order if === 0 remove order
	// else edit order total

	data := api.Cart{
		Id:          find.ID,
		ProductName: find.ProductName,
		Qty:         find.Qty,
		Buyer:       find.Buyer,
		Price:       int(find.Price),
		OrderId:     find.OrderID,
		CreatedAt:   find.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   find.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	return &api.CartSuccessResponse{
		Code:    204,
		Data:    data,
		Message: "Cart item removed successfully",
		Status:  "success",
	}, nil
}
