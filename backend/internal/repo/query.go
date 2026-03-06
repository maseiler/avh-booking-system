package repo

import (
	"fmt"
	"strconv"
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

// SqlString returns the SQL string representation of an operator
func (o Operator) SqlString() string {
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

// SqlString returns the SQL string representation of an order direction
func (o Order) SqlString() string {
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

// SqlString returns the SQL string representation of an operation
func (o Operation) SqlString() string {
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

func buildSelectSQL(q *Query) string {
	stmt := "SELECT * FROM " + string(q.Table)

	if len(q.Filter) > 0 {
		stmt += " WHERE "
		for i, f := range q.Filter {
			stmt += f.Column + " " + f.Operator.SqlString() + " " + f.Value
			if i != len(q.Filter)-1 {
				stmt += " AND "
			}
		}
	}

	if q.Sort != nil {
		stmt += " ORDER BY " + q.Sort.Column + " " + q.Sort.Order.SqlString()
	}

	if q.Limit != nil {
		stmt += " LIMIT " + strconv.Itoa(*q.Limit)
	}

	return stmt
}
