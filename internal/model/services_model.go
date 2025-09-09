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
