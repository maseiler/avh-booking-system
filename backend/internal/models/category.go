package models

import (
	"context"
	"database/sql"
	"errors"
	"github.com/av-huette/avh-booking-system/internal/database"
)

type Category struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Enabled bool   `json:"enabled"`
	Icon    string `json:"icon"`
	Type    string `json:"type"`
}

type CategoryModel struct {
	DB *database.DB
}

func CreateCategory(name string, icon string, type_ string) Category {
	return Category{
		Name: name,
		Icon: icon,
		Type: type_,
	}
}

func (m *CategoryModel) Insert(category Category) (int, error) {
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

func (m *CategoryModel) Get(id int) (*Category, error) {
	ctx := context.Background()
	stmt := `SELECT category_id, name, enabled, icon, type
			FROM category
			WHERE category_id = $1`
	row := m.DB.QueryRow(ctx, stmt, id)

	var cat Category
	err := row.Scan(&cat.Id, &cat.Name, &cat.Enabled, &cat.Icon, &cat.Type)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, database.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return &cat, nil
}
