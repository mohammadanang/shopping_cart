package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/mohammadanang/shopping-cart/internal/modules/order/service"
	"github.com/mohammadanang/shopping-cart/pkg/api"
)

type Handler struct {
	svc service.Service
}

func NewOrderHandler(svc service.Service) OrderHandler {
	return &Handler{
		svc: svc,
	}
}

func (h *Handler) AddOrder(c *fiber.Ctx) error {
	body := new(api.AddOrderJSONRequestBody)
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
			Message: "Failed to add order",
			Status:  "error",
			Code:    fiber.StatusInternalServerError,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(create)
}

func (h *Handler) ListOrders(c *fiber.Ctx, params api.ListOrdersParams) error {
	paginate, err := h.svc.Paginate(c.Context(), *params.Page, *params.Size, params.OrderNumber)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(api.ErrorResponse{
			Message: "Failed to list orders",
			Status:  "error",
			Code:    fiber.StatusInternalServerError,
		})
	}

	return c.Status(fiber.StatusOK).JSON(paginate)
}

func (h *Handler) ShowOrder(c *fiber.Ctx, id api.IdParam) error {
	order, err := h.svc.Show(c.Context(), id)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(api.ErrorResponse{
			Message: "Failed to retrieve order",
			Status:  "error",
			Code:    fiber.StatusInternalServerError,
		})
	}

	return c.Status(fiber.StatusOK).JSON(order)
}

func (h *Handler) EditOrder(c *fiber.Ctx, id api.IdParam) error {
	body := new(api.EditOrderJSONRequestBody)
	if err := c.BodyParser(body); err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(api.ErrorResponse{
			Message: "Invalid request body",
			Status:  "error",
			Code:    fiber.StatusBadRequest,
		})
	}

	update, err := h.svc.Edit(c.Context(), id, *body)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(api.ErrorResponse{
			Message: "Failed to update order",
			Status:  "error",
			Code:    fiber.StatusInternalServerError,
		})
	}

	return c.Status(fiber.StatusOK).JSON(update)
}
