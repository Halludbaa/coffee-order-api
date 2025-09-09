package entity

import "time"

type Order struct {
	ID            int       `db:"id" json:"id"`
	StoreID       int       `db:"store_id" json:"store_id"`
	OrderNumber   string    `db:"order_number" json:"order_number"`
	Status        string    `db:"status" json:"status"` // pending, preparing, ready, etc.
	Total         int64     `db:"total" json:"total"`   // IDR
	CustomerNote  string    `db:"customer_note" json:"customer_note,omitempty"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time `db:"updated_at" json:"updated_at"`
}

type OrderItem struct {
	ID             int         `db:"id" json:"id"`
	OrderID        int         `db:"order_id" json:"order_id"`
	MenuItemID     int         `db:"menu_item_id" json:"menu_item_id"`
	Quantity       int         `db:"quantity" json:"quantity"`
	UnitPrice      int64       `db:"unit_price" json:"unit_price"`
	Customizations interface{} `db:"customizations" json:"customizations,omitempty"` // JSONB
	Note           string      `db:"note" json:"note,omitempty"`
	CreatedAt      time.Time   `db:"created_at" json:"created_at"`
}