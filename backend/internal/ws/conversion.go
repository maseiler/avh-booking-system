package ws

import (
	"encoding/json"
	"reflect"

	"github.com/av-huette/avh-booking-system/internal/repo"
)

func unmarshalInterface[T any](payload interface{}) (*T, *WSError) {
	if payload == nil {
		return nil, &WSError{
			Code:    WSBadInterface,
			Message: "Payload is nil",
		}
	}

	// Convert the interface{} to JSON bytes
	var b []byte
	var err error
	if b, err = json.Marshal(payload); err != nil {
		return nil, &WSError{
			Code:    WSBadInterface,
			Message: err.Error(),
			Details: "Could not marshal interface",
		}
	}

	// Convert JSON bytes to target type T
	var result T
	if err := json.Unmarshal(b, &result); err != nil {
		return nil, &WSError{
			Code:    WSBadJSON,
			Message: err.Error(),
			Details: "Could not unmarshal to type " + reflect.TypeOf(result).Name(),
		}
	}

	return &result, nil
}

func (s *Service) marshalAndValidateMessage(message *Message) ([]byte, *WSError) {
	// Marshall message
	b, err := json.Marshal(message)
	if err != nil {
		return nil, &WSError{Code: WSBadInterface, Message: err.Error(), Details: "Could not marshal message"}
	}

	// Validate message
	err = s.validator.validate(b)
	if err != nil {
		return nil, &WSError{Code: WSBadJSON, Message: err.Error(), Details: "JSON does not comply with Message schema"}
	}

	return b, nil
}

func (s *Service) marshalQueryResultList(table repo.TableName, data *[]json.RawMessage) ([]byte, *WSError) {
	res := QueryResultList{Table: table, Data: *data}
	msg := Message{
		Type:    MsgTypeQueryResultList,
		Payload: res,
	}

	return s.marshalAndValidateMessage(&msg)
}

func (s *Service) marshalResultMutation(table repo.TableName, operation repo.Operation, id int) ([]byte, *WSError) {
	res := ResultMutation{Table: table, Operation: operation, ID: id}
	msg := Message{
		Type:    MsgTypeMutationResult,
		Payload: res,
	}

	return s.marshalAndValidateMessage(&msg)
}

func (s *Service) marshalBroadcastQueryResult(table repo.TableName, data json.RawMessage) ([]byte, *WSError) {
	res := QueryResult{Table: table, Data: data}
	msg := Message{
		Type:    MsgTypeBroadcast,
		Payload: res,
	}

	return s.marshalAndValidateMessage(&msg)
}

func (s *Service) marshalBroadcastQueryResultList(table repo.TableName, data *[]json.RawMessage) ([]byte, *WSError) {
	res := QueryResultList{Table: table, Data: *data}
	msg := Message{
		Type:    MsgTypeBroadcast,
		Payload: res,
	}

	return s.marshalAndValidateMessage(&msg)
}
