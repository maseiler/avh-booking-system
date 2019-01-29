package data

// User represents an user as in database
type User struct {
	UserID    int
	BierName  string
	FirstName string
	LastName  string
	Status    string
	Phone     string
	Email     string
	Balance   float32
}
