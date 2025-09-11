package model

import (
	"coffee/internal/entity"
	"coffee/internal/model/apperrors"
	"context"

	"github.com/google/uuid"
)

type MenuService interface {
    // Global Menu Item
    CreateMenuItem(ctx context.Context, item *entity.MenuItem) error
    GetAllMenuItems(ctx context.Context) ([]*entity.MenuItem, error)
    GetMenuItemByID(ctx context.Context, id uuid.UUID) (*entity.MenuItem, error)
    UpdateMenuItem(ctx context.Context, item *entity.MenuItem) error
    DeleteMenuItem(ctx context.Context, id uuid.UUID) error

    // Store-Specific Menu
    AddToStoreMenu(ctx context.Context, storeID uuid.UUID, item *entity.StoreMenu) error
    GetStoreMenu(ctx context.Context, storeID uuid.UUID) ([]*entity.MenuItem, error)
    UpdateStoreMenuItem(ctx context.Context, storeID, menuItemID uuid.UUID, item *entity.StoreMenu) error
    RemoveFromStoreMenu(ctx context.Context, storeID, menuItemID uuid.UUID) error
}


type JWTServices interface{
	GenerateAccessToken(userID string) (string, error)
	GenerateRefreshToken(userID string) (string, error)
	ValidateAccessToken(tokenString string) (string, *apperrors.Apperrors)
	ValidateRefreshToken(tokenString string) (string, *apperrors.Apperrors)
}

type CategoryService interface {
    // Global Category
    CreateCategory(ctx context.Context, category *entity.Category) error
    GetAllCategories(ctx context.Context) ([]*entity.Category, error)
    GetCategoryByID(ctx context.Context, id uuid.UUID) (*entity.Category, error)
    UpdateCategory(ctx context.Context, category *entity.Category) error
    DeleteCategory(ctx context.Context, id uuid.UUID) error

    // Store-Specific Category
    AddToStoreCategory(ctx context.Context, storeID uuid.UUID, category *entity.StoreCategory) error
    GetStoreCategories(ctx context.Context, storeID uuid.UUID) ([]*entity.Category, error)
    UpdateStoreCategory(ctx context.Context, storeID, categoryID uuid.UUID, category *entity.StoreCategory) error
    RemoveFromStoreCategory(ctx context.Context, storeID, categoryID uuid.UUID) error
}

type StoreService interface {
    CreateStore(ctx context.Context, store *entity.Store) error
    GetAllStores(ctx context.Context) ([]*entity.Store, error)
    GetStoreByID(ctx context.Context, id uuid.UUID) (*entity.Store, error)
    GetStoreBySlug(ctx context.Context, slug string) (*entity.Store, error)
    UpdateStore(ctx context.Context, store *entity.Store) error
    DeleteStore(ctx context.Context, id uuid.UUID) error
    GetActiveStores(ctx context.Context) ([]*entity.Store, error)
    SetStoreActive(ctx context.Context, id uuid.UUID, active bool) error
}

type CustomizationService interface {
    CreateCustomizationGroup(ctx context.Context, group *entity.CustomizationGroup) error
    GetStoreCustomizations(ctx context.Context, storeID uuid.UUID) ([]*entity.CustomizationGroup, error)
    UpdateCustomizationGroup(ctx context.Context, group *entity.CustomizationGroup) error
    DeleteCustomizationGroup(ctx context.Context, id uuid.UUID) error
    
    AddCustomizationOption(ctx context.Context, option *entity.CustomizationOption) error
    UpdateCustomizationOption(ctx context.Context, option *entity.CustomizationOption) error
    DeleteCustomizationOption(ctx context.Context, id uuid.UUID) error
    
    AssignCustomizationToMenuItem(ctx context.Context, menuItemID, groupID uuid.UUID, isDefault bool) error
    GetMenuItemCustomizations(ctx context.Context, menuItemID uuid.UUID) ([]*entity.CustomizationGroup, error)
}
