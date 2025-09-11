package handlers

import (
	"coffee/internal/entity"
	"coffee/internal/model"
	"coffee/internal/model/dto"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type StoreHandler struct {
    Service model.StoreService
    Log     *logrus.Entry
}

func NewStoreHandler(service model.StoreService, logger *logrus.Logger) *StoreHandler {
    return &StoreHandler{
        Service: service,
        Log:     logger.WithField("layer", "handler").WithField("struct", "StoreHandler"),
    }
}

// CreateStore handles POST /api/stores
func (h *StoreHandler) CreateStore(c *fiber.Ctx) error {
    h.Log.Info("Handling create store request")

    var req dto.StoreRequest
    if err := c.BodyParser(&req); err != nil {
        return c.Status(http.StatusBadRequest).JSON(dto.WebErrorResponse{
            Error: "Invalid request body",
        })
    }

    store := &entity.Store{
        Name:      req.Name,
        Location:  req.Location,
        Address:   req.Address,
        Phone:     req.Phone,
        Email:     req.Email,
        StoreSlug: req.StoreSlug,
        IsActive:  req.IsActive,
    }

    if err := h.Service.CreateStore(c.Context(), store); err != nil {
        h.Log.WithError(err).Error("Failed to create store")
        return c.Status(http.StatusInternalServerError).JSON(dto.WebErrorResponse{
            Error: "Failed to create store",
        })
    }

    resp := dto.StoreResponse{
        ID:        store.ID,
        Name:      store.Name,
        Location:  store.Location,
        Address:   store.Address,
        Phone:     store.Phone,
        Email:     store.Email,
        StoreSlug: store.StoreSlug,
        IsActive:  store.IsActive,
        CreatedAt: store.CreatedAt,
        UpdatedAt: store.UpdatedAt,
    }

    return c.Status(http.StatusCreated).JSON(dto.WebSuccessResponse[dto.StoreResponse]{
        Data: resp,
    })
}

// GetAllStores handles GET /api/stores
func (h *StoreHandler) GetAllStores(c *fiber.Ctx) error {
    stores, err := h.Service.GetAllStores(c.Context())
    if err != nil {
        h.Log.WithError(err).Error("Failed to fetch stores")
        return c.Status(http.StatusInternalServerError).JSON(dto.WebErrorResponse{
            Error: "Failed to fetch stores",
        })
    }

    var responses []dto.StoreResponse
    for _, store := range stores {
        responses = append(responses, dto.StoreResponse{
            ID:        store.ID,
            Name:      store.Name,
            Location:  store.Location,
            Address:   store.Address,
            Phone:     store.Phone,
            Email:     store.Email,
            StoreSlug: store.StoreSlug,
            IsActive:  store.IsActive,
            CreatedAt: store.CreatedAt,
            UpdatedAt: store.UpdatedAt,
        })
    }

    return c.JSON(dto.WebSuccessResponse[[]dto.StoreResponse]{
        Data: responses,
    })
}

// GetStoreByID handles GET /api/stores/:id
func (h *StoreHandler) GetStoreByID(c *fiber.Ctx) error {
    id, err := uuid.Parse(c.Params("id"))
    if err != nil {
        return c.Status(http.StatusBadRequest).JSON(dto.WebErrorResponse{
            Error: "Invalid store ID",
        })
    }

    store, err := h.Service.GetStoreByID(c.Context(), id)
    if err != nil {
        h.Log.WithError(err).Error("Failed to fetch store")
        return c.Status(http.StatusNotFound).JSON(dto.WebErrorResponse{
            Error: "Store not found",
        })
    }

    resp := dto.StoreResponse{
        ID:        store.ID,
        Name:      store.Name,
        Location:  store.Location,
        Address:   store.Address,
        Phone:     store.Phone,
        Email:     store.Email,
        StoreSlug: store.StoreSlug,
        IsActive:  store.IsActive,
        CreatedAt: store.CreatedAt,
        UpdatedAt: store.UpdatedAt,
    }

    return c.JSON(dto.WebSuccessResponse[dto.StoreResponse]{
        Data: resp,
    })
}

// GetStoreBySlug handles GET /api/stores/slug/:slug
func (h *StoreHandler) GetStoreBySlug(c *fiber.Ctx) error {
    slug := c.Params("slug")
    store, err := h.Service.GetStoreBySlug(c.Context(), slug)
    if err != nil {
        h.Log.WithError(err).Error("Failed to fetch store by slug")
        return c.Status(http.StatusNotFound).JSON(dto.WebErrorResponse{
            Error: "Store not found",
        })
    }

    resp := dto.StoreResponse{
        ID:        store.ID,
        Name:      store.Name,
        Location:  store.Location,
        Address:   store.Address,
        Phone:     store.Phone,
        Email:     store.Email,
        StoreSlug: store.StoreSlug,
        IsActive:  store.IsActive,
        CreatedAt: store.CreatedAt,
        UpdatedAt: store.UpdatedAt,
    }

    return c.JSON(dto.WebSuccessResponse[dto.StoreResponse]{
        Data: resp,
    })
}

// UpdateStore handles PUT /api/stores/:id
func (h *StoreHandler) UpdateStore(c *fiber.Ctx) error {
    id, err := uuid.Parse(c.Params("id"))
    if err != nil {
        return c.Status(http.StatusBadRequest).JSON(dto.WebErrorResponse{
            Error: "Invalid store ID",
        })
    }

    var req dto.StoreRequest
    if err := c.BodyParser(&req); err != nil {
        return c.Status(http.StatusBadRequest).JSON(dto.WebErrorResponse{
            Error: "Invalid request body",
        })
    }

    store := &entity.Store{
        ID:        id,
        Name:      req.Name,
        Location:  req.Location,
        Address:   req.Address,
        Phone:     req.Phone,
        Email:     req.Email,
        StoreSlug: req.StoreSlug,
        IsActive:  req.IsActive,
    }

    if err := h.Service.UpdateStore(c.Context(), store); err != nil {
        h.Log.WithError(err).Error("Failed to update store")
        return c.Status(http.StatusInternalServerError).JSON(dto.WebErrorResponse{
            Error: "Failed to update store",
        })
    }

    return c.JSON(dto.WebSuccessResponse[string]{
        Data: "Store updated successfully",
    })
}

// DeleteStore handles DELETE /api/stores/:id
func (h *StoreHandler) DeleteStore(c *fiber.Ctx) error {
    id, err := uuid.Parse(c.Params("id"))
    if err != nil {
        return c.Status(http.StatusBadRequest).JSON(dto.WebErrorResponse{
            Error: "Invalid store ID",
        })
    }

    if err := h.Service.DeleteStore(c.Context(), id); err != nil {
        h.Log.WithError(err).Error("Failed to delete store")
        return c.Status(http.StatusInternalServerError).JSON(dto.WebErrorResponse{
            Error: "Failed to delete store",
        })
    }

    return c.JSON(dto.WebSuccessResponse[string]{
        Data: "Store deleted successfully",
    })
}

// SetStoreActive handles PATCH /api/stores/:id/active
func (h *StoreHandler) SetStoreActive(c *fiber.Ctx) error {
    id, err := uuid.Parse(c.Params("id"))
    if err != nil {
        return c.Status(http.StatusBadRequest).JSON(dto.WebErrorResponse{
            Error: "Invalid store ID",
        })
    }

    var req struct {
        IsActive bool `json:"is_active"`
    }
    if err := c.BodyParser(&req); err != nil {
        return c.Status(http.StatusBadRequest).JSON(dto.WebErrorResponse{
            Error: "Invalid request body",
        })
    }

    if err := h.Service.SetStoreActive(c.Context(), id, req.IsActive); err != nil {
        h.Log.WithError(err).Error("Failed to update store active status")
        return c.Status(http.StatusInternalServerError).JSON(dto.WebErrorResponse{
            Error: "Failed to update store active status",
        })
    }

    return c.JSON(dto.WebSuccessResponse[string]{
        Data: "Store active status updated successfully",
    })
}