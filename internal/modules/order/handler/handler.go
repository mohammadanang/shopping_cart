package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/mohammadanang/shopping-cart/internal/modules/order/domain"
	"github.com/mohammadanang/shopping-cart/internal/modules/order/service"
	"github.com/mohammadanang/shopping-cart/pkg/wrapper"
)

type Handler struct {
	svc service.Service
}

func NewOrderHandler(svc service.Service) OrderHandler {
	return &Handler{
		svc: svc,
	}
}

func (h *Handler) RegisterRoutes(r fiber.Router) {
	orderV1 := r.Group("/order/v1")

	orderV1.Get("/", h.listOrder)
	orderV1.Put("/:id", h.updateOrderAndPayment)
	orderV1.Get("/:id", h.showOrder)
}

func (h *Handler) updateOrderAndPayment(c *fiber.Ctx) error {
	return nil
}

func (h *Handler) showOrder(c *fiber.Ctx) error {
	return nil
}

// func (h *Handler) AddOrder(c *fiber.Ctx) error {
// 	body := new(api.AddOrderJSONRequestBody)
// 	if err := c.BodyParser(body); err != nil {
// 		log.Println(err.Error())
// 		return c.Status(fiber.StatusBadRequest).JSON(api.ErrorResponse{
// 			Message: "Invalid request body",
// 			Status:  "error",
// 			Code:    fiber.StatusBadRequest,
// 		})
// 	}

// 	create, err := h.svc.Add(c.Context(), *body)
// 	if err != nil {
// 		log.Println(err.Error())
// 		return c.Status(fiber.StatusInternalServerError).JSON(api.ErrorResponse{
// 			Message: "Failed to add order",
// 			Status:  "error",
// 			Code:    fiber.StatusInternalServerError,
// 		})
// 	}

// 	return c.Status(fiber.StatusCreated).JSON(create)
// }

func (h *Handler) listOrder(c *fiber.Ctx) error {
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
			Message: "Failed to list orders",
			Status:  "error",
			Code:    fiber.StatusInternalServerError,
		})
	}

	var orders []*domain.Order
	for _, item := range paginate.Items {
		orders = append(orders, &item)
	}

	return c.Status(fiber.StatusOK).JSON(wrapper.PaginateResponse[*domain.Order]{
		Code:    fiber.StatusOK,
		Message: "Get orders successfully",
		Status:  "success",
		Data:    orders,
		Meta:    paginate.Meta,
	})
}

// func (h *Handler) ShowOrder(c *fiber.Ctx, id api.IdParam) error {
// 	order, err := h.svc.Show(c.Context(), id)
// 	if err != nil {
// 		log.Println(err.Error())
// 		return c.Status(fiber.StatusInternalServerError).JSON(api.ErrorResponse{
// 			Message: "Failed to retrieve order",
// 			Status:  "error",
// 			Code:    fiber.StatusInternalServerError,
// 		})
// 	}

// 	return c.Status(fiber.StatusOK).JSON(order)
// }
