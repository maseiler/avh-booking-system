package repo

import (
	"context"
	"errors"

	"github.com/av-huette/avh-booking-system/internal/database"
	"github.com/av-huette/avh-booking-system/internal/models"
	"github.com/jackc/pgx/v5"
)

// SettingsStore provides database operations for key/value settings tables.
// Set the Table field to TableSettingsFrontend, TableSettingsPayment, or TableSettingsEmail.
type SettingsStore struct {
	DB    DBTx
	Table TableName
}

// GetAll retrieves all entries from the settings table.
func (m *SettingsStore) GetAll(ctx context.Context) ([]models.Setting, error) {
	query := Query{Table: m.Table}
	stmt, args, err := buildSelectSQL(&query)
	if err != nil {
		return nil, err
	}
	rows, err := m.DB.Query(ctx, stmt, args...)
	if err != nil {
		return nil, err
	}

	settings, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Setting])
	if err != nil {
		return nil, err
	}

	return settings, nil
}

// Get retrieves a single setting by key.
func (m *SettingsStore) Get(ctx context.Context, key string) (*models.Setting, error) {
	stmt := "SELECT * FROM " + pgx.Identifier{string(m.Table)}.Sanitize() + " WHERE key = $1"
	row := m.DB.QueryRow(ctx, stmt, key)

	var s models.Setting
	err := row.Scan(&s.Key, &s.Value)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, database.ErrNoRecord
		}
		return nil, err
	}

	return &s, nil
}

// Insert adds a new setting entry.
func (m *SettingsStore) Insert(ctx context.Context, setting models.Setting) error {
	stmt := "INSERT INTO " + pgx.Identifier{string(m.Table)}.Sanitize() + " (key, value) VALUES ($1, $2)"
	_, err := m.DB.Exec(ctx, stmt, setting.Key, setting.Value)
	return err
}

// Update modifies the value of an existing setting.
func (m *SettingsStore) Update(ctx context.Context, setting models.Setting) error {
	stmt := "UPDATE " + pgx.Identifier{string(m.Table)}.Sanitize() + " SET value = $1 WHERE key = $2"
	_, err := m.DB.Exec(ctx, stmt, setting.Value, setting.Key)
	return err
}

// Delete removes a setting by key.
func (m *SettingsStore) Delete(ctx context.Context, key string) error {
	stmt := "DELETE FROM " + pgx.Identifier{string(m.Table)}.Sanitize() + " WHERE key = $1"
	_, err := m.DB.Exec(ctx, stmt, key)
	return err
}
