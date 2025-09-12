package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/mohammadanang/shopping-cart/db/dbgen"
	"github.com/mohammadanang/shopping-cart/internal/server"
	"github.com/mohammadanang/shopping-cart/pkg/config"
	"github.com/mohammadanang/shopping-cart/pkg/database"
)

func main() {
	// Global recover for main goroutine
	defer func() {
		if r := recover(); r != nil {
			log.Printf("❌ Application panicked: %v", r)
		}
	}()

	// === CONFIGURATION ===
	cfg := config.NewConfig(".")
	ctx := context.Background()
	db := <-database.NewPostgres(ctx, cfg)
	if db.Err != nil {
		log.Fatalf("DB connection failed: %v", db.Err)
	}

	defer db.Pool.Close()
	dbStore := dbgen.NewStore(db.Pool)
	// === CONFIGURATION ===

	app := fiber.New()
	appServer := server.NewServer(dbStore, app, cfg)

	// === MIDDLEWARES ===
	appServer.SetMiddlewares()
	// === MIDDLEWARES ===

	// === ROUTES ===
	appServer.SetRoutes()
	// === SWAGGER DOCS ===
	app.Static("/openapi.json", filepath.Join("api", "api.yaml"))
	app.Static("/docs", filepath.Join("docs", "swagger-ui"))
	// === SWAGGER DOCS ===
	// === ROUTES ===

	// Run server in goroutine
	appPort := cfg.Env.Port
	log.Println("🚀 Server running on :" + appPort)
	go func() {
		if err := app.Listen(":" + appPort); err != nil {
			log.Printf("❌ Fiber error: %v", err)
		}
	}()

	// === GRACEFULLY SHUTDOWN ===
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("⏳ Shutting down server...")
	_ = app.Shutdown() // Shutdown Fiber app

	db.Pool.Close() // Close DB pool
	log.Println("✅ Database pool closed, shutdown complete")
	// === GRACEFULLY SHUTDOWN ===
}
