package data

import "time"

// BookEntry represents a booking transaction as in database
type BookEntry struct {
	ID            int
	TimeStamp     time.Time
	UserID        int
	ItemID        int
	Amount        int
	TotalPrice    float32
	Comment       string
	PaymentMethod string
}
