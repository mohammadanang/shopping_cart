package handler

import (
	"github.com/gofiber/fiber/v2"
)

type CartHandler interface {
	RegisterRoutes(r fiber.Router)
}
