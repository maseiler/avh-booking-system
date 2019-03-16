package data

import "time"

// BookEntry represents a booking transaction as in database
type BookEntry struct {
	BookEntryID int       `json:",string"`
	TimeStamp   time.Time `json:",string"`
	UserID      int       `json:",string"`
	ItemID      int       `json:",string"`
	Amount      int       `json:",string"`
	TotalPrice  float32   `json:",string"`
	Comment     string
}
