package ws

import (
	"encoding/json"
	"github.com/av-huette/avh-booking-system/internal/models"
	"log/slog"
	"time"
)

func (s *Service) processPing(message models.Message) ([]byte, *models.WsError) {
	b, err := json.Marshal(message.Payload)
	if err != nil {
		return nil, &models.WsError{Code: models.WsBadInterface, Message: err.Error(), Details: "Marshal payload"}
	}

	var ping models.PingPong
	if err := json.Unmarshal(b, &ping); err != nil {
		return nil, &models.WsError{Code: models.WsBadJson, Message: err.Error(), Details: "Unmarshal to PingPong"}
	}

	s.log.Info("Received ping with timestamp", slog.String("timestamp", ping.Timestamp.String()))

	pong := models.PingPong{Timestamp: time.Now()}
	response := models.Message{Type: models.MsgTypePong, Payload: pong}
	b, _ = json.Marshal(response)

	err = s.validator.ValidateMessage(b)
	if err != nil {
		return nil, &models.WsError{Code: models.WsBadJson, Message: err.Error(), Details: "JSON does not comply with Message schema"}
	}

	return b, nil
}

// processQuery unmarshals the message, fetches the data from the database and returns the object as JSON
func (s *Service) processQuery(message models.Message) ([]byte, *models.WsError) {
	//query, err := getQuery(message.Payload)
	query, err := unmarshalInterface[models.Query](message.Payload)
	if err != nil {
		return nil, err
	}

	switch query.Table {
	case models.TableAccount:
		{
			accounts, _ := s.dbModels.Account.Get(query)
			var rawJsonSlice []json.RawMessage
			for _, account := range accounts {
				rawAccount, _ := json.Marshal(account)
				rawJsonSlice = append(rawJsonSlice, rawAccount)
			}

			return s.marshalQueryResult(models.TableAccount, &rawJsonSlice)
		}

	case models.TableProduct:
		// TODO
	}

	return nil, &models.WsError{
		Code:    models.WsInvalidTable,
		Message: "Invalid table",
		Details: "Could not process query for table " + string(query.Table),
	}
}

// processMutation unmarshals the message and initiates a database mutation
func (s *Service) processMutation(message models.Message) ([]byte, *models.WsError) {

	mutation, wsErr := unmarshalInterface[models.Mutation](message.Payload)
	if wsErr != nil {
		return nil, wsErr
	}

	switch mutation.Operation {
	case models.OpInsert:
		{
			switch mutation.Table {
			case models.TableAccount:
				account, wsErr := unmarshalInterface[models.Account](mutation.Values)
				if wsErr != nil {
					return nil, wsErr
				}

				newId, err := s.dbModels.Account.Insert(*account)
				if err != nil {
					return nil, &models.WsError{Code: models.WsInternalError, Message: err.Error(), Details: "Could not create account"}
				}

				return s.marshalResultMutation(mutation.Table, mutation.Operation, newId)
			}

			return nil, &models.WsError{
				Code:    models.WsInvalidTable,
				Message: "Invalid table",
				Details: "Could not process mutation for table " + string(mutation.Table),
			}
		}
	case models.OpUpdate:
		{
			//TODO
			return nil, &models.WsError{Code: models.WsUnknown, Message: "Not yet implemented", Details: "TBD"}
		}

	case models.OpDelete:
		{
			// TODO
			return nil, &models.WsError{Code: models.WsUnknown, Message: "Not yet implemented", Details: "TBD"}
		}
	}

	return nil, &models.WsError{
		Code:    models.WsInvalidOperation,
		Message: "Invalid operation",
		Details: "Could not process mutation for operation " + string(mutation.Operation),
	}
}
