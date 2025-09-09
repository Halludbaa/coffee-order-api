package entity

import (
	"time"

	"github.com/google/uuid"
)

type MenuItem struct {
	ID          uuid.UUID `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Description string    `db:"description" json:"description,omitempty"`
	BasePrice   int64     `db:"base_price" json:"base_price"` // IDR: no cents
	CategoryID  uuid.UUID `db:"category_id" json:"category_id"`
	IsActive    bool      `db:"is_active" json:"is_active"`
	ImageURL    string    `db:"image_url" json:"image_url,omitempty"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

type StoreMenu struct {
	ID             uuid.UUID `db:"id" json:"id"`
	StoreID        uuid.UUID `db:"store_id" json:"store_id"`
	MenuItemID     uuid.UUID `db:"menu_item_id" json:"menu_item_id"`
	PriceOverride  *int64    `db:"price_override" json:"price_override,omitempty"` // nil = use base
	IsAvailable    bool      `db:"is_available" json:"is_available"`
	SortOrder      int       `db:"sort_order" json:"sort_order"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
}