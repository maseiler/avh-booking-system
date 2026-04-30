package repo

import (
	"context"

	"github.com/av-huette/avh-booking-system/internal/models"
	"github.com/jackc/pgx/v5"
)

// FavoritesStore provides database operations for Favorite entities.
type FavoritesStore struct {
	DB DBTx
}

// Get retrieves favorites based on the provided query specification.
func (m *FavoritesStore) Get(ctx context.Context, query *Query) ([]models.Favorite, error) {
	stmt, args, err := buildSelectSQL(query)
	if err != nil {
		return nil, err
	}
	rows, err := m.DB.Query(ctx, stmt, args...)
	if err != nil {
		return nil, err
	}

	favorites, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Favorite])
	if err != nil {
		return nil, err
	}

	return favorites, nil
}

// Insert adds a new favorite entry for an account and product.
func (m *FavoritesStore) Insert(ctx context.Context, fav models.Favorite) error {
	query := `
        INSERT INTO favorites (product, account, count)
        VALUES ($1, $2, $3)`
	_, err := m.DB.Exec(ctx, query, fav.ProductID, fav.AccountID, fav.Count)
	return err
}

// Update modifies the count of an existing favorite entry.
func (m *FavoritesStore) Update(ctx context.Context, fav models.Favorite) error {
	query := `
        UPDATE favorites
        SET count = $1
        WHERE product = $2 AND account = $3`
	_, err := m.DB.Exec(ctx, query, fav.Count, fav.ProductID, fav.AccountID)
	return err
}

// Delete removes a favorite entry by product ID and account ID.
func (m *FavoritesStore) Delete(ctx context.Context, productID, accountID int) error {
	query := `
        DELETE FROM favorites
        WHERE product = $1 AND account = $2`
	_, err := m.DB.Exec(ctx, query, productID, accountID)
	return err
}
