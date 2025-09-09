package entity

import (
	"time"
)

type User struct {
	ID           int       `db:"id" json:"id"`
	FullName     string    `db:"full_name" json:"full_name"`
	Email        string    `db:"email" json:"email,omitempty"`
	Role         string    `db:"role" json:"role"` // barista, manager, admin
	StoreID      int       `db:"store_id" json:"store_id"`
	PinHash      []byte    `db:"pin_hash" json:"-"` // never expose
	IsActive     bool      `db:"is_active" json:"is_active"`
	MustResetPin bool      `db:"must_reset_pin" json:"must_reset_pin"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
}
