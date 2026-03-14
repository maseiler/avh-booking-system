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

// ProductModel provides database operations for Product entities.
type ProductModel struct {
	DB *database.DB
}

// Get retrieves products based on the provided query specification.
func (m *ProductModel) Get(query *Query) ([]models.Product, error) {
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

	products, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Product])
	if err != nil {
		fmt.Printf("CollectRows error: %v", err)
		return nil, err
	}

	return products, nil
}

// GetById retrieves a product by its ID.
func (m *ProductModel) GetById(productId int) (*models.Product, error) {
	ctx := context.Background()
	stmt := `SELECT product_id, name, price, vat, product_group, size, unit, category, created_at
			FROM product
			WHERE product_id = $1`
	row := m.DB.QueryRow(ctx, stmt, productId)

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
func (m *ProductModel) Insert(product models.Product) (int, error) {
	ctx := context.Background()
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
