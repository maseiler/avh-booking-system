package models

// Setting represents a key/value configuration entry.
// It is shared across all settings tables (frontend, payment, email).
type Setting struct {
	Key   string `json:"key" db:"key"`
	Value string `json:"value" db:"value"`
}

// CreateSetting creates a new Setting with the given key and value.
func CreateSetting(key, value string) Setting {
	return Setting{Key: key, Value: value}
}
