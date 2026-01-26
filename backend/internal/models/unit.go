package models

import (
	"context"
	"database/sql"
	"errors"
	"github.com/av-huette/avh-booking-system/internal/database"
)

type Unit struct {
	Id   int    `json:"id" db:"unit_id"`
	Name string `json:"name" db:"name"`
}

type UnitModel struct {
	DB *database.DB
}

func CreateUnit(name string) Unit {
	return Unit{
		Name: name,
	}
}

func (m *UnitModel) Insert(unit Unit) (int, error) {
	ctx := context.Background()
	query := `
        INSERT INTO unit (name) 
        VALUES ($1)
        RETURNING unit_id;`
	var id int
	err := m.DB.QueryRow(ctx, query, unit.Name).Scan(&id)

	return id, err
}

func (m *UnitModel) Get(id int) (*Unit, error) {
	ctx := context.Background()
	stmt := `SELECT unit_id, name
			FROM unit
			WHERE unit_id = $1`
	row := m.DB.QueryRow(ctx, stmt, id)

	var unit Unit
	err := row.Scan(&unit.Id, &unit.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, database.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return &unit, nil
}
