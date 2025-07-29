package apperrors

import "fmt"

type Type = int

const (
	Authorization        Type = 401
	BadRequest           Type = 400
	Conflict             Type = 409
	Internal             Type = 500
	NotFound             Type = 404
	PayloadTooLarge      Type = 413
	ServiceUnavailable   Type = 503
	UnsupportedMediaType Type = 415
)

type APIError struct {
	Field	string 	`json:"field"`
	Message	string	`json:"message"`
}


type Apperrors struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Errors	[]APIError `json:"errors,omitempty"`
}

func (aer *Apperrors) Error() string {
	return aer.Message
}

func PasswordNotMatch() []APIError {
	return []APIError{
		{
			Field: "Password",
			Message: "Passwords don't match",
		},
	}
}

func NewInternal() *Apperrors {
	return &Apperrors{
		Code:    Internal,
		Message: "Internal server error.",
	}
}

func NewAuthorization(reason string) *Apperrors {
	return &Apperrors{
		Code:    Authorization,
		Message: reason,
	}
}

func NewBadRequest(reason string, errors []APIError) *Apperrors {
	return &Apperrors{
		Code:    BadRequest,
		Message: fmt.Sprintf("Bad request. Reason: %v", reason),
		Errors: errors,
	}
}

func NewConflict(name string, value string) *Apperrors {
	return &Apperrors{
		Code:    Conflict,
		Message: fmt.Sprintf("resource: %v with value: %v already exists", name, value),
	}
}

func NewNotFound(name string, value string) *Apperrors {
	return &Apperrors{
		Code:    NotFound,
		Message: fmt.Sprintf("resource: %v with value: %v not found", name, value),
	}
}