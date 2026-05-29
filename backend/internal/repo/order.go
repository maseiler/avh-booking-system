package repo

import (
	"context"
	"strconv"

	"github.com/av-huette/avh-booking-system/internal/database"
	"github.com/av-huette/avh-booking-system/internal/models"
	"github.com/jackc/pgx/v5"
)

// OrderStore provides database operations for Order entities.
type OrderStore struct {
	DB DBTx
}

// Get retrieves orders based on the provided query specification.
func (m *OrderStore) Get(ctx context.Context, query *Query) ([]models.Order, error) {
	stmt, args, err := buildSelectSQL(query)
	if err != nil {
		return nil, err
	}
	rows, err := m.DB.Query(ctx, stmt, args...)
	if err != nil {
		return nil, err
	}

	orders, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Order])
	if err != nil {
		return nil, err
	}

	return orders, nil
}

// GetByID retrieves an order by its ID.
func (m *OrderStore) GetByID(ctx context.Context, id int) (*models.Order, error) {
	query := Query{Table: TableOrder, Filter: []Filter{{Column: "order_id", Operator: Eq, Value: strconv.Itoa(id)}}}
	orders, err := m.Get(ctx, &query)
	if err != nil {
		return nil, err
	}
	if len(orders) == 0 {
		return nil, database.ErrNoRecord
	}
	return &orders[0], nil
}

// Insert adds a new order to the database.
func (m *OrderStore) Insert(ctx context.Context, order models.Order) (int, error) {
	query := `
        INSERT INTO "order" (account)
        VALUES ($1)
        RETURNING order_id`
	var id int
	err := m.DB.QueryRow(ctx, query, order.AccountID).Scan(&id)
	return id, err
}
