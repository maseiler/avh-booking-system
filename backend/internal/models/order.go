package models

type Order string

const (
	Asc  = "asc"
	Desc = "desc"
)

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
