package models

// ServiceSewobe holds credentials for the Sewobe external service.
type ServiceSewobe struct {
	ID           int    `json:"id" db:"service_sewobe_id"`
	SewobeApiKey string `json:"sewobeApiKey" db:"sewobe_api_key"`
	SewobeUrl    string `json:"sewobeUrl" db:"sewobe_url"`
}

// CreateServiceSewobe creates a new ServiceSewobe entry.
func CreateServiceSewobe(apiKey, url string) ServiceSewobe {
	return ServiceSewobe{SewobeApiKey: apiKey, SewobeUrl: url}
}
