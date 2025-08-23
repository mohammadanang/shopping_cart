package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/mohammadanang/shopping-cart/internal/modules/cart/service"
	"github.com/mohammadanang/shopping-cart/pkg/api"
)

type Handler struct {
	svc service.Service
}

func NewCartHandler(svc service.Service) CartHandler {
	return &Handler{
		svc: svc,
	}
}

func (h *Handler) AddCart(c *fiber.Ctx) error {
	body := new(api.AddCartJSONRequestBody)
	if err := c.BodyParser(body); err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(api.ErrorResponse{
			Message: "Invalid request body",
			Status:  "error",
			Code:    fiber.StatusBadRequest,
		})
	}

	create, err := h.svc.Add(c.Context(), *body)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(api.ErrorResponse{
			Message: "Failed to add item to cart",
			Status:  "error",
			Code:    fiber.StatusInternalServerError,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(create)
}

func (h *Handler) ListCarts(c *fiber.Ctx) error {
	all, err := h.svc.List(c.Context())
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(api.ErrorResponse{
			Message: "Failed to list carts",
			Status:  "error",
			Code:    fiber.StatusInternalServerError,
		})
	}

	return c.Status(fiber.StatusOK).JSON(all)
}

func (h *Handler) RemoveCart(c *fiber.Ctx, id api.IdParam) error {
	deleted, err := h.svc.Remove(c.Context(), int64(id))
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(api.ErrorResponse{
			Message: "Failed to remove cart",
			Status:  "error",
			Code:    fiber.StatusInternalServerError,
		})
	}

	return c.Status(fiber.StatusNoContent).JSON(deleted)
}
