package entity

import (
	"time"

	"github.com/google/uuid"
)

type Store struct {
	ID        uuid.UUID       `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Location  string    `db:"location" json:"location,omitempty"`
	Address   string    `db:"address" json:"address,omitempty"`
	Phone     string    `db:"phone" json:"phone,omitempty"`
	Email     string    `db:"email" json:"email,omitempty"`
	StoreSlug string    `db:"store_slug" json:"store_slug"`
	IsActive  bool      `db:"is_active" json:"is_active"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}