package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/mohammadanang/shopping-cart/internal/modules/cart/domain"
	"github.com/mohammadanang/shopping-cart/internal/modules/cart/service"
	"github.com/mohammadanang/shopping-cart/pkg/wrapper"
)

type Handler struct {
	svc service.Service
}

func NewCartHandler(svc service.Service) CartHandler {
	return &Handler{
		svc: svc,
	}
}

func (h *Handler) RegisterRoutes(r fiber.Router) {
	cartV1 := r.Group("/cart/v1")

	cartV1.Post("/", h.addCart)
}

func (h *Handler) addCart(c *fiber.Ctx) error {
	body := new(domain.AddRequest)
	if err := c.BodyParser(body); err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(wrapper.ErrResponse{
			Code:    fiber.StatusBadRequest,
			Status:  "error",
			Message: "Invalid request body",
		})
	}

	create, err := h.svc.AddAndOrEdit(c.Context(), *body)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(wrapper.ErrResponse{
			Message: "Failed to add item to cart",
			Status:  "error",
			Code:    fiber.StatusInternalServerError,
		})
	}

	var result []*domain.Cart
	for _, item := range create.Items {
		result = append(result, &item)
	}

	return c.Status(fiber.StatusCreated).JSON(wrapper.OkResponse[[]*domain.Cart]{
		Code:    fiber.StatusCreated,
		Status:  "success",
		Message: "Item added to cart",
		Data:    result,
	})
}

// func (h *Handler) listCarts(c *fiber.Ctx) error {
// 	all, err := h.svc.List(c.Context())
// 	if err != nil {
// 		log.Println(err.Error())
// 		return c.Status(fiber.StatusInternalServerError).JSON(wrapper.ErrResponse{
// 			Message: "Failed to list carts",
// 			Status:  "error",
// 			Code:    fiber.StatusInternalServerError,
// 		})
// 	}

// 	return c.Status(fiber.StatusOK).JSON(all)
// }
