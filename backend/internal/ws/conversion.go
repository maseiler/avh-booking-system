package ws

import (
	"encoding/json"

	"github.com/av-huette/avh-booking-system/internal/models"
	"github.com/go-playground/validator/v10"
)

var structValidator = validator.New(validator.WithRequiredStructEnabled())

func getQuery(payload interface{}) (*models.Query, *models.WsError) {
	b, err := json.Marshal(payload)
	if err != nil {
		return nil, &models.WsError{Code: models.WsBadInterface, Message: err.Error(), Details: "Marshal payload"}
	}

	var q models.Query
	if err := json.Unmarshal(b, &q); err != nil {
		return nil, &models.WsError{Code: models.WsBadJson, Message: err.Error(), Details: "Unmarshal to Query"}
	}

	if err := structValidator.Struct(q); err != nil {
		return nil, &models.WsError{Code: models.WsBadStruct, Message: err.Error(), Details: "Query struct invalid"}
	}

	return &q, nil
}

func getMutation(payload interface{}) (*models.Mutation, *models.WsError) {
	b, err := json.Marshal(payload)
	if err != nil {
		return nil, &models.WsError{Code: models.WsBadInterface, Message: err.Error(), Details: "Marshal payload"}
	}

	var m models.Mutation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, &models.WsError{Code: models.WsBadJson, Message: err.Error(), Details: "Unmarshal to Mutation"}
	}

	if err := structValidator.Struct(m); err != nil {
		return nil, &models.WsError{Code: models.WsBadStruct, Message: err.Error(), Details: "Mutation struct invalid"}
	}

	return &m, nil
}
