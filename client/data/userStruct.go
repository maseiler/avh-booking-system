package data

import "time"

// User represents an user as defined in the database
type User struct {
	Id        int
	Creation  time.Time
	Client    string
	BierName  string
	FirstName string
	LastName  string
	BoatName  string
	Status    string
	Email     string
	Phone     string
	Balance   float32
	MaxDebt   int
}
