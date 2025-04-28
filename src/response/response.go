package response

import "app/src/model"

type Common struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type SuccessWithUser struct {
	Code    int        `json:"code"`
	Status  string     `json:"status"`
	Message string     `json:"message"`
	User    model.User `json:"user"`
}

type SuccessWithTokens struct {
	Code    int        `json:"code"`
	Status  string     `json:"status"`
	Message string     `json:"message"`
	User    model.User `json:"user"`
	Tokens  Tokens     `json:"tokens"`
}

type SuccessWithData[T any] struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

type Pagination struct {
	Page         int   `json:"page"`
	Limit        int   `json:"limit"`
	TotalPages   int64 `json:"totalPages"`
	TotalResults int64 `json:"totalResults"`
}

type SuccessWithPaginate[T any] struct {
	Code       int        `json:"code"`
	Status     string     `json:"status"`
	Message    string     `json:"message"`
	Data       []T        `json:"data"`
	Pagination Pagination `json:"pagination"`
}

type ErrorDetails struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
}
