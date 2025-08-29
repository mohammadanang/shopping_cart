package server

import "github.com/gofiber/fiber/v2"

type ApiServer interface {
	RegisterRoutes(r fiber.Router)
}
