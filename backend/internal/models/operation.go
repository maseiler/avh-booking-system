package models

import (
	"fmt"
)

type Operation string

const (
	OpInsert = "insert"
	OpUpdate = "update"
	OpDelete = "delete"
)

// String returns the SQL string representation of an operation
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
