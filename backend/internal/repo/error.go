package repo

import "errors"

// --------------------------------------------------
// Database errors
// --------------------------------------------------

var DbQueryError = errors.New("database query failed")
var ErrInvalidColumn = errors.New("invalid column name")
