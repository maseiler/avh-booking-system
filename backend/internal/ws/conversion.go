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

func unmarshalAccount(mutationValues *interface{}) (*models.Account, *models.WsError) {
	// Convert the interface{} to JSON bytes
	b, err := json.Marshal(mutationValues)
	if err != nil {
		return nil, &models.WsError{Code: models.WsBadInterface, Message: err.Error(), Details: "Could not marshal interface"}
	}

	// Convert JSON bytes to Account
	account := &models.Account{}
	err = json.Unmarshal(b, &account)
	if err != nil {
		return nil, &models.WsError{Code: models.WsBadJson, Message: err.Error(), Details: "Could not unmarshal JSON to Account"}
	}

	return account, nil
}

func (s *Service) marshalResult(table models.TableName, data *[]json.RawMessage) ([]byte, *models.WsError) {
	result := models.Result{Table: table, Data: *data}

	msg := models.Message{
		Type:    models.MsgTypeResult,
		Payload: result,
	}

	b, err := json.Marshal(msg)
	if err != nil {
		return nil, &models.WsError{Code: models.WsBadInterface, Message: err.Error(), Details: "Could not marshal response including Account"}
	}

	err = s.validator.ValidateMessage(b)
	if err != nil {
		return nil, &models.WsError{Code: models.WsBadJson, Message: err.Error(), Details: "JSON does not comply with Message schema"}
	}

	return b, nil
}

func (s *Service) marshalResultInsertion(id int) ([]byte, *models.WsError) {
	// Create the message
	resInsertion := models.ResultInsertion{Id: id}
	msg := models.Message{
		Type:    models.MsgTypeResultInsertion,
		Payload: resInsertion,
	}

	// Marshall message
	b, err := json.Marshal(msg)
	if err != nil {
		return nil, &models.WsError{Code: models.WsBadInterface, Message: err.Error(), Details: "Could not marshal ResultInsertion"}
	}

	// Validate message
	err = s.validator.ValidateMessage(b)
	if err != nil {
		return nil, &models.WsError{Code: models.WsBadJson, Message: err.Error(), Details: "JSON does not comply with Message schema"}
	}

	return b, nil
}
