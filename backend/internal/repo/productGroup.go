package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/av-huette/avh-booking-system/internal/database"
	"github.com/av-huette/avh-booking-system/internal/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// ProductGroupModel provides database operations for ProductGroup entities.
type ProductGroupModel struct {
	DB *database.DB
}

// Get retrieves product groups based on the provided query specification.
func (m *ProductGroupModel) Get(query *Query) ([]models.ProductGroup, error) {
	ctx := context.Background()
	stmt, args, err := buildSelectSQL(query)
	if err != nil {
		return nil, err
	}
	// pgx will panic when it tries to assign null to int as might be the case for the parent field. COALESCE in the SQL select statement will replace null values with 0.
	stmt = strings.Replace(stmt, "SELECT *", "SELECT product_group_id, name, COALESCE(parent, 0) AS parent", 1)
	rows, err := m.DB.Query(ctx, stmt, args...)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			fmt.Println(pgErr.Message)
			fmt.Println(pgErr.Code)
		}
	}

	groups, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.ProductGroup])
	if err != nil {
		fmt.Printf("CollectRows error: %v", err)
		return nil, err
	}

	return groups, nil
}

// GetById retrieves a product group by its ID.
func (m *ProductGroupModel) GetById(id int) (*models.ProductGroup, error) {
	ctx := context.Background()
	stmt := `SELECT product_group_id, name, parent
			FROM product_group
			WHERE product_group_id = $1`
	row := m.DB.QueryRow(ctx, stmt, id)

	var productGroup models.ProductGroup
	var parentId *int // pointer to read null
	err := row.Scan(&productGroup.ID, &productGroup.Name, &parentId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, database.ErrNoRecord
		} else {
			return nil, err
		}
	}
	if parentId == nil {
		productGroup.ParentID = 0
	} else {
		productGroup.ParentID = *parentId
	}

	return &productGroup, nil
}

// Insert adds a new product group to the database.
func (m *ProductGroupModel) Insert(group models.ProductGroup) (int, error) {
	ctx := context.Background()
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
