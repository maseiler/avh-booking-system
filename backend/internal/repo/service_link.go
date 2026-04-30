package repo

import (
	"context"
	"strconv"

	"github.com/av-huette/avh-booking-system/internal/database"
	"github.com/av-huette/avh-booking-system/internal/models"
	"github.com/jackc/pgx/v5"
)

// ServiceLinkStore provides database operations for ServiceLink entities.
type ServiceLinkStore struct {
	DB DBTx
}

// Get retrieves service links based on the provided query specification.
func (m *ServiceLinkStore) Get(ctx context.Context, query *Query) ([]models.ServiceLink, error) {
	stmt, args, err := buildSelectSQL(query)
	if err != nil {
		return nil, err
	}
	rows, err := m.DB.Query(ctx, stmt, args...)
	if err != nil {
		return nil, err
	}

	entries, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.ServiceLink])
	if err != nil {
		return nil, err
	}

	return entries, nil
}

// GetByID retrieves a service link by its foreign user ID.
func (m *ServiceLinkStore) GetByID(ctx context.Context, foreignUserID int) (*models.ServiceLink, error) {
	query := Query{Table: TableServiceLink, Filter: []Filter{{Column: "foreign_user_id", Operator: Eq, Value: strconv.Itoa(foreignUserID)}}}
	entries, err := m.Get(ctx, &query)
	if err != nil {
		return nil, err
	}
	if len(entries) == 0 {
		return nil, database.ErrNoRecord
	}
	return &entries[0], nil
}

// Insert adds a new service link. ForeignUserID is provided by the caller.
func (m *ServiceLinkStore) Insert(ctx context.Context, s models.ServiceLink) error {
	query := `
        INSERT INTO service_link (foreign_user_id, "user", service_sewobe)
        VALUES ($1, $2, $3)`
	_, err := m.DB.Exec(ctx, query, s.ForeignUserID, s.UserID, s.ServiceSewobeID)
	return err
}

// Update modifies an existing service link.
func (m *ServiceLinkStore) Update(ctx context.Context, s models.ServiceLink) error {
	query := `
        UPDATE service_link
        SET "user" = $1, service_sewobe = $2
        WHERE foreign_user_id = $3`
	_, err := m.DB.Exec(ctx, query, s.UserID, s.ServiceSewobeID, s.ForeignUserID)
	return err
}

// Delete removes a service link by foreign user ID.
func (m *ServiceLinkStore) Delete(ctx context.Context, foreignUserID int) error {
	query := `DELETE FROM service_link WHERE foreign_user_id = $1`
	_, err := m.DB.Exec(ctx, query, foreignUserID)
	return err
}
