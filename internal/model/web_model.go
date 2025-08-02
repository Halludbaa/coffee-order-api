package model

type WebResponse[T any] struct {
	Data    T      `json:"data,omitempty"`
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

func NewWebResponse[T any](data T, code int32) *WebResponse[T] {
	return &WebResponse[T]{
		Data:    data,
		Code:    code,
		Message: "idk",
	}
}