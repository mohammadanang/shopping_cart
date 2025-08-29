package wrapper

type ErrResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

type OkResponse[T interface{}] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Status  string `json:"status"`
	Data    T      `json:"data,omitempty"`
}

type PaginateResponse[T interface{}] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Status  string `json:"status"`
	Data    []T    `json:"data,omitempty"`
	Meta    Meta   `json:"meta,omitempty"`
}

type Meta struct {
	Page       int `json:"page"`
	Limit      int `json:"size"`
	TotalData  int `json:"totalData"`
	TotalPages int `json:"totalPages"`
}
