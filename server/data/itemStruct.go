package data

// Item represents an item as in database
type Item struct {
	// ItemID int             `json:"ItemID"`
	// Name   string          `json:"Name"`
	// Price  float32         `json:"Price"`
	// Size   sql.NullFloat64 `json:"Size"`
	// Unit   sql.NullString  `json:"Unit"`
	ItemID int
	Name   string
	Price  float32
	Size   float32
	Unit   string
	Type   string
}
