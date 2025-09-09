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