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

	"github.com/mohammadanang/shopping-cart/pkg/config"
)

type handlers struct {
	cartHandler.CartHandler
	// orderHandler.OrderHandler
}

type Server struct {
	app      *fiber.App
	store    dbgen.Store
	Handlers ApiServer
	conf     *config.Config
}

func NewServer(db *dbgen.Queries, store dbgen.Store, app *fiber.App, cfg *config.Config) *Server {
	cartRepo := cartRepository.NewCartRepository(db)
	cartSvc := cartService.NewCartService(store, cartRepo)
	cartHdl := cartHandler.NewCartHandler(cartSvc)

	return &Server{
		app:   app,
		store: store,
		conf:  cfg,
		Handlers: &handlers{
			// OrderHandler: orderHdl,
			CartHandler: cartHdl,
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
		Max:        10,              // allow 5 requests
		Expiration: 1 * time.Minute, // per 30 seconds
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
		AllowOrigins:     s.conf.Env.AllowedOrigins,
		AllowMethods:     s.conf.Env.AllowedMethods,
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
	s.Handlers.RegisterRoutes(apiRoutes)
}
