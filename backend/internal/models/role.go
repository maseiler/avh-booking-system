package models

// Role represents a user role used for access control.
type Role struct {
	ID   int    `json:"id" db:"role_id"`
	Name string `json:"name" db:"name"`
}

// CreateRole creates a new Role with the given name.
func CreateRole(name string) Role {
	return Role{Name: name}
}
