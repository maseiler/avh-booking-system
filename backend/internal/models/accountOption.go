package models

import (
	"context"
	"database/sql"
	"errors"
	"github.com/av-huette/avh-booking-system/internal/database"
)

// AccountOption represents a key-value configuration option for an account.
// It allows storing arbitrary account-specific settings and preferences.
// The combination of AccountId and Key forms a composite primary key.
type AccountOption struct {
	AccountId int    `json:"accountId" db:"account"`
	Key       string `json:"key" db:"key"`
	Value     string `json:"value" db:"value"`
}

// AccountOptionModel provides database operations for AccountOption entities.
type AccountOptionModel struct {
	DB *database.DB
}

// CreateAccountOption creates a new AccountOption with the provided details.
// This is a convenience function for constructing AccountOption structs.
func CreateAccountOption(accountId int, key string, value string) AccountOption {
	return AccountOption{
		AccountId: accountId,
		Key:       key,
		Value:     value,
	}
}

// Insert adds a new account option to the database.
// Returns the account ID and key of the inserted option, along with any error encountered.
// The combination of account ID and key must be unique.
func (m *AccountOptionModel) Insert(opt AccountOption) (int, string, error) {
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
// Returns a pointer to the AccountOption if found, nil otherwise.
// Returns database.ErrNoRecord if no matching option exists.
func (m *AccountOptionModel) Get(accountId int, key string) (*AccountOption, error) {
	ctx := context.Background()
	stmt := `SELECT account, key, value
			FROM account_option
			WHERE account = $1 AND key = $2`
	row := m.DB.QueryRow(ctx, stmt, accountId, key)

	var opt AccountOption
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
