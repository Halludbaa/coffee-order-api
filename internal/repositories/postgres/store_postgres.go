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

type StoreRepository struct {
    DB  *sqlx.DB
    Log *logrus.Entry
}

func NewStoreRepository(db *sqlx.DB, log *logrus.Logger) model.StoreRepository {
    return &StoreRepository{
        DB:  db,
        Log: log.WithField("layer", "repository").WithField("struct", "StoreRepository"),
    }
}

func (r *StoreRepository) CreateStore(ctx context.Context, store *entity.Store) error {
    r.Log.WithField("name", store.Name).Info("Creating new store")

    query := `
        INSERT INTO stores (id, name, location, address, phone, email, store_slug, is_active)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
        RETURNING created_at, updated_at
    `

    err := r.DB.QueryRowContext(ctx, query,
        store.ID,
        store.Name,
        store.Location,
        store.Address,
        store.Phone,
        store.Email,
        store.StoreSlug,
        store.IsActive,
    ).Scan(&store.CreatedAt, &store.UpdatedAt)

    if err != nil {
        r.Log.WithError(err).Error("Failed to create store")
        return fmt.Errorf("failed to create store: %w", err)
    }

    return nil
}

func (r *StoreRepository) GetAllStores(ctx context.Context) ([]*entity.Store, error) {
    r.Log.Info("Fetching all stores")

    query := `
        SELECT id, name, location, address, phone, email, store_slug, 
               is_active, created_at, updated_at
        FROM stores
        ORDER BY name
    `

    var stores []*entity.Store
    if err := r.DB.SelectContext(ctx, &stores, query); err != nil {
        r.Log.WithError(err).Error("Failed to fetch stores")
        return nil, fmt.Errorf("failed to fetch stores: %w", err)
    }

    return stores, nil
}

func (r *StoreRepository) GetStoreByID(ctx context.Context, id uuid.UUID) (*entity.Store, error) {
    r.Log.WithField("id", id).Info("Fetching store by ID")

    query := `
        SELECT id, name, location, address, phone, email, store_slug, 
               is_active, created_at, updated_at
        FROM stores
        WHERE id = $1
    `

    var store entity.Store
    if err := r.DB.GetContext(ctx, &store, query, id); err != nil {
        if err == sql.ErrNoRows {
            return nil, fmt.Errorf("store not found")
        }
        r.Log.WithError(err).Error("Failed to fetch store")
        return nil, fmt.Errorf("failed to fetch store: %w", err)
    }

    return &store, nil
}

func (r *StoreRepository) GetStoreBySlug(ctx context.Context, slug string) (*entity.Store, error) {
    r.Log.WithField("slug", slug).Info("Fetching store by slug")

    query := `
        SELECT id, name, location, address, phone, email, store_slug, 
               is_active, created_at, updated_at
        FROM stores
        WHERE store_slug = $1
    `

    var store entity.Store
    if err := r.DB.GetContext(ctx, &store, query, slug); err != nil {
        if err == sql.ErrNoRows {
            return nil, fmt.Errorf("store not found")
        }
        r.Log.WithError(err).Error("Failed to fetch store")
        return nil, fmt.Errorf("failed to fetch store: %w", err)
    }

    return &store, nil
}

func (r *StoreRepository) UpdateStore(ctx context.Context, store *entity.Store) error {
    r.Log.WithField("id", store.ID).Info("Updating store")

    query := `
        UPDATE stores
        SET name = $1, location = $2, address = $3, phone = $4, 
            email = $5, store_slug = $6, is_active = $7, updated_at = NOW()
        WHERE id = $8
        RETURNING updated_at
    `

    result, err := r.DB.ExecContext(ctx, query,
        store.Name,
        store.Location,
        store.Address,
        store.Phone,
        store.Email,
        store.StoreSlug,
        store.IsActive,
        store.ID,
    )

    if err != nil {
        r.Log.WithError(err).Error("Failed to update store")
        return fmt.Errorf("failed to update store: %w", err)
    }

    rows, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("error checking rows affected: %w", err)
    }
    if rows == 0 {
        return fmt.Errorf("store not found")
    }

    return nil
}

func (r *StoreRepository) DeleteStore(ctx context.Context, id uuid.UUID) error {
    r.Log.WithField("id", id).Info("Deleting store")

    query := "DELETE FROM stores WHERE id = $1"

    result, err := r.DB.ExecContext(ctx, query, id)
    if err != nil {
        r.Log.WithError(err).Error("Failed to delete store")
        return fmt.Errorf("failed to delete store: %w", err)
    }

    rows, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("error checking rows affected: %w", err)
    }
    if rows == 0 {
        return fmt.Errorf("store not found")
    }

    return nil
}

func (r *StoreRepository) GetActiveStores(ctx context.Context) ([]*entity.Store, error) {
    r.Log.Info("Fetching active stores")

    query := `
        SELECT id, name, location, address, phone, email, store_slug, 
               is_active, created_at, updated_at
        FROM stores
        WHERE is_active = true
        ORDER BY name
    `

    var stores []*entity.Store
    if err := r.DB.SelectContext(ctx, &stores, query); err != nil {
        r.Log.WithError(err).Error("Failed to fetch active stores")
        return nil, fmt.Errorf("failed to fetch active stores: %w", err)
    }

    return stores, nil
}

func (r *StoreRepository) SetStoreActive(ctx context.Context, id uuid.UUID, active bool) error {
    r.Log.WithFields(logrus.Fields{
        "id":     id,
        "active": active,
    }).Info("Setting store active status")

    query := `
        UPDATE stores
        SET is_active = $1, updated_at = NOW()
        WHERE id = $2
    `

    result, err := r.DB.ExecContext(ctx, query, active, id)
    if err != nil {
        r.Log.WithError(err).Error("Failed to update store active status")
        return fmt.Errorf("failed to update store active status: %w", err)
    }

    rows, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("error checking rows affected: %w", err)
    }
    if rows == 0 {
        return fmt.Errorf("store not found")
    }

    return nil
}