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

// LocationModel provides database operations for Location entities.
type LocationModel struct {
	DB *database.DB
}

// Get retrieves locations based on the provided query specification.
func (m *LocationModel) Get(query *Query) ([]models.Location, error) {
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

	locations, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Location])
	if err != nil {
		fmt.Printf("CollectRows error: %v", err)
		return nil, err
	}

	return locations, nil
}

// GetByID retrieves a location by its ID.
func (m *LocationModel) GetByID(id int) (*models.Location, error) {
	ctx := context.Background()
	stmt := `SELECT location_id, name
			FROM location
			WHERE location_id = $1`
	row := m.DB.QueryRow(ctx, stmt, id)

	var location models.Location
	err := row.Scan(&location.ID, &location.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, database.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return &location, nil
}

// Insert adds a new location to the database.
func (m *LocationModel) Insert(location models.Location) (int, error) {
	ctx := context.Background()
	query := `
        INSERT INTO location (name)
        VALUES ($1)
        RETURNING location_id;`
	var id int
	err := m.DB.QueryRow(ctx, query, location.Name).Scan(&id)

	return id, err
}
