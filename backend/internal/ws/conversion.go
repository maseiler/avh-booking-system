package ws

import (
	"encoding/json"
	"reflect"

	"github.com/av-huette/avh-booking-system/internal/models"
)

func unmarshalInterface[T any](payload interface{}) (*T, *models.WsError) {
	if payload == nil {
		return nil, &models.WsError{
			Code:    models.WsBadInterface,
			Message: "Payload is nil",
		}
	}

	// Convert the interface{} to JSON bytes
	var b []byte
	var err error
	if b, err = json.Marshal(payload); err != nil {
		return nil, &models.WsError{
			Code:    models.WsBadInterface,
			Message: err.Error(),
			Details: "Could not marshal interface",
		}
	}

	// Convert JSON bytes to target type T
	var result T
	if err := json.Unmarshal(b, &result); err != nil {
		return nil, &models.WsError{
			Code:    models.WsBadJson,
			Message: err.Error(),
			Details: "Could not unmarshal to type " + reflect.TypeOf(result).Name(),
		}
	}

	return &result, nil
}

func (s *Service) marshalResult(table models.TableName, data *[]json.RawMessage) ([]byte, *models.WsError) {
	// Create the message
	result := models.Result{Table: table, Data: *data}
	msg := models.Message{
		Type:    models.MsgTypeResult,
		Payload: result,
	}

	// Marshall message
	b, err := json.Marshal(msg)
	if err != nil {
		return nil, &models.WsError{Code: models.WsBadInterface, Message: err.Error(), Details: "Could not marshal message with Result containing []Account"}
	}

	// Validate message
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
		return nil, &models.WsError{Code: models.WsBadInterface, Message: err.Error(), Details: "Could not marshal message with ResultInsertion"}
	}

	// Validate message
	err = s.validator.ValidateMessage(b)
	if err != nil {
		return nil, &models.WsError{Code: models.WsBadJson, Message: err.Error(), Details: "JSON does not comply with Message schema"}
	}

	return b, nil
}
