package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/mohammadanang/shopping-cart/internal/server"
	"github.com/mohammadanang/shopping-cart/pkg/api"
	"github.com/mohammadanang/shopping-cart/pkg/config"
	"github.com/mohammadanang/shopping-cart/pkg/database"
	swgui "github.com/swaggest/swgui/v5"
)

func main() {
	cfg := config.LoadConfig(".")
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Shopping Cart API")
	})

	app.Get("/openapi.json", func(c *fiber.Ctx) error {
		swagger, err := api.GetSwagger()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.JSON(swagger)
	})

	swaggerUI := swgui.NewHandler(
		"Shopping Cart API Docs",
		"/openapi.json",
		"/docs",
	)
	app.Get("/docs/*", adaptor.HTTPHandler(swaggerUI))

	ctx := context.Background()
	dbConfig, dbPool := database.NewPostgresPool(ctx, cfg)
	defer dbPool.Close()

	appServer := server.NewServer(dbConfig)
	api.RegisterHandlers(app, appServer)

	appPort := cfg.Port
	log.Println("🚀 Server running on :" + appPort)
	if err := app.Listen(":" + appPort); err != nil {
		log.Fatal(err)
	}
}
