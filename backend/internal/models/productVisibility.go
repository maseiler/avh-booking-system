package models

import (
	"context"
	"database/sql"
	"errors"
	"github.com/av-huette/avh-booking-system/internal/database"
)

type ProductVisibility struct {
	Id         int `json:"id"`
	CategoryId int `json:"category_id"`
	ProductId  int `json:"product_id"`
}

type ProductVisibilityModel struct {
	DB *database.DB
}

func CreateProductVisibility(category int, product int) ProductVisibility {
	return ProductVisibility{
		CategoryId: category,
		ProductId:  product,
	}
}

func (m *ProductVisibilityModel) Insert(visibility ProductVisibility) (int, error) {
	ctx := context.Background()
	var id int
	query := `
        INSERT INTO product_visibility (category, product) 
        VALUES ($1, $2)
        RETURNING product_visibility_id`
	err := m.DB.QueryRow(ctx, query, visibility.CategoryId, visibility.ProductId).Scan(&id)

	return id, err
}

func (m *ProductVisibilityModel) Get(id int) (*ProductVisibility, error) {
	ctx := context.Background()
	stmt := `SELECT product_visibility_id, category, product
			FROM product_visibility
			WHERE product_visibility_id = $1`
	row := m.DB.QueryRow(ctx, stmt, id)

	var productVisibility ProductVisibility
	err := row.Scan(&productVisibility.Id, &productVisibility.CategoryId, &productVisibility.ProductId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, database.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return &productVisibility, nil
}
