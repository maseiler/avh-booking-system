package models

import (
	"context"
	"database/sql"
	"errors"
	"github.com/av-huette/avh-booking-system/internal/database"
)

type AccountOption struct {
	AccountId int    `json:"id"`
	Key       string `json:"key"`
	Value     string `json:"value"`
}

type AccountOptionModel struct {
	DB *database.DB
}

func CreateAccountOption(accountId int, key string, value string) AccountOption {
	return AccountOption{
		AccountId: accountId,
		Key:       key,
		Value:     value,
	}
}

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
