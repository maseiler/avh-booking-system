package repo

import (
	"context"
	"database/sql"
	"errors"

	"github.com/av-huette/avh-booking-system/internal/database"
	"github.com/av-huette/avh-booking-system/internal/models"
	"github.com/jackc/pgx/v5"
)

// VatModel provides database operations for Vat entities.
type VatModel struct {
	DB *database.DB
}

// Get retrieves VATs based on the provided query specification.
func (m *VatModel) Get(query *Query) ([]models.Vat, error) {
	ctx := context.Background()
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
func (m *VatModel) GetByID(id int) (*models.Vat, error) {
	ctx := context.Background()
	stmt := `SELECT vat_id, rate
			FROM vat
			WHERE vat_id = $1`
	row := m.DB.QueryRow(ctx, stmt, id)

	var vat models.Vat
	err := row.Scan(&vat.ID, &vat.Rate)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, database.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return &vat, nil
}

// Insert adds a new VAT rate to the database.
func (m *VatModel) Insert(vat models.Vat) (int, error) {
	ctx := context.Background()
	query := `
        INSERT INTO vat (rate)
        VALUES ($1)
        RETURNING vat_id;`
	var id int
	err := m.DB.QueryRow(ctx, query, vat.Rate).Scan(&id)

	return id, err
}
