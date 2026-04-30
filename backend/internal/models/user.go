package models

// User represents an authenticated system user.
// Password is stored as received — encryption is the caller's responsibility.
type User struct {
	ID       int    `json:"id" db:"user_id"`
	Name     string `json:"name" db:"name"`
	RoleID   int    `json:"roleId" db:"role"`
	Password string `json:"password" db:"password"`
}

// CreateUser creates a new User. Password must be pre-encrypted by the caller.
func CreateUser(name string, roleID int, password string) User {
	return User{Name: name, RoleID: roleID, Password: password}
}
