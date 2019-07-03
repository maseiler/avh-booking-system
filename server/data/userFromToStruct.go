package data

// UserFromTo is needed to unmarshal a user and two strings from API
type UserFromTo struct {
	User User
	From string
	To   string
}
