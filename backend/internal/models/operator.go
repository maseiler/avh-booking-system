package models

import (
	"encoding/json"
	"fmt"
)

type Operator uint

const (
	Eq Operator = iota + 1
	Gt
	Gte
	Lt
	Lte
	Ne
)

// String representations of operators
const (
	OperatorEq  = "eq"
	OperatorGt  = "gt"
	OperatorGte = "gte"
	OperatorLt  = "lt"
	OperatorLte = "lte"
	OperatorNe  = "ne"
)

// String returns the string representation of an operator
func (o Operator) String() string {
	switch o {
	case Eq:
		return OperatorEq
	case Gt:
		return OperatorGt
	case Gte:
		return OperatorGte
	case Lt:
		return OperatorLt
	case Lte:
		return OperatorLte
	case Ne:
		return OperatorNe
	default:
		return "unknown"
	}
}

// String returns the SQL string representation of an operator
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

// UnmarshalJSON implements the json.Unmarshaler interface
func (o *Operator) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	switch s {
	case OperatorEq:
		*o = Eq
	case OperatorGt:
		*o = Gt
	case OperatorGte:
		*o = Gte
	case OperatorLt:
		*o = Lt
	case OperatorLte:
		*o = Lte
	case OperatorNe:
		*o = Ne
	default:
		return fmt.Errorf("invalid operator: %s", s)
	}

	return nil
}

// MarshalJSON implements the json.Marshaler interface
func (o Operator) MarshalJSON() ([]byte, error) {
	var s string
	switch o {
	case Eq:
		s = OperatorEq
	case Gt:
		s = OperatorGt
	case Gte:
		s = OperatorGte
	case Lt:
		s = OperatorLt
	case Lte:
		s = OperatorLte
	case Ne:
		s = OperatorNe
	default:
		return nil, fmt.Errorf("invalid operator value: %d", o)
	}
	return json.Marshal(s)
}
