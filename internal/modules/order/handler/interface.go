package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mohammadanang/shopping-cart/pkg/api"
)

type OrderHandler interface {
	ListOrders(c *fiber.Ctx, params api.ListOrdersParams) error
	AddOrder(c *fiber.Ctx) error
	ShowOrder(c *fiber.Ctx, id api.IdParam) error
	EditOrder(c *fiber.Ctx, id api.IdParam) error
}
