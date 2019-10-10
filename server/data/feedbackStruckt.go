package data

import "time"

// Feedback represented as in database
type Feedback struct {
	ID        int
	TimeStamp time.Time
	Name      string
	Content   string
}
