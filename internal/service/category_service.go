package service

import (
	"coffee/internal/entity"
	"coffee/internal/model"
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type categoryService struct {
    repo model.CategoryRepository
    log  *logrus.Entry
}

func NewCategoryService(repo model.CategoryRepository, logger *logrus.Logger) model.CategoryService {
    return &categoryService{
        repo: repo,
        log:  logger.WithField("layer", "service").WithField("struct", "categoryService"),
    }
}

func (s *categoryService) CreateCategory(ctx context.Context, category *entity.Category) error {
    s.log.WithField("name", category.Name).Info("Creating new category")

    if category.ID == uuid.Nil {
        category.ID = uuid.New()
    }

    if category.Name == "" {
        return fmt.Errorf("category name is required")
    }

    return s.repo.CreateCategory(ctx, category)
}

func (s *categoryService) GetAllCategories(ctx context.Context) ([]*entity.Category, error) {
    s.log.Info("Fetching all categories")
    return s.repo.GetAllCategories(ctx)
}

func (s *categoryService) GetCategoryByID(ctx context.Context, id uuid.UUID) (*entity.Category, error) {
    s.log.WithField("category_id", id).Info("Fetching category by ID")
    
    if id == uuid.Nil {
        return nil, fmt.Errorf("invalid category ID")
    }
    
    return s.repo.GetCategoryByID(ctx, id)
}

func (s *categoryService) UpdateCategory(ctx context.Context, category *entity.Category) error {
    s.log.WithField("category_id", category.ID).Info("Updating category")

    if category.ID == uuid.Nil {
        return fmt.Errorf("invalid category ID")
    }

    if category.Name == "" {
        return fmt.Errorf("category name is required")
    }

    // Check if category exists
    existing, err := s.repo.GetCategoryByID(ctx, category.ID)
    if err != nil {
        return fmt.Errorf("category not found: %w", err)
    }

    // Update fields
    existing.Name = category.Name
    existing.Icon = category.Icon
    existing.Color = category.Color
    existing.IsActive = category.IsActive

    return s.repo.UpdateCategory(ctx, existing)
}

func (s *categoryService) DeleteCategory(ctx context.Context, id uuid.UUID) error {
    s.log.WithField("category_id", id).Info("Deleting category")

    if id == uuid.Nil {
        return fmt.Errorf("invalid category ID")
    }

    return s.repo.DeleteCategory(ctx, id)
}

func (s *categoryService) AddToStoreCategory(ctx context.Context, storeID uuid.UUID, category *entity.StoreCategory) error {
    s.log.WithFields(logrus.Fields{
        "store_id": storeID,
        "category_id": category.CategoryID,
    }).Info("Adding category to store")

    if storeID == uuid.Nil {
        return fmt.Errorf("invalid store ID")
    }

    if category.CategoryID == uuid.Nil {
        return fmt.Errorf("invalid category ID")
    }

    // Verify category exists
    _, err := s.repo.GetCategoryByID(ctx, category.CategoryID)
    if err != nil {
        return fmt.Errorf("category not found: %w", err)
    }

    return s.repo.AddToStoreCategory(ctx, storeID, category)
}

func (s *categoryService) GetStoreCategories(ctx context.Context, storeID uuid.UUID) ([]*entity.Category, error) {
    s.log.WithField("store_id", storeID).Info("Fetching store categories")

    if storeID == uuid.Nil {
        return nil, fmt.Errorf("invalid store ID")
    }

    return s.repo.GetStoreCategories(ctx, storeID)
}

func (s *categoryService) UpdateStoreCategory(ctx context.Context, storeID, categoryID uuid.UUID, category *entity.StoreCategory) error {
    s.log.WithFields(logrus.Fields{
        "store_id": storeID,
        "category_id": categoryID,
    }).Info("Updating store category")

    if storeID == uuid.Nil {
        return fmt.Errorf("invalid store ID")
    }

    if categoryID == uuid.Nil {
        return fmt.Errorf("invalid category ID")
    }

    category.StoreID = storeID
    category.CategoryID = categoryID

    return s.repo.UpdateStoreCategory(ctx, storeID, categoryID, category)
}

func (s *categoryService) RemoveFromStoreCategory(ctx context.Context, storeID, categoryID uuid.UUID) error {
    s.log.WithFields(logrus.Fields{
        "store_id": storeID,
        "category_id": categoryID,
    }).Info("Removing category from store")

    if storeID == uuid.Nil {
        return fmt.Errorf("invalid store ID")
    }

    if categoryID == uuid.Nil {
        return fmt.Errorf("invalid category ID")
    }

    return s.repo.RemoveFromStoreCategory(ctx, storeID, categoryID)
}