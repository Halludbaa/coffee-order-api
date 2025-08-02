package model

type PingRequest struct {
	Message string `json:"message" binding:"required,min=8"`
}