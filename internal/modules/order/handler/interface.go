package handler

import (
	"github.com/gofiber/fiber/v2"
)

type OrderHandler interface {
	RegisterRoutes(r fiber.Router)
}
