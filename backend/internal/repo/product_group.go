package repo

import (
	"context"
	"strconv"
	"strings"

	"github.com/av-huette/avh-booking-system/internal/database"
	"github.com/av-huette/avh-booking-system/internal/models"
	"github.com/jackc/pgx/v5"
)

// ProductGroupModel provides database operations for ProductGroup entities.
type ProductGroupModel struct {
	DB *database.DB
}

// Get retrieves product groups based on the provided query specification.
func (m *ProductGroupModel) Get(ctx context.Context, query *Query) ([]models.ProductGroup, error) {
	stmt, args, err := buildSelectSQL(query)
	if err != nil {
		return nil, err
	}
	// pgx will panic when it tries to assign null to int as might be the case for the parent field. COALESCE in the SQL select statement will replace null values with 0.
	stmt = strings.Replace(stmt, "SELECT *", "SELECT product_group_id, name, COALESCE(parent, 0) AS parent", 1)
	rows, err := m.DB.Query(ctx, stmt, args...)
	if err != nil {
		return nil, err
	}

	groups, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.ProductGroup])
	if err != nil {
		return nil, err
	}

	return groups, nil
}

// GetByID retrieves a product group by its ID.
func (m *ProductGroupModel) GetByID(ctx context.Context, id int) (*models.ProductGroup, error) {
	query := Query{Table: TableProductGroup, Filter: []Filter{{Column: "product_group_id", Operator: Eq, Value: strconv.Itoa(id)}}}
	groups, err := m.Get(ctx, &query)
	if err != nil {
		return nil, err
	}
	if len(groups) == 0 {
		return nil, database.ErrNoRecord
	}
	return &groups[0], nil
}

// Insert adds a new product group to the database.
func (m *ProductGroupModel) Insert(ctx context.Context, group models.ProductGroup) (int, error) {
	var id int
	var err error
	if group.ParentID <= 0 {
		query := `
        INSERT INTO product_group (name)
        VALUES ($1)
        RETURNING product_group_id`
		err = m.DB.QueryRow(ctx, query, group.Name).Scan(&id)
	} else {
		query := `
        INSERT INTO product_group (name, parent)
        VALUES ($1, $2)
        RETURNING product_group_id`
		err = m.DB.QueryRow(ctx, query, group.Name, group.ParentID).Scan(&id)
	}

	return id, err
}
