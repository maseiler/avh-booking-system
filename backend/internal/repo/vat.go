package repo

import (
	"context"
	"strconv"

	"github.com/av-huette/avh-booking-system/internal/database"
	"github.com/av-huette/avh-booking-system/internal/models"
	"github.com/jackc/pgx/v5"
)

// VatModel provides database operations for Vat entities.
type VatModel struct {
	DB DBTx
}

// Get retrieves VATs based on the provided query specification.
func (m *VatModel) Get(ctx context.Context, query *Query) ([]models.Vat, error) {
	stmt, args, err := buildSelectSQL(query)
	if err != nil {
		return nil, err
	}
	rows, err := m.DB.Query(ctx, stmt, args...)
	if err != nil {
		return nil, err
	}

	vats, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Vat])
	if err != nil {
		return nil, err
	}

	return vats, nil
}

// GetByID retrieves a VAT rate by its ID.
func (m *VatModel) GetByID(ctx context.Context, id int) (*models.Vat, error) {
	query := Query{Table: TableVat, Filter: []Filter{{Column: "vat_id", Operator: Eq, Value: strconv.Itoa(id)}}}
	vats, err := m.Get(ctx, &query)
	if err != nil {
		return nil, err
	}
	if len(vats) == 0 {
		return nil, database.ErrNoRecord
	}
	return &vats[0], nil
}

// Insert adds a new VAT rate to the database.
func (m *VatModel) Insert(ctx context.Context, vat models.Vat) (int, error) {
	query := `
        INSERT INTO vat (rate)
        VALUES ($1)
        RETURNING vat_id;`
	var id int
	err := m.DB.QueryRow(ctx, query, vat.Rate).Scan(&id)

	return id, err
}
