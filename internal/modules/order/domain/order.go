package domain

import (
	"github.com/mohammadanang/shopping-cart/pkg/wrapper"
)

type Order struct {
	ID          int64   `json:"id"`
	OrderNumber string  `json:"order_number"`
	Discount    float64 `json:"discount"`
	Status      string  `json:"status"`
	Total       float64 `json:"total"`
	Buyer       string  `json:"buyer"`
	CreatedAt   *string `json:"created_at,omitempty"`
	UpdatedAt   *string `json:"updated_at,omitempty"`
}

type Cart struct {
	ProductName string  `json:"product_name"`
	Qty         int32   `json:"qty"`
	Price       float64 `json:"price"`
}

type OrderWithCarts struct {
	Order
	TotalProduct int32  `json:"total_product"`
	Carts        []Cart `json:"carts"`
}

type UpdateRequest struct {
	Discount      float64 `json:"discount"`
	PaymentMethod string  `json:"payment_method"`
}

type UpdateResponse struct {
	OrderId       int64   `json:"order_id"`
	OrderNumber   string  `json:"order_number"`
	Buyer         string  `json:"buyer"`
	PaymentNumber string  `json:"payment_number"`
	Status        string  `json:"status"`
	TotalOrder    float64 `json:"total_order"`
	Discount      float64 `json:"discount"`
	TotalPaid     float64 `json:"total_paid"`
}

type Result struct {
	Value interface{} `json:"value"`
	Error error       `json:"error"`
}

type PaginateRequest struct {
	Page        *int32  `json:"page" query:"page"`
	Size        *int32  `json:"size" query:"size"`
	OrderNumber *string `json:"order_number" query:"order_number"`
}

type PaginateParam struct {
	Offset      int32   `json:"offset"`
	Limit       int32   `json:"limit"`
	OrderNumber *string `json:"order_number"`
}

type DetailRequest struct {
	ID int64 `json:"id"`
}

type PaginateResponse struct {
	Items []OrderWithCarts `json:"items"`
	Meta  wrapper.Meta     `json:"meta"`
}

type Payment struct {
	PaymentNumber string  `json:"payment_number"`
	Method        string  `json:"method"`
	Total         float64 `json:"total"`
}

type ShowResponse struct {
	Order
	Carts    []Cart `json:"carts"`
	*Payment `json:"payment"`
}
