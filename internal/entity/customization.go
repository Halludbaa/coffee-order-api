package entity

import "time"

type CustomizationGroup struct {
	ID         int       `db:"id" json:"id"`
	Name       string    `db:"name" json:"name"`
	StoreID    int       `db:"store_id" json:"store_id"`
	IsRequired bool      `db:"is_required" json:"is_required"`
	SortOrder  int       `db:"sort_order" json:"sort_order"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
}

type CustomizationOption struct {
	ID                 int       `db:"id" json:"id"`
	GroupID            int       `db:"group_id" json:"group_id"`
	Label              string    `db:"label" json:"label"`
	AdditionalPrice    int64     `db:"additional_price" json:"additional_price"`
	IsAvailable        bool      `db:"is_available" json:"is_available"`
	SortOrder          int       `db:"sort_order" json:"sort_order"`
	CreatedAt          time.Time `db:"created_at" json:"created_at"`
}

type MenuItemCustomization struct {
	ID           int `db:"id" json:"id"`
	MenuItemID   int `db:"menu_item_id" json:"menu_item_id"`
	GroupID      int `db:"group_id" json:"group_id"`
	IsDefault    bool `db:"is_default" json:"is_default"`
}