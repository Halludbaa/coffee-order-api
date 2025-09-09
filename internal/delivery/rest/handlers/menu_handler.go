package handlers

import (
	"coffee/internal/entity"
	"coffee/internal/model"
	"coffee/internal/model/converter"
	"coffee/internal/model/dto"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type MenuHandler struct {
    Service model.MenuService
    Log     *logrus.Entry
}

func NewMenuHandler(service model.MenuService, logger *logrus.Logger) *MenuHandler {
    return &MenuHandler{
        Service: service,
        Log:     logger.WithField("layer", "handler").WithField("struct", "MenuHandler"),
    }
}

// CreateMenuItem handles POST /api/menu
func (h *MenuHandler) CreateMenuItem(c *fiber.Ctx) error {
    h.Log.WithField("method", "CreateMenuItem").Info("Received create menu item request")

    var req dto.MenuItemRequest
    if err := c.BodyParser(&req); err != nil {
        h.Log.WithError(err).Warn("Invalid request body")
        return c.Status(http.StatusBadRequest).JSON(dto.WebErrorResponse{
            Error: "Invalid request body",
        })
    }

    item := converter.ToMenuItemEntity(&req)

    if err := h.Service.CreateMenuItem(c.Context(), item); err != nil {
        h.Log.WithError(err).Error("Service failed to create menu item")
        return c.Status(http.StatusInternalServerError).JSON(dto.WebErrorResponse{
            Error: err.Error(),
        })
    }

    resp := converter.ToMenuItemResponse(item)
    h.Log.WithField("menu_id", item.ID).Info("Menu item created and returned")

    return c.Status(http.StatusCreated).JSON(dto.WebSuccessResponse[dto.MenuItemResponse]{Data: resp})
}

// GetAllMenuItems handles GET /api/menu
func (h *MenuHandler) GetAllMenuItems(c *fiber.Ctx) error {
    h.Log.Info("Fetching all menu items")

    items, err := h.Service.GetAllMenuItems(c.Context())
    if err != nil {
        h.Log.WithError(err).Error("Failed to get all menu items")
        return c.Status(http.StatusInternalServerError).JSON(dto.WebErrorResponse{
            Error: "Failed to fetch menu items",
        })
    }

    responses := converter.ToMenuItemResponses(items)
    h.Log.WithField("count", len(responses)).Info("Returned all menu items")

    return c.JSON(dto.WebSuccessResponse[[]dto.MenuItemResponse]{Data: responses})
}

// GetMenuItemByID handles GET /api/menu/:id
func (h *MenuHandler) GetMenuItemByID(c *fiber.Ctx) error {
    id, err := uuid.Parse(c.Params("id"))
    if err != nil {
        return c.Status(http.StatusBadRequest).JSON(dto.WebErrorResponse{
            Error: "Invalid menu item ID",
        })
    }

    h.Log.WithField("menu_id", id).Info("Fetching menu item by ID")

    item, err := h.Service.GetMenuItemByID(c.Context(), id)
    if err != nil {
        h.Log.WithError(err).WithField("menu_id", id).Warn("Menu item not found")
        return c.Status(http.StatusNotFound).JSON(dto.WebErrorResponse{
            Error: "Menu item not found",
        })
    }

    resp := converter.ToMenuItemResponse(item)
    return c.JSON(dto.WebSuccessResponse[dto.MenuItemResponse]{Data: resp})
}

// UpdateMenuItem handles PUT /api/menu/:id
func (h *MenuHandler) UpdateMenuItem(c *fiber.Ctx) error {
    id, err := uuid.Parse(c.Params("id"))
    if err != nil {
        return c.Status(http.StatusBadRequest).JSON(dto.WebErrorResponse{
            Error: "Invalid menu item ID",
        })
    }

    var req dto.MenuItemRequest
    if err := c.BodyParser(&req); err != nil {
        return c.Status(http.StatusBadRequest).JSON(dto.WebErrorResponse{
            Error: "Invalid request body",
        })
    }

    h.Log.WithField("menu_id", id).Info("Updating menu item")

    // Fetch existing item
    item, err := h.Service.GetMenuItemByID(c.Context(), id)
    if err != nil {
        return c.Status(http.StatusNotFound).JSON(dto.WebErrorResponse{
            Error: "Menu item not found",
        })
    }

    // Update fields
    converter.UpdateMenuItemEntity(&req, item)

    if err := h.Service.UpdateMenuItem(c.Context(), item); err != nil {
        h.Log.WithError(err).Error("Failed to update menu item")
        return c.Status(http.StatusInternalServerError).JSON(dto.WebErrorResponse{
            Error: err.Error(),
        })
    }

    resp := converter.ToMenuItemResponse(item)
    return c.JSON(dto.WebSuccessResponse[dto.MenuItemResponse]{Data: resp})
}

// DeleteMenuItem handles DELETE /api/menu/:id
func (h *MenuHandler) DeleteMenuItem(c *fiber.Ctx) error {
    id, err := uuid.Parse(c.Params("id"))
    if err != nil {
        return c.Status(http.StatusBadRequest).JSON(dto.WebErrorResponse{
            Error: "Invalid menu item ID",
        })
    }

    h.Log.WithField("menu_id", id).Info("Deleting menu item")

    if err := h.Service.DeleteMenuItem(c.Context(), id); err != nil {
        h.Log.WithError(err).Error("Failed to delete menu item")
        return c.Status(http.StatusInternalServerError).JSON(dto.WebErrorResponse{
            Error: err.Error(),
        })
    }

    return c.JSON(dto.WebSuccessResponse[bool]{Data: true})
}

// AddToStoreMenu handles POST /api/stores/:storeId/menu
func (h *MenuHandler) AddToStoreMenu(c *fiber.Ctx) error {
    storeID, err := uuid.Parse(c.Params("storeId"))
    if err != nil {
        return c.Status(http.StatusBadRequest).JSON(dto.WebErrorResponse{
            Error: "Invalid store ID",
        })
    }

    var req dto.StoreMenuRequest
    if err := c.BodyParser(&req); err != nil {
        return c.Status(http.StatusBadRequest).JSON(dto.WebErrorResponse{
            Error: "Invalid request body",
        })
    }

    h.Log.WithFields(logrus.Fields{
        "store_id": storeID,
        "price_override": req.PriceOverride,
        "is_available": req.IsAvailable,
    }).Info("Adding menu item to store menu")

    item := &entity.StoreMenu{
        MenuItemID:    uuid.Nil, // Will be set from body or URL
        PriceOverride: req.PriceOverride,
        IsAvailable:   req.IsAvailable,
        SortOrder:     req.SortOrder,
    }

    // You can extend this to accept menu_item_id in body or URL
    menuItemID, err := uuid.Parse(c.Query("menu_item_id"))
    if err != nil {
        return c.Status(http.StatusBadRequest).JSON(dto.WebErrorResponse{
            Error: "Query param 'menu_item_id' is required",
        })
    }
    item.MenuItemID = menuItemID

    if err := h.Service.AddToStoreMenu(c.Context(), storeID, item); err != nil {
        h.Log.WithError(err).Error("Failed to add to store menu")
        return c.Status(http.StatusInternalServerError).JSON(dto.WebErrorResponse{
            Error: err.Error(),
        })
    }

    return c.JSON(dto.WebSuccessResponse[bool]{Data: true})
}

// GetStoreMenu handles GET /api/stores/:storeId/menu
func (h *MenuHandler) GetStoreMenu(c *fiber.Ctx) error {
    storeID, err := uuid.Parse(c.Params("storeId"))
    if err != nil {
        return c.Status(http.StatusBadRequest).JSON(dto.WebErrorResponse{
            Error: "Invalid store ID",
        })
    }

    h.Log.WithField("store_id", storeID).Info("Fetching store menu")

    items, err := h.Service.GetStoreMenu(c.Context(), storeID)
    if err != nil {
        h.Log.WithError(err).Error("Failed to get store menu")
        return c.Status(http.StatusInternalServerError).JSON(dto.WebErrorResponse{
            Error: "Failed to fetch store menu",
        })
    }

    responses := converter.ToMenuItemResponses(items)
    return c.JSON(dto.WebSuccessResponse[[]dto.MenuItemResponse]{Data: responses})
}

// UpdateStoreMenu handles PUT /api/stores/:storeId/menu/:menuItemId
func (h *MenuHandler) UpdateStoreMenu(c *fiber.Ctx) error {
    storeID, err := uuid.Parse(c.Params("storeId"))
    if err != nil {
        return c.Status(http.StatusBadRequest).JSON(dto.WebErrorResponse{
            Error: "Invalid store ID",
        })
    }

    menuItemID, err := uuid.Parse(c.Params("menuItemId"))
    if err != nil {
        return c.Status(http.StatusBadRequest).JSON(dto.WebErrorResponse{
            Error: "Invalid menu item ID",
        })
    }

    var req dto.StoreMenuRequest
    if err := c.BodyParser(&req); err != nil {
        return c.Status(http.StatusBadRequest).JSON(dto.WebErrorResponse{
            Error: "Invalid request body",
        })
    }

    h.Log.WithFields(logrus.Fields{
        "store_id": storeID,
        "menu_item_id": menuItemID,
    }).Info("Updating store menu item")

    item := &entity.StoreMenu{
        MenuItemID:    menuItemID,
        PriceOverride: req.PriceOverride,
        IsAvailable:   req.IsAvailable,
        SortOrder:     req.SortOrder,
    }

    if err := h.Service.UpdateStoreMenuItem(c.Context(), storeID, menuItemID, item); err != nil {
        h.Log.WithError(err).Error("Failed to update store menu item")
        return c.Status(http.StatusInternalServerError).JSON(dto.WebErrorResponse{
            Error: err.Error(),
        })
    }

    return c.JSON(dto.WebSuccessResponse[bool]{Data: true})
}

// RemoveFromStoreMenu handles DELETE /api/stores/:storeId/menu/:menuItemId
func (h *MenuHandler) DeleteStoreMenu(c *fiber.Ctx) error {
    storeID, err := uuid.Parse(c.Params("storeId"))
    if err != nil {
        return c.Status(http.StatusBadRequest).JSON(dto.WebErrorResponse{
            Error: "Invalid store ID",
        })
    }

    menuItemID, err := uuid.Parse(c.Params("menuItemId"))
    if err != nil {
        return c.Status(http.StatusBadRequest).JSON(dto.WebErrorResponse{
            Error: "Invalid menu item ID",
        })
    }

    h.Log.WithFields(logrus.Fields{
        "store_id": storeID,
        "menu_item_id": menuItemID,
    }).Info("Removing item from store menu")

    if err := h.Service.RemoveFromStoreMenu(c.Context(), storeID, menuItemID); err != nil {
        h.Log.WithError(err).Error("Failed to remove item from store menu")
        return c.Status(http.StatusInternalServerError).JSON(dto.WebErrorResponse{
            Error: err.Error(),
        })
    }

    return c.JSON(dto.WebSuccessResponse[bool]{Data: true})
}