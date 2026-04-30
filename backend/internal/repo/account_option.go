package repo

import (
	"context"
	"errors"

	"github.com/av-huette/avh-booking-system/internal/database"
	"github.com/av-huette/avh-booking-system/internal/models"
	"github.com/jackc/pgx/v5"
)

// AccountOptionStore provides database operations for AccountOption entities.
type AccountOptionStore struct {
	DB DBTx
}

// Insert adds a new account option to the database.
func (m *AccountOptionStore) Insert(ctx context.Context, opt models.AccountOption) (int, string, error) {
	query := `
        INSERT INTO account_option (account, key, value)
        VALUES ($1, $2, $3)
        RETURNING account, key`
	var account int
	var key string
	err := m.DB.QueryRow(ctx, query,
		opt.AccountID,
		opt.Key,
		opt.Value,
	).Scan(&account, &key)

	return account, key, err
}

// Update modifies the value of an existing account option.
func (m *AccountOptionStore) Update(ctx context.Context, opt models.AccountOption) error {
	query := `
        UPDATE account_option
        SET value = $1
        WHERE account = $2 AND key = $3`
	_, err := m.DB.Exec(ctx, query, opt.Value, opt.AccountID, opt.Key)

	return err
}

// Get retrieves a specific account option by account ID and key.
func (m *AccountOptionStore) Get(ctx context.Context, accountID int, key string) (*models.AccountOption, error) {
	stmt := `SELECT account, key, value
			FROM account_option
			WHERE account = $1 AND key = $2`
	row := m.DB.QueryRow(ctx, stmt, accountID, key)

	var opt models.AccountOption
	err := row.Scan(&opt.AccountID, &opt.Key, &opt.Value)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, database.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return &opt, nil
}
