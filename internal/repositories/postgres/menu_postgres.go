package postgres

import (
	"coffee/internal/entity"
	"coffee/internal/model"
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type MenuRepository struct {
	DB *sqlx.DB
	Log *logrus.Entry
}

func NewMenuRepository(db *sqlx.DB, log *logrus.Logger) model.MenuRepository {
	return &MenuRepository{
		DB: db,
		Log: log.WithField("layer", "repository").WithField("struct", "MenuRepository"),
	}
}
func (r *MenuRepository) CreateMenuItem(ctx context.Context, item *entity.MenuItem) error {
    r.Log.WithFields(logrus.Fields{
        "method": "CreateMenuItem",
        "menu_id": item.ID,
        "name":    item.Name,
        "price":   item.BasePrice,
    }).Info("Creating new menu item")

    query := `
        INSERT INTO menu_items (id, name, description, base_price, category_id, is_active, image_url)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
    `

    _, err := r.DB.ExecContext(ctx, query,
        item.ID, item.Name, item.Description, item.BasePrice,
        item.CategoryID, item.IsActive, item.ImageURL,
    )

    if err != nil {
        r.Log.WithError(err).WithField("menu_id", item.ID).Error("Failed to create menu item")
        return fmt.Errorf("failed to create menu item: %w", err)
    }

    r.Log.WithField("menu_id", item.ID).Info("Menu item created successfully")
    return nil
}

func (r *MenuRepository) GetAllMenuItems(ctx context.Context) ([]*entity.MenuItem, error) {
    r.Log.WithField("method", "GetAllMenuItems").Info("Fetching all menu items")

    query := `
        SELECT id, name, description, base_price, category_id, is_active, image_url
        FROM menu_items
        ORDER BY name
    `

    rows, err := r.DB.QueryContext(ctx, query)
    if err != nil {
        r.Log.WithError(err).Error("Failed to query all menu items")
        return nil, fmt.Errorf("failed to query menu items: %w", err)
    }
    defer rows.Close()

    var items []*entity.MenuItem
    for rows.Next() {
        var item entity.MenuItem
        err := rows.Scan(
            &item.ID, &item.Name, &item.Description, &item.BasePrice,
            &item.CategoryID, &item.IsActive, &item.ImageURL,
        )
        if err != nil {
            r.Log.WithError(err).Error("Failed to scan menu item")
            continue
        }
        items = append(items, &item)
    }

    if err = rows.Err(); err != nil {
        r.Log.WithError(err).Error("Row iteration error")
        return nil, fmt.Errorf("row iteration error: %w", err)
    }

    r.Log.WithField("count", len(items)).Info("Fetched all menu items")
    return items, nil
}

func (r *MenuRepository) GetMenuItemByID(ctx context.Context, id uuid.UUID) (*entity.MenuItem, error) {
    r.Log.WithFields(logrus.Fields{
        "method": "GetMenuItemByID",
        "menu_id": id,
    }).Info("Fetching menu item by ID")

    var item entity.MenuItem
    query := `
        SELECT id, name, description, base_price, category_id, is_active, image_url
        FROM menu_items
        WHERE id = $1
    `

    err := r.DB.QueryRowContext(ctx, query, id).Scan(
        &item.ID, &item.Name, &item.Description, &item.BasePrice,
        &item.CategoryID, &item.IsActive, &item.ImageURL,
    )

    if err == sql.ErrNoRows {
        r.Log.WithField("menu_id", id).Warn("Menu item not found")
        return nil, fmt.Errorf("menu item not found")
    } else if err != nil {
        r.Log.WithError(err).WithField("menu_id", id).Error("Database error")
        return nil, fmt.Errorf("database error: %w", err)
    }

    r.Log.WithField("menu_id", id).Info("Menu item found")
    return &item, nil
}

func (r *MenuRepository) UpdateMenuItem(ctx context.Context, item *entity.MenuItem) error {
    r.Log.WithFields(logrus.Fields{
        "method": "UpdateMenuItem",
        "menu_id": item.ID,
        "name":    item.Name,
        "price":   item.BasePrice,
    }).Info("Updating menu item")

    query := `
        UPDATE menu_items
        SET name = $1, description = $2, base_price = $3,
            category_id = $4, is_active = $5, image_url = $6
        WHERE id = $7
    `

    result, err := r.DB.ExecContext(ctx, query,
        item.Name, item.Description, item.BasePrice,
        item.CategoryID, item.IsActive, item.ImageURL, item.ID,
    )

    if err != nil {
        r.Log.WithError(err).WithField("menu_id", item.ID).Error("Failed to update menu item")
        return fmt.Errorf("failed to update menu item: %w", err)
    }

    rows, _ := result.RowsAffected()
    if rows == 0 {
        r.Log.WithField("menu_id", item.ID).Warn("No rows affected — menu item may not exist")
    }

    r.Log.WithField("menu_id", item.ID).Info("Menu item updated")
    return nil
}

func (r *MenuRepository) DeleteMenuItem(ctx context.Context, id uuid.UUID) error {
    r.Log.WithFields(logrus.Fields{
        "method": "DeleteMenuItem",
        "menu_id": id,
    }).Info("Deleting menu item")

    result, err := r.DB.ExecContext(ctx, "DELETE FROM menu_items WHERE id = $1", id)
    if err != nil {
        r.Log.WithError(err).WithField("menu_id", id).Error("Failed to delete menu item")
        return fmt.Errorf("failed to delete menu item: %w", err)
    }

    rows, _ := result.RowsAffected()
    if rows == 0 {
        r.Log.WithField("menu_id", id).Warn("No rows deleted — menu item may not exist")
    }

    r.Log.WithField("menu_id", id).Info("Menu item deleted")
    return nil
}

func (r *MenuRepository) AddToStoreMenu(ctx context.Context, storeID uuid.UUID, item *entity.StoreMenu) error {
    r.Log.WithFields(logrus.Fields{
        "method": "AddToStoreMenu",
        "store_id": storeID,
        "menu_item_id": item.MenuItemID,
        "price_override": item.PriceOverride,
        "is_available": item.IsAvailable,
    }).Info("Adding menu item to store menu")

    query := `
        INSERT INTO store_menu (store_id, menu_item_id, price_override, is_available, sort_order)
        VALUES ($1, $2, $3, $4, $5)
        ON CONFLICT (store_id, menu_item_id) DO UPDATE
        SET price_override = EXCLUDED.price_override,
            is_available = EXCLUDED.is_available,
            sort_order = EXCLUDED.sort_order
    `

    _, err := r.DB.ExecContext(ctx, query,
        storeID, item.MenuItemID, item.PriceOverride,
        item.IsAvailable, item.SortOrder,
    )

    if err != nil {
        r.Log.WithError(err).WithFields(logrus.Fields{
            "store_id":     storeID,
            "menu_item_id": item.MenuItemID,
        }).Error("Failed to add item to store menu")
        return fmt.Errorf("failed to add to store menu: %w", err)
    }

    r.Log.WithFields(logrus.Fields{
        "store_id":     storeID,
        "menu_item_id": item.MenuItemID,
    }).Info("Menu item added to store menu")
    return nil
}

func (r *MenuRepository) GetStoreMenu(ctx context.Context, storeID uuid.UUID) ([]*entity.MenuItem, error) {
    r.Log.WithFields(logrus.Fields{
        "method": "GetStoreMenu",
        "store_id": storeID,
    }).Info("Fetching store-specific menu")

    query := `
        SELECT 
            mi.id, mi.name, mi.description, mi.image_url,
            COALESCE(sm.price_override, mi.base_price) AS base_price,
            mi.category_id, mi.is_active
        FROM store_menu sm
        JOIN menu_items mi ON mi.id = sm.menu_item_id
        WHERE sm.store_id = $1 AND sm.is_available = true AND mi.is_active = true
        ORDER BY sm.sort_order
    `

    rows, err := r.DB.QueryContext(ctx, query, storeID)
    if err != nil {
        r.Log.WithError(err).WithField("store_id", storeID).Error("Failed to query store menu")
        return nil, fmt.Errorf("failed to get store menu: %w", err)
    }
    defer rows.Close()

    var items []*entity.MenuItem
    for rows.Next() {
        var item entity.MenuItem
        var price int64
        err := rows.Scan(
            &item.ID, &item.Name, &item.Description, &item.ImageURL,
            &price, &item.CategoryID, &item.IsActive,
        )
        if err != nil {
            r.Log.WithError(err).Error("Failed to scan menu item row")
            continue
        }
        item.BasePrice = price
        items = append(items, &item)
    }

    if err = rows.Err(); err != nil {
        r.Log.WithError(err).Error("Row iteration error in store menu")
        return nil, fmt.Errorf("row iteration error: %w", err)
    }

    r.Log.WithFields(logrus.Fields{
        "store_id": storeID,
        "count":    len(items),
    }).Info("Store menu fetched successfully")
    return items, nil
}

func (r *MenuRepository) UpdateStoreMenuItem(ctx context.Context, storeID, menuItemID uuid.UUID, item *entity.StoreMenu) error {
    r.Log.WithFields(logrus.Fields{
        "method": "UpdateStoreMenuItem",
        "store_id": storeID,
        "menu_item_id": menuItemID,
        "price_override": item.PriceOverride,
        "is_available": item.IsAvailable,
    }).Info("Updating store-specific menu item")

    query := `
        UPDATE store_menu
        SET price_override = $1, is_available = $2, sort_order = $3
        WHERE store_id = $4 AND menu_item_id = $5
    `

    result, err := r.DB.ExecContext(ctx, query,
        item.PriceOverride, item.IsAvailable, item.SortOrder,
        storeID, menuItemID,
    )

    if err != nil {
        r.Log.WithError(err).WithFields(logrus.Fields{
            "store_id": storeID,
            "menu_item_id": menuItemID,
        }).Error("Failed to update store menu item")
        return fmt.Errorf("failed to update store menu item: %w", err)
    }

    rows, _ := result.RowsAffected()
    if rows == 0 {
        r.Log.WithFields(logrus.Fields{
            "store_id": storeID,
            "menu_item_id": menuItemID,
        }).Warn("No rows updated — item may not exist in store menu")
    }

    r.Log.WithFields(logrus.Fields{
        "store_id": storeID,
        "menu_item_id": menuItemID,
        "rows_affected": rows,
    }).Info("Store menu item updated")
    return nil
}

func (r *MenuRepository) RemoveFromStoreMenu(ctx context.Context, storeID, menuItemID uuid.UUID) error {
    r.Log.WithFields(logrus.Fields{
        "method": "RemoveFromStoreMenu",
        "store_id": storeID,
        "menu_item_id": menuItemID,
    }).Info("Removing menu item from store menu")

    result, err := r.DB.ExecContext(ctx, "DELETE FROM store_menu WHERE store_id = $1 AND menu_item_id = $2", storeID, menuItemID)
    if err != nil {
        r.Log.WithError(err).WithFields(logrus.Fields{
            "store_id": storeID,
            "menu_item_id": menuItemID,
        }).Error("Failed to remove from store menu")
        return fmt.Errorf("failed to remove from store menu: %w", err)
    }

    rows, _ := result.RowsAffected()
    if rows == 0 {
        r.Log.WithFields(logrus.Fields{
            "store_id": storeID,
            "menu_item_id": menuItemID,
        }).Warn("No rows deleted — item may not exist in store menu")
    }

    r.Log.WithFields(logrus.Fields{
        "store_id": storeID,
        "menu_item_id": menuItemID,
        "rows_affected": rows,
    }).Info("Menu item removed from store menu")
    return nil
}

