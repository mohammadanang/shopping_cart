package server

import (
	"github.com/mohammadanang/shopping-cart/db/dbgen"
	cartHandler "github.com/mohammadanang/shopping-cart/internal/modules/cart/handler"
	cartRepository "github.com/mohammadanang/shopping-cart/internal/modules/cart/repository"
	cartService "github.com/mohammadanang/shopping-cart/internal/modules/cart/service"
	orderHandler "github.com/mohammadanang/shopping-cart/internal/modules/order/handler"
	orderRepository "github.com/mohammadanang/shopping-cart/internal/modules/order/repository"
	orderService "github.com/mohammadanang/shopping-cart/internal/modules/order/service"
	"github.com/mohammadanang/shopping-cart/pkg/api"
)

type Server struct {
	orderHandler.OrderHandler
	cartHandler.CartHandler
}

func NewServer(db *dbgen.Queries) api.ServerInterface {
	cartRepo := cartRepository.NewCartRepository(db)
	cartSvc := cartService.NewCartService(cartRepo)
	cartHdl := cartHandler.NewCartHandler(cartSvc)

	orderRepo := orderRepository.NewOrderRepository(db)
	orderSvc := orderService.NewOrderService(orderRepo)
	orderHdl := orderHandler.NewOrderHandler(orderSvc)

	return &Server{
		CartHandler:  cartHdl,
		OrderHandler: orderHdl,
	}
}
