package database

import "errors"

var (
	ErrNoRecord = errors.New("models: no matching record found")
)
