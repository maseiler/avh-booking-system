package repo

import (
	"context"
	"database/sql"
	"errors"

	"github.com/av-huette/avh-booking-system/internal/database"
	"github.com/av-huette/avh-booking-system/internal/models"
	"github.com/jackc/pgx/v5"
)

// ProductModel provides database operations for Product entities.
type ProductModel struct {
	DB *database.DB
}

// Get retrieves products based on the provided query specification.
func (m *ProductModel) Get(ctx context.Context, query *Query) ([]models.Product, error) {
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
func (m *ProductModel) GetByID(ctx context.Context, productID int) (*models.Product, error) {
	stmt := `SELECT product_id, name, price, vat, product_group, size, unit, category, created_at
			FROM product
			WHERE product_id = $1`
	row := m.DB.QueryRow(ctx, stmt, productID)

	var product models.Product
	err := row.Scan(&product.ID, &product.Name, &product.Price, &product.VatID, &product.ProductGroupID,
		&product.Size, &product.UnitID, &product.CategoryID, &product.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, database.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return &product, nil
}

// Insert adds a new product to the database.
func (m *ProductModel) Insert(ctx context.Context, product models.Product) (int, error) {
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
