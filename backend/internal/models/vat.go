package models

import (
	"context"
	"database/sql"
	"errors"
	"github.com/av-huette/avh-booking-system/internal/database"
)

type Vat struct {
	Id   int `json:"id"`
	Rate int `json:"rate"`
}

type VatModel struct {
	DB *database.DB
}

func CreateVat(rate int) Vat {
	return Vat{
		Rate: rate,
	}
}

func (m *VatModel) Insert(vat Vat) (int, error) {
	ctx := context.Background()
	query := `
        INSERT INTO vat (rate) 
        VALUES ($1)
        RETURNING vat_id;`
	var id int
	err := m.DB.QueryRow(ctx, query, vat.Rate).Scan(&id)

	return id, err
}

func (m *VatModel) Get(id int) (*Vat, error) {
	ctx := context.Background()
	stmt := `SELECT vat_id, rate
			FROM vat
			WHERE vat_id = $1`
	row := m.DB.QueryRow(ctx, stmt, id)

	var vat Vat
	err := row.Scan(&vat.Id, &vat.Rate)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, database.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return &vat, nil
}
