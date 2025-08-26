package common

import "time"

// swagger:response Response
type Response struct {
	Code    int         `json:"code"` // Status Code. Default: 200
	Body    interface{} `json:"body"`
	Message string      `json:"message"`
}

// swagger:response ResponseBody
type ResponseBody struct {
	Message string      `json:"message"`
	Body    interface{} `json:"body"`
}

// swagger:response SuccessResponse
type SuccessResponse struct {
	Success bool `json:"success"`
}

// swagger:response PaginationResponse
type PaginationResponse[T any] struct {
	Items []T   `json:"items"`
	Total int64 `json:"total"`
}

type BaseFilterInput struct {
	TableName string         `json:"table_name"`
	CreatedAt *[2]*time.Time `json:"created_at"`
	Search    *string        `json:"search"`
	StartDate *time.Time     `json:"start_date"`
	EndDate   *time.Time     `json:"end_date"`
}

// swagger:parameters PaginateInput
type PaginateInput struct {
	Page   int                `json:"page"`
	Limit  int                `json:"limit"`
	Sorter *map[string]string `json:"sorter"`
}

func NewPaginationResponse[T any](items []T, total int64) *PaginationResponse[T] {
	return &PaginationResponse[T]{
		Items: items,
		Total: total,
	}
}

func NewSuccessReponse() *SuccessResponse {
	return &SuccessResponse{
		Success: true,
	}
}
