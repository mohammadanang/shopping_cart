package domain

import "time"

type Cart struct {
	Id          int64      `json:"id"`
	ProductName string     `json:"product_name"`
	Qty         int32      `json:"qty"`
	Price       float64    `json:"price"`
	OrderId     int64      `json:"order_id"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

type CartWithOrder struct {
	Cart
	OrderNumber string `json:"order_number"`
}

type AddInput struct {
	ProductName string  `json:"product_name"`
	Qty         int32   `json:"qty"`
	Price       float64 `json:"price"`
}

type AddRequest struct {
	OrderId *int64     `json:"order_id,omitempty"`
	Buyer   string     `json:"buyer"`
	Data    []AddInput `json:"data"`
}

type Result struct {
	Value interface{} `json:"value"`
	Error error       `json:"error"`
}

type AddResponse struct {
	Type        string `json:"type"`
	OrderId     int64  `json:"order_id"`
	OrderNumber string `json:"order_number"`
	Buyer       string `json:"buyer"`
	Items       []Cart `json:"items"`
}

type ListRequest struct {
	OrderId int64 `json:"order_id" query:"order_id"`
}
