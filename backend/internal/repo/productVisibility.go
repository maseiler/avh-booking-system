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

// ProductVisibilityModel provides database operations for ProductVisibility entities.
type ProductVisibilityModel struct {
	DB *database.DB
}

// Get retrieves ProductVisibilities based on the provided query specification.
func (m *ProductVisibilityModel) Get(query *Query) ([]models.ProductVisibility, error) {
	ctx := context.Background()
	stmt := buildSelectSQL(query)
	rows, err := m.DB.Query(ctx, stmt)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			fmt.Println(pgErr.Message)
			fmt.Println(pgErr.Code)
		}
	}

	visibilities, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.ProductVisibility])
	if err != nil {
		fmt.Printf("CollectRows error: %v", err)
		return nil, err
	}

	return visibilities, nil
}

// GetById retrieves a product visibility rule by its ID.
func (m *ProductVisibilityModel) GetById(id int) (*models.ProductVisibility, error) {
	ctx := context.Background()
	stmt := `SELECT product_visibility_id, category, location, product
			FROM product_visibility
			WHERE product_visibility_id = $1`
	row := m.DB.QueryRow(ctx, stmt, id)

	var productVisibility models.ProductVisibility
	err := row.Scan(&productVisibility.Id, &productVisibility.CategoryId, &productVisibility.LocationId, &productVisibility.ProductId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, database.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return &productVisibility, nil
}

// Insert adds a new product visibility rule to the database.
func (m *ProductVisibilityModel) Insert(visibility models.ProductVisibility) (int, error) {
	ctx := context.Background()
	var id int
	query := `
        INSERT INTO product_visibility (category, location, product)
        VALUES ($1, $2, $3)
        RETURNING product_visibility_id`
	err := m.DB.QueryRow(ctx, query, visibility.CategoryId, visibility.LocationId, visibility.ProductId).Scan(&id)

	return id, err
}

// Delete removes an existing visibility from the database.
func (m *ProductVisibilityModel) Delete(id int) (int, error) {
	ctx := context.Background()
	query := `
        DELETE FROM product_visibility
        WHERE product_visibility_id = $1
			RETURNING $1;
			`
	err := m.DB.QueryRow(ctx, query, id).Scan(&id)

	return id, err
}
