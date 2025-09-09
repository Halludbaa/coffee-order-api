package dto

import "github.com/google/uuid"

type MenuItemRequest struct {
    Name        string  `json:"name" validate:"required,min=2"`
    Description string  `json:"description,omitempty"`
    BasePrice   int64   `json:"base_price" validate:"required,gt=0"`
    CategoryID  uuid.UUID `json:"category_id" validate:"required"`
    ImageURL    string  `json:"image_url,omitempty"`
    IsActive    bool    `json:"is_active"`
}

type MenuItemResponse struct {
    ID          uuid.UUID `json:"id"`
    Name        string    `json:"name"`
    Description string    `json:"description,omitempty"`
    BasePrice   int64     `json:"base_price"`
    CategoryID  uuid.UUID `json:"category_id"`
    IsActive    bool      `json:"is_active"`
    ImageURL    string    `json:"image_url,omitempty"`
}

type StoreMenuRequest struct {
    PriceOverride *int64 `json:"price_override"`
    IsAvailable   bool   `json:"is_available"`
    SortOrder     int    `json:"sort_order"`
}