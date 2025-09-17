package handler

import (
	"github.com/gofiber/fiber/v2"
)

type PaymentHandler interface {
	RegisterRoutes(r fiber.Router)
}
