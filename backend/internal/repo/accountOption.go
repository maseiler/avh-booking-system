package repo

import (
	"context"
	"database/sql"
	"errors"

	"github.com/av-huette/avh-booking-system/internal/database"
	"github.com/av-huette/avh-booking-system/internal/models"
)

// AccountOptionModel provides database operations for AccountOption entities.
type AccountOptionModel struct {
	DB *database.DB
}

// Insert adds a new account option to the database.
func (m *AccountOptionModel) Insert(opt models.AccountOption) (int, string, error) {
	ctx := context.Background()
	query := `
        INSERT INTO account_option (account, key, value)
        VALUES ($1, $2, $3)
        RETURNING account, key`
	var account int
	var key string
	err := m.DB.QueryRow(ctx, query,
		opt.AccountId,
		opt.Key,
		opt.Value,
	).Scan(&account, &key)

	return account, key, err
}

// Get retrieves a specific account option by account ID and key.
func (m *AccountOptionModel) Get(accountId int, key string) (*models.AccountOption, error) {
	ctx := context.Background()
	stmt := `SELECT account, key, value
			FROM account_option
			WHERE account = $1 AND key = $2`
	row := m.DB.QueryRow(ctx, stmt, accountId, key)

	var opt models.AccountOption
	err := row.Scan(&opt.AccountId, &opt.Key, &opt.Value)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, database.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return &opt, nil
}
