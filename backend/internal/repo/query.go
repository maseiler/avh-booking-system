package repo

import (
	"strconv"

	"github.com/jackc/pgx/v5"
)

// ============================================================================
// Table names
// ============================================================================

type TableName string

const (
	TableAccount           TableName = "account"
	TableAccountOption     TableName = "account_option"
	TableCategory          TableName = "category"
	TableLocation          TableName = "location"
	TableFavorites         TableName = "favorites"
	TableServiceLink       TableName = "service_link"
	TableServiceSewobe     TableName = "service_sewobe"
	TableRights            TableName = "rights"
	TableRole              TableName = "role"
	TableUser              TableName = "user"
	TableUserOption        TableName = "user_option"
	TableSettingsEmail     TableName = "settings_email"
	TableSettingsFrontend  TableName = "settings_frontend"
	TableSettingsPayment   TableName = "settings_payment"
	TableOrder             TableName = "order"
	TableProductOrder      TableName = "product_order"
	TableProduct           TableName = "product"
	TableProductGroup      TableName = "product_group"
	TableProductVisibility TableName = "product_visibility"
	TableUnit              TableName = "unit"
	TableVat               TableName = "vat"
)

// ============================================================================
// Operator
// ============================================================================

type Operator string

const (
	Eq  = "eq"
	Gt  = "gt"
	Gte = "gte"
	Lt  = "lt"
	Lte = "lte"
	Ne  = "ne"
)

// SQLString returns the SQL string representation of an operator
func (o Operator) SQLString() string {
	switch o {
	case Eq:
		return "="
	case Gt:
		return ">"
	case Gte:
		return ">="
	case Lt:
		return "<"
	case Lte:
		return "<="
	case Ne:
		return "<>"
	default:
		return "unknown"
	}
}

// ============================================================================
// Order
// ============================================================================

type Order string

const (
	Asc  = "asc"
	Desc = "desc"
)

// SQLString returns the SQL string representation of an order direction
func (o Order) SQLString() string {
	switch o {
	case Asc:
		return "ASC"
	case Desc:
		return "DESC"
	default:
		return "unknown"
	}
}

// ============================================================================
// Operation
// ============================================================================

type Operation string

const (
	OpInsert = "insert"
	OpUpdate = "update"
	OpDelete = "delete"
)

// SQLString returns the SQL string representation of an operation
func (o Operation) SQLString() string {
	switch o {
	case OpInsert:
		return "INSERT"
	case OpUpdate:
		return "UPDATE"
	case OpDelete:
		return "DELETE"
	default:
		return "unknown"
	}
}

// ============================================================================
// Query types
// ============================================================================

type Query struct {
	Table  TableName `json:"table" validate:"required"`
	Filter []Filter  `json:"filter,omitempty"`
	Limit  *int      `json:"limit,omitempty"`
	Sort   *Sorting  `json:"sort,omitempty"`
}

type Filter struct {
	Column   string   `json:"column" validate:"required"`
	Operator Operator `json:"operator" validate:"required"`
	Value    string   `json:"value" validate:"required"`
}

type Sorting struct {
	Column string `json:"column" validate:"required"`
	Order  Order  `json:"order" validate:"required"`
}

// ============================================================================
// Query builder
// ============================================================================

func buildSelectSQL(q *Query) (string, []any, error) {
	stmt := "SELECT * FROM " + pgx.Identifier{string(q.Table)}.Sanitize()
	var args []any

	if len(q.Filter) > 0 {
		stmt += " WHERE "
		for i, f := range q.Filter {
			args = append(args, f.Value)
			placeholder := "$" + strconv.Itoa(len(args))
			if i > 0 {
				stmt += " AND "
			}
			stmt += pgx.Identifier{f.Column}.Sanitize() + " " + f.Operator.SQLString() + " " + placeholder
		}
	}

	if q.Sort != nil {
		stmt += " ORDER BY " + pgx.Identifier{q.Sort.Column}.Sanitize() + " " + q.Sort.Order.SQLString()
	}

	if q.Limit != nil {
		stmt += " LIMIT " + strconv.Itoa(*q.Limit)
	}

	return stmt, args, nil
}
