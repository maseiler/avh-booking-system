package models

// AccountOption represents a key-value configuration option for an account.
// It allows storing arbitrary account-specific settings and preferences.
// The combination of AccountID and Key forms a composite primary key.
type AccountOption struct {
	AccountID int    `json:"accountId" db:"account"`
	Key       string `json:"key" db:"key"`
	Value     string `json:"value" db:"value"`
}

// CreateAccountOption creates a new AccountOption with the provided details.
// This is a convenience function for constructing AccountOption structs.
func CreateAccountOption(accountID int, key string, value string) AccountOption {
	return AccountOption{
		AccountID: accountID,
		Key:       key,
		Value:     value,
	}
}
