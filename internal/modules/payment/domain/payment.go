package domain

import "github.com/mohammadanang/shopping-cart/pkg/wrapper"

type Payment struct {
	ID            int64   `json:"id"`
	OrderID       int64   `json:"order_id"`
	PaymentNumber string  `json:"payment_number"`
	Method        string  `json:"method"`
	Total         float64 `json:"total"`
	PaidAt        *string `json:"paid_at,omitempty"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
}

type PaymentWithOrder struct {
	Payment
	OrderNumber string  `json:"order_number"`
	Discount    float64 `json:"discount"`
	TotalOrder  float64 `json:"total_order"`
}

type PaginateRequest struct {
	Page          *int32  `json:"page" query:"page"`
	Size          *int32  `json:"size" query:"size"`
	PaymentNumber *string `json:"payment_number" query:"payment_number"`
}

type PaginateParam struct {
	Offset        int32   `json:"offset"`
	Limit         int32   `json:"limit"`
	PaymentNumber *string `json:"payment_number"`
}

type PaginateResponse struct {
	Items []PaymentWithOrder `json:"items"`
	Meta  wrapper.Meta       `json:"meta"`
}

type Result struct {
	Error error       `json:"error"`
	Value interface{} `json:"value"`
}

type CompleteRequest struct {
	Total float64 `json:"total"`
}

type CompleteResponse struct {
	OrderId       int64   `json:"order_id"`
	OrderNumber   string  `json:"order_number"`
	PaymentNumber string  `json:"payment_number"`
	Buyer         string  `json:"buyer"`
	Status        string  `json:"status"`
	Method        string  `json:"method"`
	TotalOrder    float64 `json:"total_order"`
	Discount      float64 `json:"discount"`
	TotalPaid     float64 `json:"total_paid"`
	PaidAt        string  `json:"paid_at"`
}
