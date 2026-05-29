package repo

import (
	"context"
	"strconv"

	"github.com/av-huette/avh-booking-system/internal/database"
	"github.com/av-huette/avh-booking-system/internal/models"
	"github.com/jackc/pgx/v5"
)

// CategoryStore provides database operations for Category entities.
type CategoryStore struct {
	DB DBTx
}

// Get retrieves categories based on the provided query specification.
func (m *CategoryStore) Get(ctx context.Context, query *Query) ([]models.Category, error) {
	stmt, args, err := buildSelectSQL(query)
	if err != nil {
		return nil, err
	}
	rows, err := m.DB.Query(ctx, stmt, args...)
	if err != nil {
		return nil, err
	}

	categories, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Category])
	if err != nil {
		return nil, err
	}

	return categories, nil
}

// GetByID retrieves a category by its ID.
func (m *CategoryStore) GetByID(ctx context.Context, id int) (*models.Category, error) {
	query := Query{Table: TableCategory, Filter: []Filter{{Column: "category_id", Operator: Eq, Value: strconv.Itoa(id)}}}
	categories, err := m.Get(ctx, &query)
	if err != nil {
		return nil, err
	}
	if len(categories) == 0 {
		return nil, database.ErrNoRecord
	}
	return &categories[0], nil
}

// Insert adds a new category to the database.
func (m *CategoryStore) Insert(ctx context.Context, category models.Category) (int, error) {
	query := `
        INSERT INTO category (name, enabled, icon, type)
        VALUES ($1, $2, $3, $4)
        RETURNING category_id`
	var id int
	err := m.DB.QueryRow(ctx, query,
		category.Name,
		category.Enabled,
		category.Icon,
		category.Type,
	).Scan(&id)

	return id, err
}

// Update modifies an existing category in the database.
func (m *CategoryStore) Update(ctx context.Context, category models.Category) (int, error) {
	query := `
        UPDATE category
        SET name = $1, enabled = $2, icon = $3, type = $4
        WHERE category_id = $5
        RETURNING category_id`
	var id int
	err := m.DB.QueryRow(ctx, query,
		category.Name,
		category.Enabled,
		category.Icon,
		category.Type,
		category.ID,
	).Scan(&id)

	return id, err
}

// Delete removes a category from the database.
func (m *CategoryStore) Delete(ctx context.Context, id int) (int, error) {
	query := `
        DELETE FROM category
        WHERE category_id = $1
        RETURNING category_id`
	err := m.DB.QueryRow(ctx, query, id).Scan(&id)

	return id, err
}
