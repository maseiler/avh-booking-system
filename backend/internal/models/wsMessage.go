package models

import (
	"log"
	"strconv"
)

type Message struct {
	Type    string      `json:"type" validate:"required"`
	Payload interface{} `json:"payload,omitempty"`
}
type Query struct {
	Table  string   `json:"table" validate:"required"`
	Filter []Filter `json:"filter,omitempty"`
	Limit  int      `json:"limit,omitempty"`
	Sort   Sorting  `json:"sort,omitempty"`
}

// SqlStatement returns the SQL statement for this query
func (q *Query) SqlStatement() string {
	stmt := "SELECT * FROM " + q.Table

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

	if &q.Sort != nil {
		stmt += " ORDER BY " + q.Sort.Column + " " + q.Sort.Order.SqlString()
	}

	if q.Limit != 0 {
		stmt += " LIMIT " + strconv.Itoa(q.Limit)
	}

	log.Printf("Query statement: %s\n", stmt) // TODO debug mode

	return stmt
}

type Filter struct {
	Column   string   `json:"column" validate:"required"`
	Operator Operator `json:"operator" validate:"required"`
	//Operator string `json:"operator" validate:"required,validate_operator"`
	Value string `json:"value" validate:"required"`
}

type Sorting struct {
	Column string `json:"column" validate:"required"`
	Order  Order  `json:"order" validate:"required"`
}
