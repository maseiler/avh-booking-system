package ws

import (
	"encoding/json"
	"reflect"

	"github.com/av-huette/avh-booking-system/internal/models"
)

func unmarshalInterface[T any](payload interface{}) (*T, *WsError) {
	if payload == nil {
		return nil, &WsError{
			Code:    WsBadInterface,
			Message: "Payload is nil",
		}
	}

	// Convert the interface{} to JSON bytes
	var b []byte
	var err error
	if b, err = json.Marshal(payload); err != nil {
		return nil, &WsError{
			Code:    WsBadInterface,
			Message: err.Error(),
			Details: "Could not marshal interface",
		}
	}

	// Convert JSON bytes to target type T
	var result T
	if err := json.Unmarshal(b, &result); err != nil {
		return nil, &WsError{
			Code:    WsBadJson,
			Message: err.Error(),
			Details: "Could not unmarshal to type " + reflect.TypeOf(result).Name(),
		}
	}

	return &result, nil
}

func (s *Service) marshalAndValidateMessage(message *Message) ([]byte, *WsError) {
	// Marshall message
	b, err := json.Marshal(message)
	if err != nil {
		return nil, &WsError{Code: WsBadInterface, Message: err.Error(), Details: "Could not marshal message"}
	}

	// Validate message
	err = s.validator.ValidateMessage(b)
	if err != nil {
		return nil, &WsError{Code: WsBadJson, Message: err.Error(), Details: "JSON does not comply with Message schema"}
	}

	return b, nil
}

func (s *Service) marshalQueryResultList(table models.TableName, data *[]json.RawMessage) ([]byte, *WsError) {
	res := models.QueryResultList{Table: table, Data: *data}
	msg := Message{
		Type:    MsgTypeQueryResultList,
		Payload: res,
	}

	return s.marshalAndValidateMessage(&msg)
}

func (s *Service) marshalResultMutation(table models.TableName, operation models.Operation, id int) ([]byte, *WsError) {
	res := models.ResultMutation{Table: table, Operation: operation, Id: id}
	msg := Message{
		Type:    MsgTypeMutationResult,
		Payload: res,
	}

	return s.marshalAndValidateMessage(&msg)
}

func (s *Service) marshalBroadcastQueryResult(table models.TableName, data json.RawMessage) ([]byte, *WsError) {
	res := models.QueryResult{Table: table, Data: data}
	msg := Message{
		Type:    MsgTypeBroadcast,
		Payload: res,
	}

	return s.marshalAndValidateMessage(&msg)
}

func (s *Service) marshalBroadcastQueryResultList(table models.TableName, data *[]json.RawMessage) ([]byte, *WsError) {
	res := models.QueryResultList{Table: table, Data: *data}
	msg := Message{
		Type:    MsgTypeBroadcast,
		Payload: res,
	}

	return s.marshalAndValidateMessage(&msg)
}
