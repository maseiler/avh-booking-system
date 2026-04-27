package repo

import (
	"context"
	"database/sql"
	"errors"

	"github.com/av-huette/avh-booking-system/internal/database"
	"github.com/av-huette/avh-booking-system/internal/models"
	"github.com/jackc/pgx/v5"
)

// CategoryModel provides database operations for Category entities.
type CategoryModel struct {
	DB *database.DB
}

// Get retrieves categories based on the provided query specification.
func (m *CategoryModel) Get(ctx context.Context, query *Query) ([]models.Category, error) {
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
func (m *CategoryModel) GetByID(ctx context.Context, id int) (*models.Category, error) {
	stmt := `SELECT category_id, name, enabled, icon, type
			FROM category
			WHERE category_id = $1`
	row := m.DB.QueryRow(ctx, stmt, id)

	var cat models.Category
	err := row.Scan(&cat.ID, &cat.Name, &cat.Enabled, &cat.Icon, &cat.Type)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, database.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return &cat, nil
}

// Insert adds a new category to the database.
func (m *CategoryModel) Insert(ctx context.Context, category models.Category) (int, error) {
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
