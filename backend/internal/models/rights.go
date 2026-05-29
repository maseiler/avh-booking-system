package models

// Rights represents a permission granted or denied to a role.
type Rights struct {
	ID         int    `json:"id" db:"rights_id"`
	RoleID     int    `json:"roleId" db:"role"`
	Permission string `json:"permission" db:"permission"`
	Allowed    bool   `json:"allowed" db:"allowed"`
}

// CreateRights creates a new Rights entry for the given role.
func CreateRights(roleID int, permission string, allowed bool) Rights {
	return Rights{RoleID: roleID, Permission: permission, Allowed: allowed}
}
