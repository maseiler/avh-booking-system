package models

import (
	"context"
	"database/sql"
	"errors"
	"github.com/av-huette/avh-booking-system/internal/database"
)

type Location struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type LocationModel struct {
	DB *database.DB
}

func CreateLocation(name string) Location {
	return Location{
		Name: name,
	}
}

func (m *LocationModel) Insert(location Location) (int, error) {
	ctx := context.Background()
	query := `
        INSERT INTO location (name) 
        VALUES ($1)
        RETURNING location_id;`
	var id int
	err := m.DB.QueryRow(ctx, query, location.Name).Scan(&id)

	return id, err
}

func (m *LocationModel) Get(id int) (*Location, error) {
	ctx := context.Background()
	stmt := `SELECT location_id, name
			FROM location
			WHERE location_id = $1`
	row := m.DB.QueryRow(ctx, stmt, id)

	var location Location
	err := row.Scan(&location.Id, &location.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, database.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return &location, nil
}
