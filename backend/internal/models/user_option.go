package models

// UserOption represents a key/value option associated with a user.
type UserOption struct {
	UserID int    `json:"userId" db:"user_id"`
	Key    string `json:"key" db:"key"`
	Value  string `json:"value" db:"value"`
}

// CreateUserOption creates a new UserOption for the given user.
func CreateUserOption(userID int, key, value string) UserOption {
	return UserOption{UserID: userID, Key: key, Value: value}
}
