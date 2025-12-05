package models

import "errors"

var (
	ErrNoRecord  = errors.New("models: no matching record found")
	UnknownError = errors.New("Unknown error")
	BadJson      = errors.New("Bad Json")
)
