package domain

import "time"

type Cart struct {
	Id          int64      `json:"id"`
	ProductName string     `json:"product_name"`
	Qty         int32      `json:"qty"`
	Buyer       string     `json:"buyer"`
	Price       float64    `json:"price"`
	OrderId     int64      `json:"order_id"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

type AddInput struct {
	ProductName string  `json:"product_name"`
	Qty         int32   `json:"qty"`
	Buyer       string  `json:"buyer"`
	Price       float64 `json:"price"`
	OrderId     int64   `json:"order_id"`
}

type AddRequest struct {
	OrderId *int64     `json:"order_id"`
	Data    []AddInput `json:"data"`
}

type Result struct {
	Value interface{} `json:"value"`
	Error error       `json:"error"`
}

type AddResponse struct {
	OrderId     int64  `json:"order_id"`
	OrderNumber string `json:"order_number"`
	Items       []Cart `json:"items"`
}
