package postgres

import (
	"coffee/internal/entity"
	"coffee/internal/model"
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type CategoryRepository struct {
    DB  *sqlx.DB
    Log *logrus.Entry
}

func NewCategoryRepository(db *sqlx.DB, log *logrus.Logger) model.CategoryRepository {
    return &CategoryRepository{
        DB:  db,
        Log: log.WithField("layer", "repository").WithField("struct", "CategoryRepository"),
    }
}

func (r *CategoryRepository) CreateCategory(ctx context.Context, category *entity.Category) error {
    r.Log.WithField("name", category.Name).Info("Creating new category")

    query := `
        INSERT INTO categories (id, name, icon, color, is_active)
        VALUES ($1, $2, $3, $4, $5)
    `

    _, err := r.DB.ExecContext(ctx, query,
        category.ID, category.Name, category.Icon,
        category.Color, category.IsActive,
    )

    if err != nil {
        r.Log.WithError(err).Error("Failed to create category")
        return fmt.Errorf("failed to create category: %w", err)
    }

    return nil
}

func (r *CategoryRepository) GetAllCategories(ctx context.Context) ([]*entity.Category, error) {
    r.Log.Info("Fetching all categories")

    query := `
        SELECT id, name, icon, color, is_active
        FROM categories
        ORDER BY name
    `

    var categories []*entity.Category
    err := r.DB.SelectContext(ctx, &categories, query)
    if err != nil {
        r.Log.WithError(err).Error("Failed to fetch categories")
        return nil, fmt.Errorf("failed to fetch categories: %w", err)
    }

    return categories, nil
}

func (r *CategoryRepository) GetCategoryByID(ctx context.Context, id uuid.UUID) (*entity.Category, error) {
    r.Log.WithField("category_id", id).Info("Fetching category by ID")

    query := `
        SELECT id, name, icon, color, is_active
        FROM categories
        WHERE id = $1
    `

    var category entity.Category
    err := r.DB.GetContext(ctx, &category, query, id)
    if err != nil {
        r.Log.WithError(err).Error("Failed to fetch category")
        return nil, fmt.Errorf("failed to fetch category: %w", err)
    }

    return &category, nil
}

func (r *CategoryRepository) UpdateCategory(ctx context.Context, category *entity.Category) error {
    r.Log.WithField("category_id", category.ID).Info("Updating category")

    query := `
        UPDATE categories
        SET name = $1, icon = $2, color = $3, is_active = $4
        WHERE id = $5
    `

    result, err := r.DB.ExecContext(ctx, query,
        category.Name, category.Icon, category.Color,
        category.IsActive, category.ID,
    )
    if err != nil {
        r.Log.WithError(err).Error("Failed to update category")
        return fmt.Errorf("failed to update category: %w", err)
    }

    rows, _ := result.RowsAffected()
    if rows == 0 {
        return fmt.Errorf("category not found")
    }

    return nil
}

func (r *CategoryRepository) DeleteCategory(ctx context.Context, id uuid.UUID) error {
    r.Log.WithField("category_id", id).Info("Deleting category")

    query := "DELETE FROM categories WHERE id = $1"

    result, err := r.DB.ExecContext(ctx, query, id)
    if err != nil {
        r.Log.WithError(err).Error("Failed to delete category")
        return fmt.Errorf("failed to delete category: %w", err)
    }

    rows, _ := result.RowsAffected()
    if rows == 0 {
        return fmt.Errorf("category not found")
    }

    return nil
}

func (r *CategoryRepository) AddToStoreCategory(ctx context.Context, storeID uuid.UUID, category *entity.StoreCategory) error {
    r.Log.WithFields(logrus.Fields{
        "store_id": storeID,
        "category_id": category.CategoryID,
    }).Info("Adding category to store")

    query := `
        INSERT INTO store_categories (store_id, category_id, name, is_visible, sort_order)
        VALUES ($1, $2, $3, $4, $5)
        ON CONFLICT (store_id, category_id) DO UPDATE
        SET name = EXCLUtDED.name,
            is_visible = EXCLUDED.is_visible,
            sort_order = EXCLUDED.sort_order
    `

    _, err := r.DB.ExecContext(ctx, query,
        storeID, category.CategoryID, category.Name,
        category.IsVisible, category.SortOrder,
    )

    if err != nil {
        r.Log.WithError(err).Error("Failed to add category to store")
        return fmt.Errorf("failed to add category to store: %w", err)
    }

    return nil
}

func (r *CategoryRepository) GetStoreCategories(ctx context.Context, storeID uuid.UUID) ([]*entity.Category, error) {
    r.Log.WithField("store_id", storeID).Info("Fetching store categories")

    query := `
        SELECT 
            c.id, COALESCE(sc.name, c.name) as name,
            c.icon, c.color, c.is_active
        FROM store_categories sc
        JOIN categories c ON c.id = sc.category_id
        WHERE sc.store_id = $1 AND sc.is_visible = true
        ORDER BY sc.sort_order
    `

    var categories []*entity.Category
    err := r.DB.SelectContext(ctx, &categories, query, storeID)
    if err != nil {
        r.Log.WithError(err).Error("Failed to fetch store categories")
        return nil, fmt.Errorf("failed to fetch store categories: %w", err)
    }

    return categories, nil
}

func (r *CategoryRepository) UpdateStoreCategory(ctx context.Context, storeID, categoryID uuid.UUID, category *entity.StoreCategory) error {
    r.Log.WithFields(logrus.Fields{
        "store_id": storeID,
        "category_id": categoryID,
    }).Info("Updating store category")

    query := `
        UPDATE store_categories
        SET name = $1, is_visible = $2, sort_order = $3
        WHERE store_id = $4 AND category_id = $5
    `

    result, err := r.DB.ExecContext(ctx, query,
        category.Name, category.IsVisible, category.SortOrder,
        storeID, categoryID,
    )

    if err != nil {
        r.Log.WithError(err).Error("Failed to update store category")
        return fmt.Errorf("failed to update store category: %w", err)
    }

    rows, _ := result.RowsAffected()
    if rows == 0 {
        return fmt.Errorf("store category not found")
    }

    return nil
}

func (r *CategoryRepository) RemoveFromStoreCategory(ctx context.Context, storeID, categoryID uuid.UUID) error {
    r.Log.WithFields(logrus.Fields{
        "store_id": storeID,
        "category_id": categoryID,
    }).Info("Removing category from store")

    query := `DELETE FROM store_categories WHERE store_id = $1 AND category_id = $2`

    result, err := r.DB.ExecContext(ctx, query, storeID, categoryID)
    if err != nil {
        r.Log.WithError(err).Error("Failed to remove category from store")
        return fmt.Errorf("failed to remove category from store: %w", err)
    }

    rows, _ := result.RowsAffected()
    if rows == 0 {
        return fmt.Errorf("store category not found")
    }

    return nil
}