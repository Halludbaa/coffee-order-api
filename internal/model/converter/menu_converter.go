// converter/menu_converter.go
package converter

import (
	"coffee/internal/entity"
	"coffee/internal/model/dto"

	"github.com/google/uuid"
)

func ToMenuItemResponse(item *entity.MenuItem) dto.MenuItemResponse {
    return dto.MenuItemResponse{
        ID:          item.ID,
        Name:        item.Name,
        Description: item.Description,
        BasePrice:   item.BasePrice,
        CategoryID:  item.CategoryID,
        IsActive:    item.IsActive,
        ImageURL:    item.ImageURL,
    }
}

func ToMenuItemResponses(items []*entity.MenuItem) []dto.MenuItemResponse {
    var responses []dto.MenuItemResponse
    for _, item := range items {
        responses = append(responses, ToMenuItemResponse(item))
    }
    return responses
}

func ToMenuItemEntity(dto *dto.MenuItemRequest) *entity.MenuItem {
    return &entity.MenuItem{
        ID:          uuid.New(), // Generate new ID
        Name:        dto.Name,
        Description: dto.Description,
        BasePrice:   dto.BasePrice,
        CategoryID:  dto.CategoryID,
        IsActive:    dto.IsActive,
        ImageURL:    dto.ImageURL,
    }
}

func UpdateMenuItemEntity(dto *dto.MenuItemRequest, entity *entity.MenuItem) {
    entity.Name = dto.Name
    entity.Description = dto.Description
    entity.BasePrice = dto.BasePrice
    entity.CategoryID = dto.CategoryID
    entity.IsActive = dto.IsActive
    entity.ImageURL = dto.ImageURL
}