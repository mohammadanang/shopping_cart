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
	Page       int32 `json:"page"`
	Limit      int32 `json:"size"`
	TotalData  int32 `json:"totalData"`
	TotalPages int32 `json:"totalPages"`
}
