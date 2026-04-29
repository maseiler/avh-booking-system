package repo

import (
	"context"
	"strconv"

	"github.com/av-huette/avh-booking-system/internal/database"
	"github.com/av-huette/avh-booking-system/internal/models"
	"github.com/jackc/pgx/v5"
)

// ProductStore provides database operations for Product entities.
type ProductStore struct {
	DB DBTx
}

// Get retrieves products based on the provided query specification.
func (m *ProductStore) Get(ctx context.Context, query *Query) ([]models.Product, error) {
	stmt, args, err := buildSelectSQL(query)
	if err != nil {
		return nil, err
	}
	rows, err := m.DB.Query(ctx, stmt, args...)
	if err != nil {
		return nil, err
	}

	products, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Product])
	if err != nil {
		return nil, err
	}

	return products, nil
}

// GetByID retrieves a product by its ID.
func (m *ProductStore) GetByID(ctx context.Context, productID int) (*models.Product, error) {
	query := Query{Table: TableProduct, Filter: []Filter{{Column: "product_id", Operator: Eq, Value: strconv.Itoa(productID)}}}
	products, err := m.Get(ctx, &query)
	if err != nil {
		return nil, err
	}
	if len(products) == 0 {
		return nil, database.ErrNoRecord
	}
	return &products[0], nil
}

// Insert adds a new product to the database.
func (m *ProductStore) Insert(ctx context.Context, product models.Product) (int, error) {
	query := `
        INSERT INTO product (name, price, vat, product_group, size, unit, category)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
        RETURNING product_id`
	var id int
	err := m.DB.QueryRow(ctx, query,
		product.Name,
		product.Price,
		product.VatID,
		product.ProductGroupID,
		product.Size,
		product.UnitID,
		product.CategoryID,
	).Scan(&id)

	return id, err
}
