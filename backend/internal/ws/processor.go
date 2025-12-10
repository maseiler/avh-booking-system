package ws

import (
	"encoding/json"
	"github.com/av-huette/avh-booking-system/internal/models"
)

// processQuery unmarshals the message, fetches the data from the database and returns the object as JSON
func (s *Service) processQuery(message models.Message) ([]byte, *models.WsError) {
	query, err := getQuery(message.Payload)
	if err != nil {
		return nil, err
	}

	switch query.Table {
	case models.TableAccount:
		{
			accounts, _ := s.dbModels.Account.Get(query)
			msg := models.Message{
				Type:    models.MsgTypeResponse,
				Payload: map[string]interface{}{"accounts": accounts},
			}
			b, err := json.Marshal(msg)
			if err != nil {
				return nil, &models.WsError{Code: models.WsBadInterface, Message: err.Error(), Details: "Could not marshal response including Account"}
			}
			return b, nil
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

	mutation, err := getMutation(message.Payload)
	if err != nil {
		return nil, err
	}

	switch mutation.Operation {
	case models.OpInsert:
		{
			switch mutation.Table {
			case models.TableAccount:
				jsonBytes, err := json.Marshal(mutation.Values)
				if err != nil {
					return nil, &models.WsError{Code: models.WsBadInterface, Message: err.Error(), Details: "Could not marshal interface"}
				}

				// Convert the interface{} back to JSON bytes
				account := models.Account{}
				err = json.Unmarshal(jsonBytes, &account)
				if err != nil {
					return nil, &models.WsError{Code: models.WsBadJson, Message: err.Error(), Details: "Could not unmarshal JSON"}
				}
				row, err := s.dbModels.Account.Insert(account)
				if err != nil {
					return nil, &models.WsError{Code: models.WsInternalError, Message: err.Error(), Details: "Could not create account"}
				}
				msg := models.Message{
					Type:    models.MsgTypeResponse,
					Payload: map[string]interface{}{"row": row},
				}
				b, err := json.Marshal(msg)
				if err != nil {
					return nil, &models.WsError{Code: models.WsBadInterface, Message: err.Error(), Details: "Could not marshal response"}
				}
				return b, nil
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
