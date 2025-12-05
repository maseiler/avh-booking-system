package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/av-huette/avh-booking-system/internal/database"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

type Account struct {
	Id        int              `json:"id" db:"account_id"`
	FirstName string           `json:"firstName" db:"first_name"`
	Nickname  string           `json:"nickname" db:"nickname"`
	LastName  string           `json:"lastName" db:"last_name"`
	Email     string           `json:"email" db:"email"`
	Phone     string           `json:"phone" db:"phone"`
	Balance   pgtype.Numeric   `json:"balance" db:"balance"`
	MaxDebt   int              `json:"maxDebt" db:"max_debt"`
	Category  int              `json:"category" db:"category"`
	Enabled   bool             `json:"enabled" db:"enabled"`
	CreatedAt pgtype.Timestamp `json:"createdAt" db:"created_at"`
}

func CreateAccount(firstName string, nickName string, lastName string,
	email string, phone string, balance string, maxDebt int, category int) Account {
	return Account{
		FirstName: firstName,
		Nickname:  nickName,
		LastName:  lastName,
		Email:     email,
		Phone:     phone,
		Balance:   NewNumeric(balance),
		MaxDebt:   maxDebt,
		Category:  category,
		Enabled:   true,
	}
}

type AccountModel struct {
	DB *database.DB
}

func (m *AccountModel) Insert(account Account) (int, error) {
	ctx := context.Background()
	query := `
        INSERT INTO account (first_name, nickname, last_name, email, phone, balance, max_debt, category, enabled) 
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
        RETURNING account_id`
	var id int
	err := m.DB.QueryRow(ctx, query, account.FirstName,
		account.Nickname,
		account.LastName,
		account.Email,
		account.Phone,
		account.Balance,
		account.MaxDebt,
		account.Category,
		account.Enabled).Scan(&id)

	return id, err
}

func (m *AccountModel) Get(accountId int) (*Account, error) {
	ctx := context.Background()
	stmt := `SELECT account_id, first_name, nickname, last_name, email, phone, balance,
       max_debt, category, enabled, created_at FROM account WHERE account_id = $1`
	row := m.DB.QueryRow(ctx, stmt, accountId)

	var account Account
	err := row.Scan(&account.Id, &account.FirstName, &account.Nickname, &account.LastName, &account.Email,
		&account.Phone, &account.Balance, &account.MaxDebt, &account.Category, &account.Enabled, &account.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return &account, nil
}

func (m *AccountModel) GetAll() ([]Account, error) {
	ctx := context.Background()
	stmt := `SELECT * FROM account`
	rows, err := m.DB.Query(ctx, stmt)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			fmt.Println(pgErr.Message)
			fmt.Println(pgErr.Code)
		}
	}

	accounts, err := pgx.CollectRows(rows, pgx.RowToStructByName[Account])
	if err != nil {
		fmt.Printf("CollectRows error: %v", err)
		return nil, err
	}

	return accounts, nil
}
