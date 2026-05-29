package repo

import (
	"context"
	"strconv"

	"github.com/av-huette/avh-booking-system/internal/database"
	"github.com/av-huette/avh-booking-system/internal/models"
	"github.com/jackc/pgx/v5"
)

// UserStore provides database operations for User entities.
type UserStore struct {
	DB DBTx
}

// Get retrieves users based on the provided query specification.
func (m *UserStore) Get(ctx context.Context, query *Query) ([]models.User, error) {
	stmt, args, err := buildSelectSQL(query)
	if err != nil {
		return nil, err
	}
	rows, err := m.DB.Query(ctx, stmt, args...)
	if err != nil {
		return nil, err
	}

	users, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.User])
	if err != nil {
		return nil, err
	}

	return users, nil
}

// GetByID retrieves a user by their ID.
func (m *UserStore) GetByID(ctx context.Context, id int) (*models.User, error) {
	query := Query{Table: TableUser, Filter: []Filter{{Column: "user_id", Operator: Eq, Value: strconv.Itoa(id)}}}
	users, err := m.Get(ctx, &query)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, database.ErrNoRecord
	}
	return &users[0], nil
}

// Insert adds a new user to the database.
func (m *UserStore) Insert(ctx context.Context, user models.User) (int, error) {
	query := `
        INSERT INTO "user" (name, role, password)
        VALUES ($1, $2, $3)
        RETURNING user_id`
	var id int
	err := m.DB.QueryRow(ctx, query, user.Name, user.RoleID, user.Password).Scan(&id)
	return id, err
}

// Update modifies an existing user.
func (m *UserStore) Update(ctx context.Context, user models.User) (int, error) {
	query := `
        UPDATE "user"
        SET name = $1, role = $2, password = $3
        WHERE user_id = $4
        RETURNING user_id`
	var id int
	err := m.DB.QueryRow(ctx, query, user.Name, user.RoleID, user.Password, user.ID).Scan(&id)
	return id, err
}

// Delete removes a user by ID.
func (m *UserStore) Delete(ctx context.Context, id int) (int, error) {
	query := `DELETE FROM "user" WHERE user_id = $1 RETURNING user_id`
	err := m.DB.QueryRow(ctx, query, id).Scan(&id)
	return id, err
}
