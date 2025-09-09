package services

import (
	"coffee/internal/entity"
	"coffee/internal/model"
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type MenuService struct {
    MenuRepo model.MenuRepository
    Log      *logrus.Entry
}

func NewMenuService(menuRepo model.MenuRepository, logger *logrus.Logger) model.MenuService {
    return &MenuService{
        MenuRepo: menuRepo,
        Log: logger.WithField("layer", "service").WithField("struct", "MenuService"),
    }
}

func (s *MenuService) CreateMenuItem(ctx context.Context, item *entity.MenuItem) error {
    s.Log.WithFields(logrus.Fields{
        "method": "CreateMenuItem",
        "name":   item.Name,
        "price":  item.BasePrice,
    }).Info("Starting to create menu item")

    if item.BasePrice <= 0 {
        s.Log.WithField("price", item.BasePrice).Warn("Invalid menu item price")
        return fmt.Errorf("base price must be greater than zero")
    }

    if item.Name == "" {
        s.Log.Warn("Menu item name is empty")
        return fmt.Errorf("name is required")
    }

    if err := s.MenuRepo.CreateMenuItem(ctx, item); err != nil {
        s.Log.WithError(err).Error("Failed to create menu item in repository")
        return err
    }

    s.Log.WithField("menu_id", item.ID).Info("Menu item created successfully")
    return nil
}

func (s *MenuService) GetAllMenuItems(ctx context.Context) ([]*entity.MenuItem, error) {
    s.Log.WithField("method", "GetAllMenuItems").Info("Fetching all menu items")

    items, err := s.MenuRepo.GetAllMenuItems(ctx)
    if err != nil {
        s.Log.WithError(err).Error("Failed to fetch menu items from repository")
        return nil, err
    }

    s.Log.WithField("count", len(items)).Info("Successfully fetched all menu items")
    return items, nil
}

func (s *MenuService) GetMenuItemByID(ctx context.Context, id uuid.UUID) (*entity.MenuItem, error) {
    s.Log.WithField("method", "GetMenuItemByID").WithField("menu_id", id).Info("Fetching menu item by ID")

    item, err := s.MenuRepo.GetMenuItemByID(ctx, id)
    if err != nil {
        s.Log.WithError(err).Error("Failed to get menu item")
        return nil, err
    }

    s.Log.WithField("menu_id", id).Info("Menu item retrieved")
    return item, nil
}

func (s *MenuService) UpdateMenuItem(ctx context.Context, item *entity.MenuItem) error {
    s.Log.WithFields(logrus.Fields{
        "method": "UpdateMenuItem",
        "menu_id": item.ID,
        "name":    item.Name,
    }).Info("Updating menu item")

    existing, err := s.MenuRepo.GetMenuItemByID(ctx, item.ID)
    if err != nil {
        s.Log.WithError(err).Warn("Menu item not found during update check")
        return fmt.Errorf("menu item not found: %w", err)
    }

    if !existing.IsActive && item.IsActive {
    }

    if item.BasePrice <= 0 {
        s.Log.WithField("price", item.BasePrice).Warn("Invalid updated price")
        return fmt.Errorf("base price must be greater than zero")
    }

    if err := s.MenuRepo.UpdateMenuItem(ctx, item); err != nil {
        s.Log.WithError(err).Error("Failed to update in repository")
        return err
    }

    s.Log.WithField("menu_id", item.ID).Info("Menu item updated successfully")
    return nil
}

func (s *MenuService) DeleteMenuItem(ctx context.Context, id uuid.UUID) error {
    s.Log.WithFields(logrus.Fields{
        "method": "DeleteMenuItem",
        "menu_id": id,
    }).Info("Deleting menu item")


    if err := s.MenuRepo.DeleteMenuItem(ctx, id); err != nil {
        s.Log.WithError(err).Error("Failed to delete menu item")
        return err
    }

    s.Log.WithField("menu_id", id).Info("Menu item deleted")
    return nil
}

func (s *MenuService) AddToStoreMenu(ctx context.Context, storeID uuid.UUID, item *entity.StoreMenu) error {
    s.Log.WithFields(logrus.Fields{
        "method": "AddToStoreMenu",
        "store_id": storeID,
        "menu_item_id": item.MenuItemID,
    }).Info("Adding menu item to store")

    _, err := s.MenuRepo.GetMenuItemByID(ctx, item.MenuItemID)
    if err != nil {
        s.Log.WithError(err).WithField("menu_item_id", item.MenuItemID).Warn("Menu item does not exist")
        return fmt.Errorf("menu item does not exist: %w", err)
    }

    if item.PriceOverride != nil && *item.PriceOverride < 0 {
        s.Log.WithField("price_override", *item.PriceOverride).Warn("Invalid price override")
        return fmt.Errorf("price override cannot be negative")
    }

    if err := s.MenuRepo.AddToStoreMenu(ctx, storeID, item); err != nil {
        s.Log.WithError(err).Error("Failed to add to store menu")
        return err
    }

    s.Log.WithFields(logrus.Fields{
        "store_id":     storeID,
        "menu_item_id": item.MenuItemID,
    }).Info("Menu item added to store")
    return nil
}

func (s *MenuService) GetStoreMenu(ctx context.Context, storeID uuid.UUID) ([]*entity.MenuItem, error) {
    s.Log.WithField("method", "GetStoreMenu").WithField("store_id", storeID).Info("Fetching store menu")

    items, err := s.MenuRepo.GetStoreMenu(ctx, storeID)
    if err != nil {
        s.Log.WithError(err).Error("Failed to get store menu from repository")
        return nil, err
    }

    s.Log.WithFields(logrus.Fields{
        "store_id": storeID,
        "count":    len(items),
    }).Info("Store menu retrieved successfully")
    return items, nil
}

func (s *MenuService) UpdateStoreMenuItem(ctx context.Context, storeID, menuItemID uuid.UUID, item *entity.StoreMenu) error {
    s.Log.WithFields(logrus.Fields{
        "method": "UpdateStoreMenuItem",
        "store_id": storeID,
        "menu_item_id": menuItemID,
    }).Info("Updating store menu item")

    storeMenuItems, err := s.MenuRepo.GetStoreMenu(ctx, storeID)
    if err != nil {
        s.Log.WithError(err).Error("Failed to verify store menu")
        return fmt.Errorf("failed to verify store context: %w", err)
    }

    found := false
    for _, i := range storeMenuItems {
        if i.ID == menuItemID {
            found = true
            break
        }
    }
    if !found {
        s.Log.WithFields(logrus.Fields{
            "store_id": storeID,
            "menu_item_id": menuItemID,
        }).Warn("Menu item not found in store")
        return fmt.Errorf("menu item not found in store menu")
    }

    if item.PriceOverride != nil && *item.PriceOverride < 0 {
        s.Log.WithField("price_override", *item.PriceOverride).Warn("Negative price override")
        return fmt.Errorf("price override cannot be negative")
    }

    if err := s.MenuRepo.UpdateStoreMenuItem(ctx, storeID, menuItemID, item); err != nil {
        s.Log.WithError(err).Error("Failed to update in repository")
        return err
    }

    s.Log.WithFields(logrus.Fields{
        "store_id": storeID,
        "menu_item_id": menuItemID,
    }).Info("Store menu item updated")
    return nil
}

func (s *MenuService) RemoveFromStoreMenu(ctx context.Context, storeID, menuItemID uuid.UUID) error {
    s.Log.WithFields(logrus.Fields{
        "method": "RemoveFromStoreMenu",
        "store_id": storeID,
        "menu_item_id": menuItemID,
    }).Info("Removing menu item from store menu")

    if err := s.MenuRepo.RemoveFromStoreMenu(ctx, storeID, menuItemID); err != nil {
        s.Log.WithError(err).Error("Failed to remove from store menu")
        return err
    }

    s.Log.WithFields(logrus.Fields{
        "store_id": storeID,
        "menu_item_id": menuItemID,
    }).Info("Menu item removed from store menu")
    return nil
}