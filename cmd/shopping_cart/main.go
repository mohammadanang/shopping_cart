package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/mohammadanang/shopping-cart/internal/server"
	"github.com/mohammadanang/shopping-cart/pkg/api"
	"github.com/mohammadanang/shopping-cart/pkg/config"
	"github.com/mohammadanang/shopping-cart/pkg/database"
	"github.com/mohammadanang/shopping-cart/pkg/docs"
)

func main() {
	// Global recover for main goroutine
	defer func() {
		if r := recover(); r != nil {
			log.Printf("❌ Application panicked: %v", r)
		}
	}()

	// === CONFIGURATION ===
	cfg := config.LoadConfig(".")
	ctx := context.Background()
	db := <-database.NewPostgres(ctx, cfg)
	if db.Err != nil {
		log.Fatalf("DB connection failed: %v", db.Err)
	}

	defer db.Pool.Close()
	// === CONFIGURATION ===

	app := fiber.New()
	appServer := server.NewServer(db.Queries, app)

	// === MIDDLEWARES ===
	appServer.SetMiddlewares()
	// === MIDDLEWARES ===

	// === ROUTES ===
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Shopping Cart API")
	})

	openapi := docs.NewDocsHandler()
	openapi.RegisterRoutes(app)
	api.RegisterHandlers(app, appServer.Handlers)
	// === ROUTES ===

	// Run server in goroutine
	appPort := cfg.Port
	log.Println("🚀 Server running on :" + appPort)
	go func() {
		if err := app.Listen(":" + cfg.Port); err != nil {
			log.Printf("❌ Fiber error: %v", err)
		}
	}()

	// === SHUTDOWN ===
	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("⏳ Shutting down server...")
	_ = app.Shutdown() // Shutdown Fiber app

	db.Pool.Close() // Close DB pool
	log.Println("✅ Database pool closed, shutdown complete")
	// === SHUTDOWN ===
}
