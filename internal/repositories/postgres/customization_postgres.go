package postgres

import (
	"coffee/internal/entity"
	"coffee/internal/model"
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type CustomizationRepository struct {
    DB  *sqlx.DB
    Log *logrus.Entry
}

func NewCustomizationRepository(db *sqlx.DB, log *logrus.Logger) model.CustomizationRepository {
    return &CustomizationRepository{
        DB:  db,
        Log: log.WithField("repository", "customization"),
    }
}

func (r *CustomizationRepository) CreateGroup(ctx context.Context, group *entity.CustomizationGroup) error {
    query := `
        INSERT INTO customization_groups (id, name, store_id, is_required, sort_order)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING created_at
    `

    if group.ID == uuid.Nil {
        group.ID = uuid.New()
    }

    err := r.DB.QueryRowContext(ctx, query,
        group.ID, group.Name, group.StoreID,
        group.IsRequired, group.SortOrder,
    ).Scan(&group.CreatedAt)

    if err != nil {
        return fmt.Errorf("failed to create customization group: %w", err)
    }

    return nil
}

func (r *CustomizationRepository) GetGroupsByStore(ctx context.Context, storeID uuid.UUID) ([]*entity.CustomizationGroup, error) {
    query := `
        SELECT id, name, store_id, is_required, sort_order, created_at
        FROM customization_groups
        WHERE store_id = $1
        ORDER BY sort_order, name
    `

    groups := []*entity.CustomizationGroup{}
    if err := r.DB.SelectContext(ctx, &groups, query, storeID); err != nil {
        return nil, fmt.Errorf("failed to fetch customization groups: %w", err)
    }

    if len(groups) == 0 {
        return groups, nil
    }

    // Collect all group IDs
    groupIDs := make([]uuid.UUID, len(groups))
    groupMap := make(map[uuid.UUID]*entity.CustomizationGroup)
    for i, group := range groups {
        groupIDs[i] = group.ID
        groupMap[group.ID] = group
    }

    // Fetch all options in a single query
    optionsQuery := `
        SELECT id, group_id, label, additional_price, is_available, sort_order, created_at
        FROM customization_options
        WHERE group_id = ANY($1)
        ORDER BY sort_order, label
    `

    options := []*entity.CustomizationOption{}
    if err := r.DB.SelectContext(ctx, &options, optionsQuery, groupIDs); err != nil {
        return nil, fmt.Errorf("failed to fetch customization options: %w", err)
    }

    // Map options to their respective groups
    for _, option := range options {
        if group, ok := groupMap[option.GroupID]; ok {
            group.Options = append(group.Options, option)
        }
    }

    return groups, nil
}

func (r *CustomizationRepository) UpdateGroup(ctx context.Context, group *entity.CustomizationGroup) error {
    query := `
        UPDATE customization_groups 
        SET name = $1, is_required = $2, sort_order = $3
        WHERE id = $4 AND store_id = $5
        RETURNING created_at
    `

    result, err := r.DB.ExecContext(ctx, query,
        group.Name,
        group.IsRequired,
        group.SortOrder,
        group.ID,
        group.StoreID,
    )
    if err != nil {
        return fmt.Errorf("failed to update customization group: %w", err)
    }

    rows, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("error checking rows affected: %w", err)
    }
    if rows == 0 {
        return fmt.Errorf("customization group not found")
    }

    return nil
}

func (r *CustomizationRepository) DeleteGroup(ctx context.Context, id uuid.UUID) error {
    query := `DELETE FROM customization_groups WHERE id = $1`

    result, err := r.DB.ExecContext(ctx, query, id)
    if err != nil {
        return fmt.Errorf("failed to delete customization group: %w", err)
    }

    rows, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("error checking rows affected: %w", err)
    }
    if rows == 0 {
        return fmt.Errorf("customization group not found")
    }

    return nil
}

func (r *CustomizationRepository) CreateOption(ctx context.Context, option *entity.CustomizationOption) error {
    query := `
        INSERT INTO customization_options (id, group_id, label, additional_price, is_available, sort_order)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING created_at
    `

    if option.ID == uuid.Nil {
        option.ID = uuid.New()
    }

    err := r.DB.QueryRowContext(ctx, query,
        option.ID,
        option.GroupID,
        option.Label,
        option.AdditionalPrice,
        option.IsAvailable,
        option.SortOrder,
    ).Scan(&option.CreatedAt)

    if err != nil {
        return fmt.Errorf("failed to create customization option: %w", err)
    }

    return nil
}

func (r *CustomizationRepository) GetOptionsByGroup(ctx context.Context, groupID uuid.UUID) ([]*entity.CustomizationOption, error) {
    query := `
        SELECT id, group_id, label, additional_price, is_available, sort_order, created_at
        FROM customization_options
        WHERE group_id = $1
        ORDER BY sort_order, label
    `

    options := []*entity.CustomizationOption{}
    if err := r.DB.SelectContext(ctx, &options, query, groupID); err != nil {
        return nil, fmt.Errorf("failed to fetch customization options: %w", err)
    }

    return options, nil
}

func (r *CustomizationRepository) UpdateOption(ctx context.Context, option *entity.CustomizationOption) error {
    query := `
        UPDATE customization_options
        SET label = $1, additional_price = $2, is_available = $3, sort_order = $4
        WHERE id = $5 AND group_id = $6
    `

    result, err := r.DB.ExecContext(ctx, query,
        option.Label,
        option.AdditionalPrice,
        option.IsAvailable,
        option.SortOrder,
        option.ID,
        option.GroupID,
    )
    if err != nil {
        return fmt.Errorf("failed to update customization option: %w", err)
    }

    rows, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("error checking rows affected: %w", err)
    }
    if rows == 0 {
        return fmt.Errorf("customization option not found")
    }

    return nil
}

func (r *CustomizationRepository) DeleteOption(ctx context.Context, id uuid.UUID) error {
    query := `DELETE FROM customization_options WHERE id = $1`

    result, err := r.DB.ExecContext(ctx, query, id)
    if err != nil {
        return fmt.Errorf("failed to delete customization option: %w", err)
    }

    rows, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("error checking rows affected: %w", err)
    }
    if rows == 0 {
        return fmt.Errorf("customization option not found")
    }

    return nil
}

func (r *CustomizationRepository) AssignGroupToMenuItem(ctx context.Context, menuItemID, groupID uuid.UUID, isDefault bool) error {
    query := `
        INSERT INTO menu_item_customizations (id, menu_item_id, group_id, is_default)
        VALUES ($1, $2, $3, $4)
    `

    result, err := r.DB.ExecContext(ctx, query,
        uuid.New(),
        menuItemID,
        groupID,
        isDefault,
    )
    if err != nil {
        return fmt.Errorf("failed to assign customization group to menu item: %w", err)
    }

    rows, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("error checking rows affected: %w", err)
    }
    if rows == 0 {
        return fmt.Errorf("failed to assign customization")
    }

    return nil
}

func (r *CustomizationRepository) GetMenuItemCustomizations(ctx context.Context, menuItemID uuid.UUID) ([]*entity.CustomizationGroup, error) {
    query := `
        SELECT g.id, g.name, g.store_id, g.is_required, g.sort_order, g.created_at
        FROM customization_groups g
        JOIN menu_item_customizations mic ON mic.group_id = g.id
        WHERE mic.menu_item_id = $1
        ORDER BY g.sort_order, g.name
    `

    groups := []*entity.CustomizationGroup{}
    if err := r.DB.SelectContext(ctx, &groups, query, menuItemID); err != nil {
        return nil, fmt.Errorf("failed to fetch menu item customizations: %w", err)
    }

    if len(groups) == 0 {
        return groups, nil
    }

    // Collect all group IDs
    groupIDs := make([]uuid.UUID, len(groups))
    groupMap := make(map[uuid.UUID]*entity.CustomizationGroup)
    for i, group := range groups {
        groupIDs[i] = group.ID
        groupMap[group.ID] = group
    }

    // Fetch all options in a single query
    optionsQuery := `
        SELECT id, group_id, label, additional_price, is_available, sort_order, created_at
        FROM customization_options
        WHERE group_id = ANY($1)
        ORDER BY sort_order, label
    `

    options := []*entity.CustomizationOption{}
    if err := r.DB.SelectContext(ctx, &options, optionsQuery, groupIDs); err != nil {
        return nil, fmt.Errorf("failed to fetch customization options: %w", err)
    }

    // Map options to their respective groups
    for _, option := range options {
        if group, ok := groupMap[option.GroupID]; ok {
            group.Options = append(group.Options, option)
        }
    }

    return groups, nil
}

func (r *CustomizationRepository) RemoveGroupFromMenuItem(ctx context.Context, menuItemID, groupID uuid.UUID) error {
    query := `DELETE FROM menu_item_customizations WHERE menu_item_id = $1 AND group_id = $2`

    result, err := r.DB.ExecContext(ctx, query, menuItemID, groupID)
    if err != nil {
        return fmt.Errorf("failed to remove customization group from menu item: %w", err)
    }

    rows, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("error checking rows affected: %w", err)
    }
    if rows == 0 {
        return fmt.Errorf("customization assignment not found")
    }

    return nil
}