package service

import (
	"coffee/internal/entity"
	"coffee/internal/model"
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type storeService struct {
    repo model.StoreRepository
    log  *logrus.Entry
}

func NewStoreService(repo model.StoreRepository, logger *logrus.Logger) model.StoreService {
    return &storeService{
        repo: repo,
        log:  logger.WithField("layer", "service").WithField("struct", "storeService"),
    }
}

func (s *storeService) CreateStore(ctx context.Context, store *entity.Store) error {
    s.log.WithField("name", store.Name).Info("Creating new store")

    if err := s.validateStore(store); err != nil {
        return err
    }

    // Generate UUID if not provided
    if store.ID == uuid.Nil {
        store.ID = uuid.New()
    }

    // Generate slug if not provided
    if store.StoreSlug == "" {
        store.StoreSlug = s.generateSlug(store.Name)
    }

    // Set default active status
    if !store.IsActive {
        store.IsActive = true
    }

    return s.repo.CreateStore(ctx, store)
}

func (s *storeService) GetAllStores(ctx context.Context) ([]*entity.Store, error) {
    return s.repo.GetAllStores(ctx)
}

func (s *storeService) GetStoreByID(ctx context.Context, id uuid.UUID) (*entity.Store, error) {
    if id == uuid.Nil {
        return nil, fmt.Errorf("invalid store ID")
    }
    return s.repo.GetStoreByID(ctx, id)
}

func (s *storeService) GetStoreBySlug(ctx context.Context, slug string) (*entity.Store, error) {
    if slug == "" {
        return nil, fmt.Errorf("store slug cannot be empty")
    }
    return s.repo.GetStoreBySlug(ctx, slug)
}

func (s *storeService) UpdateStore(ctx context.Context, store *entity.Store) error {
    s.log.WithField("id", store.ID).Info("Updating store")

    if store.ID == uuid.Nil {
        return fmt.Errorf("invalid store ID")
    }

    if err := s.validateStore(store); err != nil {
        return err
    }

    // Check if store exists
    existing, err := s.repo.GetStoreByID(ctx, store.ID)
    if err != nil {
        return err
    }

    // Preserve certain fields from existing store
    store.CreatedAt = existing.CreatedAt
    
    return s.repo.UpdateStore(ctx, store)
}

func (s *storeService) DeleteStore(ctx context.Context, id uuid.UUID) error {
    if id == uuid.Nil {
        return fmt.Errorf("invalid store ID")
    }
    return s.repo.DeleteStore(ctx, id)
}

func (s *storeService) GetActiveStores(ctx context.Context) ([]*entity.Store, error) {
    return s.repo.GetActiveStores(ctx)
}

func (s *storeService) SetStoreActive(ctx context.Context, id uuid.UUID, active bool) error {
    if id == uuid.Nil {
        return fmt.Errorf("invalid store ID")
    }
    return s.repo.SetStoreActive(ctx, id, active)
}

// Helper functions

func (s *storeService) validateStore(store *entity.Store) error {
    if store.Name == "" {
        return fmt.Errorf("store name is required")
    }

    if len(store.Name) > 100 {
        return fmt.Errorf("store name cannot exceed 100 characters")
    }

    if store.Email != "" && len(store.Email) > 100 {
        return fmt.Errorf("email cannot exceed 100 characters")
    }

    if store.Phone != "" && len(store.Phone) > 20 {
        return fmt.Errorf("phone number cannot exceed 20 characters")
    }

    return nil
}

func (s *storeService) generateSlug(name string) string {
    // Convert to lowercase
    slug := strings.ToLower(name)
    
    // Replace spaces with hyphens
    slug = strings.ReplaceAll(slug, " ", "-")
    
    // Remove special characters
    slug = strings.Map(func(r rune) rune {
        if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '-' {
            return r
        }
        return -1
    }, slug)
    
    // Remove multiple consecutive hyphens
    for strings.Contains(slug, "--") {
        slug = strings.ReplaceAll(slug, "--", "-")
    }
    
    // Trim hyphens from start and end
    slug = strings.Trim(slug, "-")
    
    return slug
}