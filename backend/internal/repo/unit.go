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

// UnitModel provides database operations for Unit entities.
type UnitModel struct {
	DB *database.DB
}

// Get retrieves units based on the provided query specification.
func (m *UnitModel) Get(query *Query) ([]models.Unit, error) {
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

	units, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Unit])
	if err != nil {
		fmt.Printf("CollectRows error: %v", err)
		return nil, err
	}

	return units, nil
}

// GetById retrieves a unit by its ID.
func (m *UnitModel) GetById(id int) (*models.Unit, error) {
	ctx := context.Background()
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
func (m *UnitModel) Insert(unit models.Unit) (int, error) {
	ctx := context.Background()
	query := `
        INSERT INTO unit (name)
        VALUES ($1)
        RETURNING unit_id;`
	var id int
	err := m.DB.QueryRow(ctx, query, unit.Name).Scan(&id)

	return id, err
}
