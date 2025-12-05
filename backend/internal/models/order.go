package models

import (
	"encoding/json"
	"fmt"
)

type Order uint

const (
	Asc Order = iota + 1
	Desc
)

// String representations of order direct
const (
	OrderAsc  = "asc"
	OrderDesc = "desc"
)

// String returns the string representation of an order diretction
func (o Order) String() string {
	switch o {
	case Asc:
		return OrderAsc
	case Desc:
		return OrderDesc
	default:
		return "unknown"
	}
}

// String returns the SQL string representation of an order diretction
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

// UnmarshalJSON implements the json.Unmarshaler interface
func (o *Order) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	switch s {
	case OrderAsc:
		*o = Asc
	case OrderDesc:
		*o = Desc
	default:
		return fmt.Errorf("invalid order direction: %s", s)
	}

	return nil
}

// MarshalJSON implements the json.Marshaler interface
func (o Order) MarshalJSON() ([]byte, error) {
	var s string
	switch o {
	case Asc:
		s = OrderAsc
	case Desc:
		s = OrderDesc
	default:
		return nil, fmt.Errorf("invalid order direction value: %d", o)
	}
	return json.Marshal(s)
}
