package dto

import "github.com/google/uuid"

type CategoryRequest struct {
    Name     string `json:"name"`
    Icon     string `json:"icon"`
    Color    string `json:"color"`
    IsActive bool   `json:"is_active"`
}

type CategoryResponse struct {
    ID       uuid.UUID `json:"id"`
    Name     string    `json:"name"`
    Icon     string    `json:"icon"`
    Color    string    `json:"color"`
    IsActive bool      `json:"is_active"`
}

type StoreCategoryRequest struct {
    CategoryID uuid.UUID `json:"category_id"`
    Name       string    `json:"name"`
    IsVisible  bool      `json:"is_visible"`
    SortOrder  int       `json:"sort_order"`
}