package repo

import "errors"

// --------------------------------------------------
// Database errors
// --------------------------------------------------

var ErrDBQuery = errors.New("database query failed")
var ErrInvalidColumn = errors.New("invalid column name")
