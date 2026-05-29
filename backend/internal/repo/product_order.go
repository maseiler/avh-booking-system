package repo

import (
	"context"

	"github.com/av-huette/avh-booking-system/internal/models"
	"github.com/jackc/pgx/v5"
)

// ProductOrderStore provides database operations for ProductOrder entities.
type ProductOrderStore struct {
	DB DBTx
}

// Get retrieves product order line items based on the provided query specification.
func (m *ProductOrderStore) Get(ctx context.Context, query *Query) ([]models.ProductOrder, error) {
	stmt, args, err := buildSelectSQL(query)
	if err != nil {
		return nil, err
	}
	rows, err := m.DB.Query(ctx, stmt, args...)
	if err != nil {
		return nil, err
	}

	items, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.ProductOrder])
	if err != nil {
		return nil, err
	}

	return items, nil
}

// Insert adds a new line item to an order.
func (m *ProductOrderStore) Insert(ctx context.Context, item models.ProductOrder) error {
	query := `
        INSERT INTO product_order (order_id, product, amount, product_price)
        VALUES ($1, $2, $3, $4)`
	_, err := m.DB.Exec(ctx, query, item.OrderID, item.ProductID, item.Amount, item.ProductPrice)
	return err
}

// Delete removes a line item from an order by order ID and product ID.
func (m *ProductOrderStore) Delete(ctx context.Context, orderID, productID int) error {
	query := `
        DELETE FROM product_order
        WHERE order_id = $1 AND product = $2`
	_, err := m.DB.Exec(ctx, query, orderID, productID)
	return err
}
