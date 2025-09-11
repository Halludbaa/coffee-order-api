package entity

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID        uuid.UUID `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Icon      string    `db:"icon" json:"icon,omitempty"`
	Color     string    `db:"color" json:"color"`
	IsActive  bool      `db:"is_active" json:"is_active"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type StoreCategory struct {
	ID          uuid.UUID    `db:"id" json:"id"`
	StoreID     uuid.UUID    `db:"store_id" json:"store_id"`
	CategoryID  uuid.UUID    `db:"category_id" json:"category_id"`
	Name        string `db:"name" json:"name"` // override
	IsVisible   bool   `db:"is_visible" json:"is_visible"`
	SortOrder   int    `db:"sort_order" json:"sort_order"`
}