package data

import "time"

// Debts represents the last payment and open debts for a user
type Debts struct {
	LastPayment time.Time
	Debts       []BookEntry
}
