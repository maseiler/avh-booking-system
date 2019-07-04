package data

// User represents an user as in database
type User struct {
	ID        int
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
