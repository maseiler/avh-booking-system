package models

type Operator string

const (
	Eq  = "eq"
	Gt  = "gt"
	Gte = "gte"
	Lt  = "lt"
	Lte = "lte"
	Ne  = "ne"
)

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
