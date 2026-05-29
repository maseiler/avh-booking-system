package repo

import (
	"context"
	"errors"

	"github.com/av-huette/avh-booking-system/internal/database"
	"github.com/av-huette/avh-booking-system/internal/models"
	"github.com/jackc/pgx/v5"
)

// UserOptionStore provides database operations for UserOption entities.
type UserOptionStore struct {
	DB DBTx
}

// Insert adds a new user option to the database.
func (m *UserOptionStore) Insert(ctx context.Context, opt models.UserOption) (int, string, error) {
	query := `
        INSERT INTO user_option (user_id, key, value)
        VALUES ($1, $2, $3)
        RETURNING user_id, key`
	var userID int
	var key string
	err := m.DB.QueryRow(ctx, query, opt.UserID, opt.Key, opt.Value).Scan(&userID, &key)
	return userID, key, err
}

// Update modifies the value of an existing user option.
func (m *UserOptionStore) Update(ctx context.Context, opt models.UserOption) error {
	query := `
        UPDATE user_option
        SET value = $1
        WHERE user_id = $2 AND key = $3`
	_, err := m.DB.Exec(ctx, query, opt.Value, opt.UserID, opt.Key)
	return err
}

// Get retrieves a specific user option by user ID and key.
func (m *UserOptionStore) Get(ctx context.Context, userID int, key string) (*models.UserOption, error) {
	stmt := `SELECT user_id, key, value FROM user_option WHERE user_id = $1 AND key = $2`
	row := m.DB.QueryRow(ctx, stmt, userID, key)

	var opt models.UserOption
	err := row.Scan(&opt.UserID, &opt.Key, &opt.Value)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, database.ErrNoRecord
		}
		return nil, err
	}

	return &opt, nil
}
