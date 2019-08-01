package data

import "time"

// BookEntry represents a booking transaction as in database
// type BookEntry struct {
// 	ID         int       `json:",string"`
// 	TimeStamp  time.Time `json:",string"`
// 	UserID     int       `json:",string"`
// 	ItemID     int       `json:",string"`
// 	Amount     int       `json:",string"`
// 	TotalPrice float32   `json:",string"`
// 	Comment    string
// }

// BookEntry represents a booking transaction as in database
type BookEntry struct {
	ID         int
	TimeStamp  time.Time
	UserID     int
	ItemID     int
	Amount     int
	TotalPrice float32
	Comment    string
}
