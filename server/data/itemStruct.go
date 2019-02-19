package data

// Item represents an item as in database
type Item struct {
	ItemID int
	Name   string
	Price  float32 `json:",string"`
	Size   float32 `json:",string"`
	Unit   string
	Type   string
}
