package repo

import (
	"context"
	"strconv"

	"github.com/av-huette/avh-booking-system/internal/database"
	"github.com/av-huette/avh-booking-system/internal/models"
	"github.com/jackc/pgx/v5"
)

// AccountModel provides database operations for Account entities.
type AccountModel struct {
	DB *database.DB
}

// Get retrieves accounts based on the provided query specification.
func (m *AccountModel) Get(ctx context.Context, query *Query) ([]models.Account, error) {
	stmt, args, err := buildSelectSQL(query)
	if err != nil {
		return nil, err
	}
	rows, err := m.DB.Query(ctx, stmt, args...)
	if err != nil {
		return nil, err
	}

	accounts, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Account])
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

// GetByID retrieves a single account by its ID.
func (m *AccountModel) GetByID(ctx context.Context, id int) (*models.Account, error) {
	filters := []Filter{{Column: "account_id", Operator: Eq, Value: strconv.Itoa(id)}}
	query := Query{Table: TableAccount, Filter: filters}

	accounts, err := m.Get(ctx, &query)
	if err != nil {
		return nil, err
	}
	if len(accounts) == 1 {
		return &accounts[0], nil
	}

	return nil, ErrDBQuery
}

// Insert adds a new account to the database.
func (m *AccountModel) Insert(ctx context.Context, account models.Account) (int, error) {
	query := `
        INSERT INTO account (first_name, nickname, last_name, email, phone, balance, max_debt, category)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
        RETURNING account_id`
	var id int
	err := m.DB.QueryRow(ctx, query,
		account.FirstName,
		account.Nickname,
		account.LastName,
		account.Email,
		account.Phone,
		account.Balance,
		account.MaxDebt,
		account.Category).Scan(&id)

	return id, err
}

// Update modifies an existing account in the database.
// Note: This method does NOT update the Balance field.
func (m *AccountModel) Update(ctx context.Context, account models.Account) (int, error) {
	query := `
        UPDATE account
        SET first_name = $1, nickname = $2, last_name = $3, email = $4, phone = $5, max_debt = $6, category = $7, enabled = $8
        WHERE account_id = $9
        RETURNING account_id`
	var id int
	err := m.DB.QueryRow(ctx, query,
		account.FirstName,
		account.Nickname,
		account.LastName,
		account.Email,
		account.Phone,
		account.MaxDebt,
		account.Category,
		account.Enabled,
		account.ID).Scan(&id)

	return id, err
}

// Note: we do NOT support deletion of accounts
