package repo

import (
	"context"
	"strconv"

	"github.com/av-huette/avh-booking-system/internal/database"
	"github.com/av-huette/avh-booking-system/internal/models"
	"github.com/jackc/pgx/v5"
)

// UnitStore provides database operations for Unit entities.
type UnitStore struct {
	DB DBTx
}

// Get retrieves units based on the provided query specification.
func (m *UnitStore) Get(ctx context.Context, query *Query) ([]models.Unit, error) {
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
func (m *UnitStore) GetByID(ctx context.Context, id int) (*models.Unit, error) {
	query := Query{Table: TableUnit, Filter: []Filter{{Column: "unit_id", Operator: Eq, Value: strconv.Itoa(id)}}}
	units, err := m.Get(ctx, &query)
	if err != nil {
		return nil, err
	}
	if len(units) == 0 {
		return nil, database.ErrNoRecord
	}
	return &units[0], nil
}

// Insert adds a new unit to the database.
func (m *UnitStore) Insert(ctx context.Context, unit models.Unit) (int, error) {
	query := `
        INSERT INTO unit (name)
        VALUES ($1)
        RETURNING unit_id;`
	var id int
	err := m.DB.QueryRow(ctx, query, unit.Name).Scan(&id)

	return id, err
}

// Update modifies an existing unit in the database.
func (m *UnitStore) Update(ctx context.Context, unit models.Unit) (int, error) {
	query := `
        UPDATE unit
        SET name = $1
        WHERE unit_id = $2
        RETURNING unit_id`
	var id int
	err := m.DB.QueryRow(ctx, query, unit.Name, unit.ID).Scan(&id)

	return id, err
}

// Delete removes a unit from the database.
func (m *UnitStore) Delete(ctx context.Context, id int) (int, error) {
	query := `
        DELETE FROM unit
        WHERE unit_id = $1
        RETURNING unit_id`
	err := m.DB.QueryRow(ctx, query, id).Scan(&id)

	return id, err
}
