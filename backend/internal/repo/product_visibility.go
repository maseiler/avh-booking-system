package repo

import (
	"context"
	"strconv"

	"github.com/av-huette/avh-booking-system/internal/database"
	"github.com/av-huette/avh-booking-system/internal/models"
	"github.com/jackc/pgx/v5"
)

// ProductVisibilityStore provides database operations for ProductVisibility entities.
type ProductVisibilityStore struct {
	DB DBTx
}

// Get retrieves ProductVisibilities based on the provided query specification.
func (m *ProductVisibilityStore) Get(ctx context.Context, query *Query) ([]models.ProductVisibility, error) {
	stmt, args, err := buildSelectSQL(query)
	if err != nil {
		return nil, err
	}
	rows, err := m.DB.Query(ctx, stmt, args...)
	if err != nil {
		return nil, err
	}

	visibilities, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.ProductVisibility])
	if err != nil {
		return nil, err
	}

	return visibilities, nil
}

// GetByID retrieves a product visibility rule by its ID.
func (m *ProductVisibilityStore) GetByID(ctx context.Context, id int) (*models.ProductVisibility, error) {
	query := Query{Table: TableProductVisibility, Filter: []Filter{{Column: "product_visibility_id", Operator: Eq, Value: strconv.Itoa(id)}}}
	visibilities, err := m.Get(ctx, &query)
	if err != nil {
		return nil, err
	}
	if len(visibilities) == 0 {
		return nil, database.ErrNoRecord
	}
	return &visibilities[0], nil
}

// Insert adds a new product visibility rule to the database.
func (m *ProductVisibilityStore) Insert(ctx context.Context, visibility models.ProductVisibility) (int, error) {
	var id int
	query := `
        INSERT INTO product_visibility (category, location, product)
        VALUES ($1, $2, $3)
        RETURNING product_visibility_id`
	err := m.DB.QueryRow(ctx, query, visibility.CategoryID, visibility.LocationID, visibility.ProductID).Scan(&id)

	return id, err
}

// Delete removes an existing visibility from the database.
func (m *ProductVisibilityStore) Delete(ctx context.Context, id int) (int, error) {
	query := `
        DELETE FROM product_visibility
        WHERE product_visibility_id = $1
        RETURNING product_visibility_id`
	err := m.DB.QueryRow(ctx, query, id).Scan(&id)

	return id, err
}
