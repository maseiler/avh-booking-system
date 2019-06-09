package data

// User represents an user as in database
type User struct {
	UserID    int
	BierName  string
	FirstName string
	LastName  string
	Status    string
	Email     string
	Balance   float32
	Phone     string
	MaxDebt   int
	BoatName  string
}
