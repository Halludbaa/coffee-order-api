package dto

type WebSuccessResponse[T any] struct {
	Data T `json:"data"`
}

type WebErrorResponse struct {
	Error string `json:"error"`
}