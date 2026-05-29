package repo

import (
	"context"
	"strconv"

	"github.com/av-huette/avh-booking-system/internal/database"
	"github.com/av-huette/avh-booking-system/internal/models"
	"github.com/jackc/pgx/v5"
)

// ServiceSewobeStore provides database operations for ServiceSewobe entities.
type ServiceSewobeStore struct {
	DB DBTx
}

// Get retrieves Sewobe service entries based on the provided query specification.
func (m *ServiceSewobeStore) Get(ctx context.Context, query *Query) ([]models.ServiceSewobe, error) {
	stmt, args, err := buildSelectSQL(query)
	if err != nil {
		return nil, err
	}
	rows, err := m.DB.Query(ctx, stmt, args...)
	if err != nil {
		return nil, err
	}

	entries, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.ServiceSewobe])
	if err != nil {
		return nil, err
	}

	return entries, nil
}

// GetByID retrieves a Sewobe service entry by its ID.
func (m *ServiceSewobeStore) GetByID(ctx context.Context, id int) (*models.ServiceSewobe, error) {
	query := Query{Table: TableServiceSewobe, Filter: []Filter{{Column: "service_sewobe_id", Operator: Eq, Value: strconv.Itoa(id)}}}
	entries, err := m.Get(ctx, &query)
	if err != nil {
		return nil, err
	}
	if len(entries) == 0 {
		return nil, database.ErrNoRecord
	}
	return &entries[0], nil
}

// Insert adds a new Sewobe service entry to the database.
func (m *ServiceSewobeStore) Insert(ctx context.Context, s models.ServiceSewobe) (int, error) {
	query := `
        INSERT INTO service_sewobe (sewobe_api_key, sewobe_url)
        VALUES ($1, $2)
        RETURNING service_sewobe_id`
	var id int
	err := m.DB.QueryRow(ctx, query, s.SewobeApiKey, s.SewobeUrl).Scan(&id)
	return id, err
}

// Update modifies an existing Sewobe service entry.
func (m *ServiceSewobeStore) Update(ctx context.Context, s models.ServiceSewobe) (int, error) {
	query := `
        UPDATE service_sewobe
        SET sewobe_api_key = $1, sewobe_url = $2
        WHERE service_sewobe_id = $3
        RETURNING service_sewobe_id`
	var id int
	err := m.DB.QueryRow(ctx, query, s.SewobeApiKey, s.SewobeUrl, s.ID).Scan(&id)
	return id, err
}

// Delete removes a Sewobe service entry by ID.
func (m *ServiceSewobeStore) Delete(ctx context.Context, id int) (int, error) {
	query := `DELETE FROM service_sewobe WHERE service_sewobe_id = $1 RETURNING service_sewobe_id`
	err := m.DB.QueryRow(ctx, query, id).Scan(&id)
	return id, err
}
