package v1

type SuccessResponse[T any] struct {
	Data T `json:"data"`
}

type SuccessListResponse[T any] struct {
	List      []T       `json:"list"`
	Paginator Paginator `json:"paginator"`
}

type Paginator struct {
	Skip  int   `json:"skip"`
	Limit int   `json:"limit"`
	Size  int   `json:"size"`
	Total int64 `json:"total"`
}
