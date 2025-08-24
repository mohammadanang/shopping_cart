package docs

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/mohammadanang/shopping-cart/pkg/api"
	swgui "github.com/swaggest/swgui/v5"
)

type DocsHandler struct{}

func NewDocsHandler() *DocsHandler {
	return &DocsHandler{}
}

func (h *DocsHandler) RegisterRoutes(app *fiber.App) {
	app.Get("/openapi.json", h.getJson)
	app.Get("/docs/*", adaptor.HTTPHandler(h.getHandler()))
}

func (h *DocsHandler) getJson(c *fiber.Ctx) error {
	swagger, err := api.GetSwagger()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(swagger)
}

func (h *DocsHandler) getHandler() *swgui.Handler {
	swaggerUI := swgui.NewHandler(
		"Shopping Cart API Docs",
		"/openapi.json",
		"/docs",
	)

	return swaggerUI
}
