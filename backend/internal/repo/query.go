package repo

import (
	"fmt"
	"strconv"
)

// tableColumns is the allowlist of valid column names per table.
var tableColumns = map[TableName]map[string]bool{
	TableAccount:           {"account_id": true, "first_name": true, "nickname": true, "last_name": true, "email": true, "phone": true, "balance": true, "max_debt": true, "category": true, "enabled": true, "created_at": true},
	TableAccountOption:     {"account_id": true, "key": true, "value": true},
	TableCategory:          {"category_id": true, "name": true, "enabled": true, "icon": true, "type": true},
	TableLocation:          {"location_id": true, "name": true},
	TableProduct:           {"product_id": true, "name": true, "price": true, "vat_id": true, "product_group_id": true, "size": true, "unit_id": true, "category_id": true, "created_at": true},
	TableProductGroup:      {"product_group_id": true, "name": true, "parent_id": true},
	TableProductVisibility: {"product_visibility_id": true, "category_id": true, "location_id": true, "product_id": true},
	TableUnit:              {"unit_id": true, "name": true},
	TableVat:               {"vat_id": true, "rate": true},
}

// ============================================================================
// Table names
// ============================================================================

type TableName string

const (
	TableAccount           TableName = "account"
	TableAccountOption     TableName = "account_option"
	TableCategory          TableName = "category"
	TableLocation          TableName = "location"
	TableOrder             TableName = "order"
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
		panic(fmt.Errorf("invalid Operation value: %v", o))
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
	allowed := tableColumns[q.Table]

	stmt := "SELECT * FROM " + string(q.Table)
	var args []any

	if len(q.Filter) > 0 {
		stmt += " WHERE "
		for i, f := range q.Filter {
			if !allowed[f.Column] {
				return "", nil, fmt.Errorf("%w: %q", ErrInvalidColumn, f.Column)
			}
			args = append(args, f.Value)
			placeholder := "$" + strconv.Itoa(len(args))
			if i > 0 {
				stmt += " AND "
			}
			stmt += f.Column + " " + f.Operator.SQLString() + " " + placeholder
		}
	}

	if q.Sort != nil {
		if !allowed[q.Sort.Column] {
			return "", nil, fmt.Errorf("%w: %q", ErrInvalidColumn, q.Sort.Column)
		}
		stmt += " ORDER BY " + q.Sort.Column + " " + q.Sort.Order.SQLString()
	}

	if q.Limit != nil {
		stmt += " LIMIT " + strconv.Itoa(*q.Limit)
	}

	return stmt, args, nil
}
