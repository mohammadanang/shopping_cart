package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mohammadanang/shopping-cart/pkg/api"
)

type CartHandler interface {
	ListCarts(c *fiber.Ctx) error
	AddCart(c *fiber.Ctx) error
	RemoveCart(c *fiber.Ctx, id api.IdParam) error
}
