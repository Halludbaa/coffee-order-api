package dto

import (
	"time"

	"github.com/google/uuid"
)

type StoreRequest struct {
    Name      string `json:"name" validate:"required,max=100"`
    Location  string `json:"location"`
    Address   string `json:"address"`
    Phone     string `json:"phone" validate:"max=20"`
    Email     string `json:"email" validate:"email,max=100"`
    StoreSlug string `json:"store_slug,omitempty"`
    IsActive  bool   `json:"is_active"`
}

type StoreResponse struct {
    ID        uuid.UUID `json:"id"`
    Name      string    `json:"name"`
    Location  string    `json:"location"`
    Address   string    `json:"address"`
    Phone     string    `json:"phone"`
    Email     string    `json:"email"`
    StoreSlug string    `json:"store_slug"`
    IsActive  bool      `json:"is_active"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}