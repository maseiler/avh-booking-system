package data

import "time"

// Transaction represents a booking transaction as in database
type Transaction struct {
	transactionID int
	DateTime      time.Time
	UserID        int
	ItemID        int
	Amount        int
	TotalPrice    float32
}
