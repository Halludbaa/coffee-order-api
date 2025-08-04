package model

type WebResponse[T any] struct {
	Data    T      `json:"data,omitempty"`
	Code    int32  `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func NewWebResponse[T any](data T, code int32) *WebResponse[T] {
	return &WebResponse[T]{
		Data:    data,
		Code:    code,
		Message: "idk",
	}
}