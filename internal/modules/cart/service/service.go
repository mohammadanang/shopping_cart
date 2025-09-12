package service

import (
	"context"
	"log"

	"github.com/mohammadanang/shopping-cart/internal/modules/cart/domain"
	cartR "github.com/mohammadanang/shopping-cart/internal/modules/cart/repository"
)

type CartService struct {
	cartRepo cartR.Repository
}

func NewCartService(cartRepo cartR.Repository) Service {
	return &CartService{
		cartRepo: cartRepo,
	}
}

func (s *CartService) AddAndOrEdit(ctx context.Context, payload domain.AddRequest) (*domain.AddResponse, error) {
	result := make(chan domain.Result)
	go func() {
		defer close(result)

		var cartResult domain.Result
		upsert, err := s.cartRepo.TxCreateOrUpdate(ctx, payload)
		if err != nil {
			cartResult.Error = err
			result <- cartResult
			return
		}

		cartResult.Value = *upsert
		result <- cartResult
	}()

	res := <-result
	if res.Error != nil {
		log.Println("create or update cart failed", res.Error.Error())
		return nil, res.Error
	}

	data := res.Value.(domain.AddResponse)

	return &data, nil
}

func (s *CartService) List(ctx context.Context, payload domain.ListRequest) (*domain.AddResponse, error) {
	result := make(chan domain.Result)
	go func() {
		defer close(result)

		var cartResult domain.Result
		getOrder, err := s.cartRepo.ShowOrder(ctx, payload.OrderId)
		if err != nil {
			cartResult.Error = err
			result <- cartResult
			return
		}

		getCarts, err := s.cartRepo.List(ctx, getOrder.ID)
		if err != nil {
			cartResult.Error = err
			result <- cartResult
			return
		}

		var carts []domain.Cart
		for _, item := range getCarts {
			carts = append(carts, domain.Cart{
				Id:          item.ID,
				ProductName: item.ProductName,
				Qty:         item.Qty,
				Price:       item.Price,
				OrderId:     item.OrderID,
				CreatedAt:   &item.CreatedAt,
				UpdatedAt:   &item.UpdatedAt,
			})
		}

		response := domain.AddResponse{
			Type:        "list",
			OrderId:     getOrder.ID,
			OrderNumber: getOrder.OrderNumber,
			Items:       carts,
		}

		cartResult.Value = response
		result <- cartResult
	}()

	res := <-result
	if res.Error != nil {
		log.Println("get list cart failed", res.Error.Error())
		return nil, res.Error
	}

	data := res.Value.(domain.AddResponse)

	return &data, nil
}
