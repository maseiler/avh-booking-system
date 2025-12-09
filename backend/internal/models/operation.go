package models

import (
	"encoding/json"
	"fmt"
)

type Operation uint

const (
	OpInsert Operation = iota + 1
	OpUpdate
	OpDelete
)

// String representations of operations
const (
	OpInsertStr = "insert"
	OpUpdateStr = "update"
	OpDeleteStr = "delete"
)

// String returns the string representation of an operation
func (o Operation) String() string {
	switch o {
	case OpInsert:
		return OpInsertStr
	case OpUpdate:
		return OpUpdateStr
	case OpDelete:
		return OpDeleteStr
	default:
		return "unknown"
	}
}

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
		return "unknown"
	}
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (o *Operation) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	switch s {
	case OpInsertStr:
		*o = OpInsert
	case OpUpdateStr:
		*o = OpUpdate
	case OpDeleteStr:
		*o = OpDelete
	default:
		return fmt.Errorf("invalid operation: %s", s)
	}

	return nil
}

// MarshalJSON implements the json.Marshaler interface
func (o Operation) MarshalJSON() ([]byte, error) {
	var s string
	switch o {
	case OpInsert:
		s = OpInsertStr
	case OpUpdate:
		s = OpUpdateStr
	case OpDelete:
		s = OpDeleteStr
	default:
		return nil, fmt.Errorf("invalid operation value: %d", o)
	}
	return json.Marshal(s)
}
