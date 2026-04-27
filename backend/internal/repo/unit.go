package repo

import (
	"context"
	"database/sql"
	"errors"

	"github.com/av-huette/avh-booking-system/internal/database"
	"github.com/av-huette/avh-booking-system/internal/models"
	"github.com/jackc/pgx/v5"
)

// UnitModel provides database operations for Unit entities.
type UnitModel struct {
	DB *database.DB
}

// Get retrieves units based on the provided query specification.
func (m *UnitModel) Get(ctx context.Context, query *Query) ([]models.Unit, error) {
	stmt, args, err := buildSelectSQL(query)
	if err != nil {
		return nil, err
	}
	rows, err := m.DB.Query(ctx, stmt, args...)
	if err != nil {
		return nil, err
	}

	units, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Unit])
	if err != nil {
		return nil, err
	}

	return units, nil
}

// GetByID retrieves a unit by its ID.
func (m *UnitModel) GetByID(ctx context.Context, id int) (*models.Unit, error) {
	stmt := `SELECT unit_id, name
			FROM unit
			WHERE unit_id = $1`
	row := m.DB.QueryRow(ctx, stmt, id)

	var unit models.Unit
	err := row.Scan(&unit.ID, &unit.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, database.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return &unit, nil
}

// Insert adds a new unit to the database.
func (m *UnitModel) Insert(ctx context.Context, unit models.Unit) (int, error) {
	query := `
        INSERT INTO unit (name)
        VALUES ($1)
        RETURNING unit_id;`
	var id int
	err := m.DB.QueryRow(ctx, query, unit.Name).Scan(&id)

	return id, err
}
