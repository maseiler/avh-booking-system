package ws

import (
	"encoding/json"

	"github.com/av-huette/avh-booking-system/internal/repo"
)

type QueryResult struct {
	Table repo.TableName  `json:"table" validate:"required"`
	Data  json.RawMessage `json:"data,omitempty"`
}

type QueryResultList struct {
	Table repo.TableName    `json:"table" validate:"required"`
	Data  []json.RawMessage `json:"data,omitempty"`
}

type ResultMutation struct {
	Table     repo.TableName `json:"table" validate:"required"`
	Operation repo.Operation `json:"operation" validate:"required"`
	ID        int            `json:"id" validate:"required"`
}

type Mutation struct {
	Operation repo.Operation    `json:"operation" validate:"required"`
	Table     repo.TableName    `json:"table" validate:"required"`
	Where     map[string]string `json:"where"`
	Values    interface{}       `json:"values"`
}
