package domain

type PaginateRequest struct {
	Limit       int32   `json:"limit"`
	Offset      int32   `json:"offset"`
	OrderNumber *string `json:"order_number,omitempty"`
}
