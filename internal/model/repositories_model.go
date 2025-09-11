package model

import (
	"coffee/internal/entity"
	"context"

	"github.com/google/uuid"
)

type UserRepository interface {
	// Store(ctx context.Context, request *SignUpRequest)  (*entity.User, error) 
	// Remove(ctx context.Context, username string) (error)
	// Update(ctx context.Context, userUpdateReq *UserUpdateRequest) (error)
	// FindByUsername(ctx context.Context, username string) (*entity.User, error)
	// FindById(ctx context.Context, Id string) (*entity.User, error)
	// FindByEmail(ctx context.Context, email string) (*entity.User, error)
}

type SessionRepo interface {
	Store(ctx context.Context, request *entity.Session) (error)
	Remove(ctx context.Context, request *entity.Session) (error)
	FindByUserId(ctx context.Context,  record *entity.Session) (error)
	FindByToken(ctx context.Context,  record *entity.Session) (error)
}

type MenuRepository interface {
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

type CategoryRepository interface {
    CreateCategory(ctx context.Context, category *entity.Category) error
    GetAllCategories(ctx context.Context) ([]*entity.Category, error)
    GetCategoryByID(ctx context.Context, id uuid.UUID) (*entity.Category, error)
    UpdateCategory(ctx context.Context, category *entity.Category) error
    DeleteCategory(ctx context.Context, id uuid.UUID) error
    
    // Store-specific category operations
    AddToStoreCategory(ctx context.Context, storeID uuid.UUID, category *entity.StoreCategory) error
    GetStoreCategories(ctx context.Context, storeID uuid.UUID) ([]*entity.Category, error)
    UpdateStoreCategory(ctx context.Context, storeID, categoryID uuid.UUID, category *entity.StoreCategory) error
    RemoveFromStoreCategory(ctx context.Context, storeID, categoryID uuid.UUID) error
}

type StoreRepository interface {
    CreateStore(ctx context.Context, store *entity.Store) error
    GetAllStores(ctx context.Context) ([]*entity.Store, error)
    GetStoreByID(ctx context.Context, id uuid.UUID) (*entity.Store, error)
    GetStoreBySlug(ctx context.Context, slug string) (*entity.Store, error)
    UpdateStore(ctx context.Context, store *entity.Store) error
    DeleteStore(ctx context.Context, id uuid.UUID) error
    
    // Optional: Add methods for active stores only
    GetActiveStores(ctx context.Context) ([]*entity.Store, error)
    SetStoreActive(ctx context.Context, id uuid.UUID, active bool) error
}

type CustomizationRepository interface {
    // Group operations
    CreateGroup(ctx context.Context, group *entity.CustomizationGroup) error
    GetGroupsByStore(ctx context.Context, storeID uuid.UUID) ([]*entity.CustomizationGroup, error)
    UpdateGroup(ctx context.Context, group *entity.CustomizationGroup) error
    DeleteGroup(ctx context.Context, id uuid.UUID) error

    // Option operations
    CreateOption(ctx context.Context, option *entity.CustomizationOption) error
    GetOptionsByGroup(ctx context.Context, groupID uuid.UUID) ([]*entity.CustomizationOption, error)
    UpdateOption(ctx context.Context, option *entity.CustomizationOption) error
    DeleteOption(ctx context.Context, id uuid.UUID) error

    // MenuItem customization operations
    AssignGroupToMenuItem(ctx context.Context, menuItemID, groupID uuid.UUID, isDefault bool) error
    GetMenuItemCustomizations(ctx context.Context, menuItemID uuid.UUID) ([]*entity.CustomizationGroup, error)
    RemoveGroupFromMenuItem(ctx context.Context, menuItemID, groupID uuid.UUID) error
}