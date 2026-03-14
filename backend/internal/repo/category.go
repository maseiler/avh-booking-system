package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/av-huette/avh-booking-system/internal/database"
	"github.com/av-huette/avh-booking-system/internal/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// CategoryModel provides database operations for Category entities.
type CategoryModel struct {
	DB *database.DB
}

// Get retrieves categories based on the provided query specification.
func (m *CategoryModel) Get(query *Query) ([]models.Category, error) {
	ctx := context.Background()
	stmt, args, err := buildSelectSQL(query)
	if err != nil {
		return nil, err
	}
	rows, err := m.DB.Query(ctx, stmt, args...)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			fmt.Println(pgErr.Message)
			fmt.Println(pgErr.Code)
		}
	}

	categories, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Category])
	if err != nil {
		fmt.Printf("CollectRows error: %v", err)
		return nil, err
	}

	return categories, nil
}

// GetById retrieves a category by its ID.
func (m *CategoryModel) GetById(id int) (*models.Category, error) {
	ctx := context.Background()
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
func (m *CategoryModel) Insert(category models.Category) (int, error) {
	ctx := context.Background()
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
