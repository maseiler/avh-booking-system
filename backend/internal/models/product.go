package models

import (
	"context"
	"database/sql"
	"errors"
	"github.com/av-huette/avh-booking-system/internal/database"
	"github.com/jackc/pgx/v5/pgtype"
)

type Product struct {
	Id             int              `json:"id"`
	Name           string           `json:"name"`
	Price          pgtype.Numeric   `json:"price"`
	VatId          int              `json:"vatId"`
	ProductGroupId int              `json:"productGroupId"`
	Size           int              `json:"size"`
	UnitId         int              `json:"unitId"`
	CategoryId     int              `json:"categoryId"`
	CreatedAt      pgtype.Timestamp `json:"createdAt"`
}

type ProductModel struct {
	DB *database.DB
}

func CreateProduct(name string, price string, vatId int, prodGroupId int, size int,
	unitId int, catId int) Product {
	return Product{
		Name:           name,
		Price:          NewNumeric(price),
		VatId:          vatId,
		ProductGroupId: prodGroupId,
		Size:           size,
		UnitId:         unitId,
		CategoryId:     catId,
	}
}

func (m *ProductModel) Insert(product Product) (int, error) {
	ctx := context.Background()
	query := `
        INSERT INTO product (name, price, vat, product_group, size, unit, category) 
        VALUES ($1, $2, $3, $4, $5, $6, $7)
        RETURNING product_id`
	var id int
	err := m.DB.QueryRow(ctx, query,
		product.Name,
		product.Price,
		product.VatId,
		product.ProductGroupId,
		product.Size,
		product.UnitId,
		product.CategoryId,
	).Scan(&id)

	return id, err
}

func (m *ProductModel) Get(productId int) (*Product, error) {
	ctx := context.Background()
	stmt := `SELECT product_id, name, price, vat, product_group, size, unit, category, created_at
			FROM product
			WHERE product_id = $1`
	row := m.DB.QueryRow(ctx, stmt, productId)

	var product Product
	err := row.Scan(&product.Id, &product.Name, &product.Price, &product.VatId, &product.ProductGroupId,
		&product.Size, &product.UnitId, &product.CategoryId, &product.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, database.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return &product, nil
}
