package models

// ServiceLink links an external service user ID to a local user
// and their corresponding accounts in external services.
type ServiceLink struct {
	ForeignUserID   int `json:"foreignUserId" db:"foreign_user_id"`
	UserID          int `json:"userId" db:"user"`
	ServiceSewobeID int `json:"serviceSewobeId" db:"service_sewobe"`
}

// CreateServiceLink creates a new ServiceLink mapping.
func CreateServiceLink(foreignUserID, userID, serviceSewobeID int) ServiceLink {
	return ServiceLink{
		ForeignUserID:   foreignUserID,
		UserID:          userID,
		ServiceSewobeID: serviceSewobeID,
	}
}
