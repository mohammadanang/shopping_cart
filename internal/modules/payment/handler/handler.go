package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/mohammadanang/shopping-cart/internal/modules/payment/domain"
	"github.com/mohammadanang/shopping-cart/internal/modules/payment/service"
	"github.com/mohammadanang/shopping-cart/pkg/wrapper"
)

type Handler struct {
	svc service.Service
}

func NewPaymentHandler(svc service.Service) PaymentHandler {
	return &Handler{
		svc: svc,
	}
}

func (h *Handler) RegisterRoutes(r fiber.Router) {
	paymentV1 := r.Group("/payment/v1")

	paymentV1.Get("/", h.listPayment)
}

func (h *Handler) listPayment(c *fiber.Ctx) error {
	queryParam := new(domain.PaginateRequest)
	if err := c.QueryParser(queryParam); err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(wrapper.ErrResponse{
			Message: "Invalid request body",
			Status:  "error",
			Code:    fiber.StatusBadRequest,
		})
	}

	paginate, err := h.svc.Paginate(c.Context(), queryParam)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(wrapper.ErrResponse{
			Message: "Failed to list payments",
			Status:  "error",
			Code:    fiber.StatusInternalServerError,
		})
	}

	var payments []*domain.PaymentWithOrder
	for _, item := range paginate.Items {
		payments = append(payments, &item)
	}

	return c.Status(fiber.StatusOK).JSON(wrapper.PaginateResponse[*domain.PaymentWithOrder]{
		Code:    fiber.StatusOK,
		Message: "Get payments successfully",
		Status:  "success",
		Data:    payments,
		Meta:    paginate.Meta,
	})
}
