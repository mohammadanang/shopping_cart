package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/mohammadanang/shopping-cart/db/dbgen"
	cartHandler "github.com/mohammadanang/shopping-cart/internal/modules/cart/handler"
	cartRepository "github.com/mohammadanang/shopping-cart/internal/modules/cart/repository"
	cartService "github.com/mohammadanang/shopping-cart/internal/modules/cart/service"
	orderHandler "github.com/mohammadanang/shopping-cart/internal/modules/order/handler"
	orderRepository "github.com/mohammadanang/shopping-cart/internal/modules/order/repository"
	orderService "github.com/mohammadanang/shopping-cart/internal/modules/order/service"
	"github.com/mohammadanang/shopping-cart/pkg/api"
)

type handlers struct {
	cartHandler.CartHandler
	orderHandler.OrderHandler
}

type Server struct {
	app      *fiber.App
	Handlers api.ServerInterface
}

func NewServer(db *dbgen.Queries, app *fiber.App) *Server {
	cartRepo := cartRepository.NewCartRepository(db)
	cartSvc := cartService.NewCartService(cartRepo)
	cartHdl := cartHandler.NewCartHandler(cartSvc)

	orderRepo := orderRepository.NewOrderRepository(db)
	orderSvc := orderService.NewOrderService(orderRepo)
	orderHdl := orderHandler.NewOrderHandler(orderSvc)

	return &Server{
		app: app,
		Handlers: &handlers{
			OrderHandler: orderHdl,
			CartHandler:  cartHdl,
		},
	}
}

func (s *Server) SetMiddlewares() {
	// Recover middleware to catch panics in routes
	s.app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))
	s.app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${ip} ${status} - ${latency} ${method} ${path}\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Local",
	}))
}
