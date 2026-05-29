package repo

import (
	"context"
	"strconv"

	"github.com/av-huette/avh-booking-system/internal/database"
	"github.com/av-huette/avh-booking-system/internal/models"
	"github.com/jackc/pgx/v5"
)

// RoleStore provides database operations for Role entities.
type RoleStore struct {
	DB DBTx
}

// Get retrieves roles based on the provided query specification.
func (m *RoleStore) Get(ctx context.Context, query *Query) ([]models.Role, error) {
	stmt, args, err := buildSelectSQL(query)
	if err != nil {
		return nil, err
	}
	rows, err := m.DB.Query(ctx, stmt, args...)
	if err != nil {
		return nil, err
	}

	roles, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Role])
	if err != nil {
		return nil, err
	}

	return roles, nil
}

// GetByID retrieves a role by its ID.
func (m *RoleStore) GetByID(ctx context.Context, id int) (*models.Role, error) {
	query := Query{Table: TableRole, Filter: []Filter{{Column: "role_id", Operator: Eq, Value: strconv.Itoa(id)}}}
	roles, err := m.Get(ctx, &query)
	if err != nil {
		return nil, err
	}
	if len(roles) == 0 {
		return nil, database.ErrNoRecord
	}
	return &roles[0], nil
}

// Insert adds a new role to the database.
func (m *RoleStore) Insert(ctx context.Context, role models.Role) (int, error) {
	query := `INSERT INTO role (name) VALUES ($1) RETURNING role_id`
	var id int
	err := m.DB.QueryRow(ctx, query, role.Name).Scan(&id)
	return id, err
}

// Update modifies an existing role.
func (m *RoleStore) Update(ctx context.Context, role models.Role) (int, error) {
	query := `UPDATE role SET name = $1 WHERE role_id = $2 RETURNING role_id`
	var id int
	err := m.DB.QueryRow(ctx, query, role.Name, role.ID).Scan(&id)
	return id, err
}

// Delete removes a role by ID.
func (m *RoleStore) Delete(ctx context.Context, id int) (int, error) {
	query := `DELETE FROM role WHERE role_id = $1 RETURNING role_id`
	err := m.DB.QueryRow(ctx, query, id).Scan(&id)
	return id, err
}
