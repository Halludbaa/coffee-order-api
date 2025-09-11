package entity

import (
	"time"

	"github.com/google/uuid"
)

type CustomizationGroup struct {
	ID         uuid.UUID `db:"id"`
	Name       string    `db:"name"`
	StoreID    uuid.UUID `db:"store_id"`
	IsRequired bool      `db:"is_required"`
	SortOrder  int       `db:"sort_order"`
	CreatedAt  time.Time `db:"created_at"`
	Options    []*CustomizationOption `db:"-"`
}

type CustomizationOption struct {
	ID              uuid.UUID `db:"id"`
	GroupID         uuid.UUID `db:"group_id"`
	Label           string    `db:"label"`
	AdditionalPrice int64     `db:"additional_price"`
	IsAvailable     bool      `db:"is_available"`
	SortOrder       int       `db:"sort_order"`
	CreatedAt       time.Time `db:"created_at"`
}

type MenuItemCustomization struct {
	ID         uuid.UUID `db:"id"`
	MenuItemID uuid.UUID `db:"menu_item_id"`
	GroupID    uuid.UUID `db:"group_id"`
	IsDefault  bool      `db:"is_default"`
}