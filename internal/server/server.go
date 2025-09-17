package server

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/mohammadanang/shopping-cart/db/dbgen"
	cartHandler "github.com/mohammadanang/shopping-cart/internal/modules/cart/handler"
	cartRepository "github.com/mohammadanang/shopping-cart/internal/modules/cart/repository"
	cartService "github.com/mohammadanang/shopping-cart/internal/modules/cart/service"
	orderHandler "github.com/mohammadanang/shopping-cart/internal/modules/order/handler"
	orderRepository "github.com/mohammadanang/shopping-cart/internal/modules/order/repository"
	orderService "github.com/mohammadanang/shopping-cart/internal/modules/order/service"
	paymentHandler "github.com/mohammadanang/shopping-cart/internal/modules/payment/handler"
	paymentRepository "github.com/mohammadanang/shopping-cart/internal/modules/payment/repository"
	paymentService "github.com/mohammadanang/shopping-cart/internal/modules/payment/service"

	"github.com/mohammadanang/shopping-cart/pkg/config"
)

type handlers struct {
	cart    cartHandler.CartHandler
	order   orderHandler.OrderHandler
	payment paymentHandler.PaymentHandler
}

type Server struct {
	app      *fiber.App
	store    dbgen.Store
	handlers handlers
	cfg      *config.Config
}

func NewServer(store dbgen.Store, app *fiber.App, cfg *config.Config) *Server {
	cartRepo := cartRepository.NewCartRepository(store)
	cartSvc := cartService.NewCartService(cartRepo)
	cartHdl := cartHandler.NewCartHandler(cartSvc)

	orderRepo := orderRepository.NewOrderRepository(store)
	orderSvc := orderService.NewOrderService(orderRepo)
	orderHdl := orderHandler.NewOrderHandler(orderSvc)

	paymentRepo := paymentRepository.NewPaymentRepository(store)
	paymentSvc := paymentService.NewPaymentService(paymentRepo)
	paymentHdl := paymentHandler.NewPaymentHandler(paymentSvc)

	return &Server{
		app:   app,
		store: store,
		cfg:   cfg,
		handlers: handlers{
			cart:    cartHdl,
			order:   orderHdl,
			payment: paymentHdl,
		},
	}
}

func (s *Server) SetMiddlewares() {
	// Recover middleware to catch panics in routes
	s.app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))
	s.app.Use(helmet.New())
	s.app.Use(idempotency.New())
	s.app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${ip} ${status} - ${latency} ${method} ${path}\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Local",
	}))
	s.app.Use(limiter.New(limiter.Config{
		Max:        60,
		Expiration: 1 * time.Minute,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP() // rate limit by client IP
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).
				JSON(fiber.Map{
					"error":  "Too many requests, slow down!",
					"status": "error",
					"code":   fiber.StatusTooManyRequests,
				})
		},
	}))
	s.app.Use(cors.New(cors.Config{
		AllowOrigins:     s.cfg.Env.AllowedOrigins,
		AllowMethods:     s.cfg.Env.AllowedMethods,
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		ExposeHeaders:    "Content-Length, X-Custom-Header",
		AllowCredentials: true,
	}))
}

func (s *Server) SetRoutes() {
	s.app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Shopping Cart API")
	})

	apiRoutes := s.app.Group("/api")
	s.handlers.cart.RegisterRoutes(apiRoutes)
	s.handlers.order.RegisterRoutes(apiRoutes)
	s.handlers.payment.RegisterRoutes(apiRoutes)
}
