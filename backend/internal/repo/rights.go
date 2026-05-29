package repo

import (
	"context"
	"strconv"

	"github.com/av-huette/avh-booking-system/internal/database"
	"github.com/av-huette/avh-booking-system/internal/models"
	"github.com/jackc/pgx/v5"
)

// RightsStore provides database operations for Rights entities.
type RightsStore struct {
	DB DBTx
}

// Get retrieves rights based on the provided query specification.
func (m *RightsStore) Get(ctx context.Context, query *Query) ([]models.Rights, error) {
	stmt, args, err := buildSelectSQL(query)
	if err != nil {
		return nil, err
	}
	rows, err := m.DB.Query(ctx, stmt, args...)
	if err != nil {
		return nil, err
	}

	rights, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Rights])
	if err != nil {
		return nil, err
	}

	return rights, nil
}

// GetByID retrieves a rights entry by its ID.
func (m *RightsStore) GetByID(ctx context.Context, id int) (*models.Rights, error) {
	query := Query{Table: TableRights, Filter: []Filter{{Column: "rights_id", Operator: Eq, Value: strconv.Itoa(id)}}}
	rights, err := m.Get(ctx, &query)
	if err != nil {
		return nil, err
	}
	if len(rights) == 0 {
		return nil, database.ErrNoRecord
	}
	return &rights[0], nil
}

// Insert adds a new rights entry to the database.
func (m *RightsStore) Insert(ctx context.Context, r models.Rights) (int, error) {
	query := `
        INSERT INTO rights (role, permission, allowed)
        VALUES ($1, $2, $3)
        RETURNING rights_id`
	var id int
	err := m.DB.QueryRow(ctx, query, r.RoleID, r.Permission, r.Allowed).Scan(&id)
	return id, err
}

// Update modifies an existing rights entry.
func (m *RightsStore) Update(ctx context.Context, r models.Rights) (int, error) {
	query := `
        UPDATE rights
        SET role = $1, permission = $2, allowed = $3
        WHERE rights_id = $4
        RETURNING rights_id`
	var id int
	err := m.DB.QueryRow(ctx, query, r.RoleID, r.Permission, r.Allowed, r.ID).Scan(&id)
	return id, err
}

// Delete removes a rights entry by ID.
func (m *RightsStore) Delete(ctx context.Context, id int) (int, error) {
	query := `DELETE FROM rights WHERE rights_id = $1 RETURNING rights_id`
	err := m.DB.QueryRow(ctx, query, id).Scan(&id)
	return id, err
}
