package data

// Item represents an item as in database
type Item struct {
	ID      int
	Name    string
	Price   float32
	Size    float32
	Unit    string
	Type    string
	Enabled bool
}
