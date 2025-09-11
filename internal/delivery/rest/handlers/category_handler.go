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

type CategoryHandler struct {
    Service model.CategoryService
    Log     *logrus.Entry
}

func NewCategoryHandler(service model.CategoryService, logger *logrus.Logger) *CategoryHandler {
    return &CategoryHandler{
        Service: service,
        Log:     logger.WithField("layer", "handler").WithField("struct", "CategoryHandler"),
    }
}

// CreateCategory handles POST /api/categories
func (h *CategoryHandler) CreateCategory(c *fiber.Ctx) error {
    h.Log.Info("Received create category request")

    var req dto.CategoryRequest
    if err := c.BodyParser(&req); err != nil {
        return c.Status(http.StatusBadRequest).JSON(dto.WebErrorResponse{
            Error: "Invalid request body",
        })
    }

    category := &entity.Category{
        Name:     req.Name,
        Icon:     req.Icon,
        Color:    req.Color,
        IsActive: true,
    }

    if err := h.Service.CreateCategory(c.Context(), category); err != nil {
        return c.Status(http.StatusInternalServerError).JSON(dto.WebErrorResponse{
            Error: "Failed to create category",
        })
    }

    resp := dto.CategoryResponse{
        ID:       category.ID,
        Name:     category.Name,
        Icon:     category.Icon,
        Color:    category.Color,
        IsActive: category.IsActive,
    }

    return c.Status(http.StatusCreated).JSON(dto.WebSuccessResponse[dto.CategoryResponse]{
        Data: resp,
    })
}

// GetAllCategories handles GET /api/categories
func (h *CategoryHandler) GetAllCategories(c *fiber.Ctx) error {
    h.Log.Info("Fetching all categories")

    categories, err := h.Service.GetAllCategories(c.Context())
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(dto.WebErrorResponse{
            Error: "Failed to fetch categories",
        })
    }

    var responses []dto.CategoryResponse
    for _, cat := range categories {
        responses = append(responses, dto.CategoryResponse{
            ID:       cat.ID,
            Name:     cat.Name,
            Icon:     cat.Icon,
            Color:    cat.Color,
            IsActive: cat.IsActive,
        })
    }

    return c.JSON(dto.WebSuccessResponse[[]dto.CategoryResponse]{
        Data: responses,
    })
}

// GetCategoryByID handles GET /api/categories/:id
func (h *CategoryHandler) GetCategoryByID(c *fiber.Ctx) error {
    id, err := uuid.Parse(c.Params("id"))
    if err != nil {
        return c.Status(http.StatusBadRequest).JSON(dto.WebErrorResponse{
            Error: "Invalid category ID",
        })
    }

    category, err := h.Service.GetCategoryByID(c.Context(), id)
    if err != nil {
        return c.Status(http.StatusNotFound).JSON(dto.WebErrorResponse{
            Error: "Category not found",
        })
    }

    resp := dto.CategoryResponse{
        ID:       category.ID,
        Name:     category.Name,
        Icon:     category.Icon,
        Color:    category.Color,
        IsActive: category.IsActive,
    }

    return c.JSON(dto.WebSuccessResponse[dto.CategoryResponse]{
        Data: resp,
    })
}

// UpdateCategory handles PUT /api/categories/:id
func (h *CategoryHandler) UpdateCategory(c *fiber.Ctx) error {
    id, err := uuid.Parse(c.Params("id"))
    if err != nil {
        return c.Status(http.StatusBadRequest).JSON(dto.WebErrorResponse{
            Error: "Invalid category ID",
        })
    }

    var req dto.CategoryRequest
    if err := c.BodyParser(&req); err != nil {
        return c.Status(http.StatusBadRequest).JSON(dto.WebErrorResponse{
            Error: "Invalid request body",
        })
    }

    category := &entity.Category{
        ID:       id,
        Name:     req.Name,
        Icon:     req.Icon,
        Color:    req.Color,
        IsActive: req.IsActive,
    }

    if err := h.Service.UpdateCategory(c.Context(), category); err != nil {
        return c.Status(http.StatusInternalServerError).JSON(dto.WebErrorResponse{
            Error: "Failed to update category",
        })
    }

    return c.Status(http.StatusOK).JSON(dto.WebSuccessResponse[string]{
        Data: "Category updated successfully",
    })
}

// DeleteCategory handles DELETE /api/categories/:id
func (h *CategoryHandler) DeleteCategory(c *fiber.Ctx) error {
    id, err := uuid.Parse(c.Params("id"))
    if err != nil {
        return c.Status(http.StatusBadRequest).JSON(dto.WebErrorResponse{
            Error: "Invalid category ID",
        })
    }

    if err := h.Service.DeleteCategory(c.Context(), id); err != nil {
        return c.Status(http.StatusInternalServerError).JSON(dto.WebErrorResponse{
            Error: "Failed to delete category",
        })
    }

    return c.Status(http.StatusOK).JSON(dto.WebSuccessResponse[string]{
        Data: "Category deleted successfully",
    })
}