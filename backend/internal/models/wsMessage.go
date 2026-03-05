package models

import (
	"encoding/json"
	"strconv"
)

// ============================================================================
// Database Access/Manipulation Types and Protocol
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

type QueryResult struct {
	Table TableName       `json:"table" validate:"required"`
	Data  json.RawMessage `json:"data,omitempty"`
}

type QueryResultList struct {
	Table TableName         `json:"table" validate:"required"`
	Data  []json.RawMessage `json:"data,omitempty"`
}

type ResultMutation struct {
	Table     TableName `json:"table" validate:"required"`
	Operation Operation `json:"operation" validate:"required"`
	Id        int       `json:"id" validate:"required"`
}

type Query struct {
	Table  TableName `json:"table" validate:"required"`
	Filter []Filter  `json:"filter,omitempty"`
	Limit  *int      `json:"limit,omitempty"`
	Sort   *Sorting  `json:"sort,omitempty"`
}

// SqlStatement returns the SQL statement for this query
func (q *Query) SqlStatement() string {
	stmt := "SELECT * FROM " + string(q.Table)

	if len(q.Filter) > 0 {
		stmt += " WHERE "
		for i, filter := range q.Filter {
			stmt += filter.Column + " "
			stmt += filter.Operator.SqlString() + " "
			stmt += filter.Value
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

type Filter struct {
	Column   string   `json:"column" validate:"required"`
	Operator Operator `json:"operator" validate:"required"`
	Value    string   `json:"value" validate:"required"`
}

type Sorting struct {
	Column string `json:"column" validate:"required"`
	Order  Order  `json:"order" validate:"required"`
}

type Mutation struct {
	Operation Operation         `json:"operation" validate:"required"`
	Table     TableName         `json:"table" validate:"required"`
	Where     map[string]string `json:"where"`
	Values    interface{}       `json:"values"`
}
