package ws

import (
	"encoding/json"
	"fmt"
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

			return s.marshalQueryResultList(models.TableAccount, &rawJsonSlice)
		}

	case models.TableCategory:
		{
			categories, _ := s.dbModels.Category.Get(query)
			var rawJsonSlice []json.RawMessage
			for _, category := range categories {
				rawCategory, _ := json.Marshal(category)
				rawJsonSlice = append(rawJsonSlice, rawCategory)
			}

			return s.marshalQueryResultList(models.TableCategory, &rawJsonSlice)
		}

	case models.TableProduct:
		{
			// TODO
		}

	case models.TableUnit:
		{
			units, _ := s.dbModels.Unit.Get(query)
			var rawJsonSlice []json.RawMessage
			for _, unit := range units {
				rawUnit, _ := json.Marshal(unit)
				rawJsonSlice = append(rawJsonSlice, rawUnit)
			}

			return s.marshalQueryResultList(models.TableUnit, &rawJsonSlice)
		}

	case models.TableVat:
		{
			vats, _ := s.dbModels.Vat.Get(query)
			var rawJsonSlice []json.RawMessage
			for _, vat := range vats {
				rawVat, _ := json.Marshal(vat)
				rawJsonSlice = append(rawJsonSlice, rawVat)
			}

			return s.marshalQueryResultList(models.TableVat, &rawJsonSlice)
		}
	}

	return nil, &models.WsError{
		Code:    models.WsInvalidTable,
		Message: "Invalid table",
		Details: "Could not process query for table " + string(query.Table),
	}
}

// processMutation unmarshals the message and initiates a database mutation
func (s *Service) processMutation(mutation *models.Mutation) ([]byte, int, *models.WsError) {

	switch mutation.Operation {
	case models.OpInsert:
		{
			switch mutation.Table {
			case models.TableAccount:
				account, wsErr := unmarshalInterface[models.Account](mutation.Values)
				if wsErr != nil {
					return nil, 0, wsErr
				}

				newId, err := s.dbModels.Account.Insert(*account)
				if err != nil {
					return nil, 0, &models.WsError{Code: models.WsInternalError, Message: err.Error(), Details: "Could not create account"}
				}

				b, wsErr := s.marshalResultMutation(mutation.Table, mutation.Operation, newId)
				return b, newId, wsErr
			}

			return nil, 0, &models.WsError{
				Code:    models.WsInvalidTable,
				Message: "Invalid table",
				Details: "Could not process mutation for table " + string(mutation.Table),
			}
		}
	case models.OpUpdate:
		{
			account, wsErr := unmarshalInterface[models.Account](mutation.Values)
			if wsErr != nil {
				return nil, 0, wsErr
			}

			newId, err := s.dbModels.Account.Update(*account)
			if err != nil {
				return nil, 0, &models.WsError{Code: models.WsInternalError, Message: err.Error(), Details: "Could not update account"}
			}

			b, wsErr := s.marshalResultMutation(mutation.Table, mutation.Operation, newId)
			return b, newId, wsErr
		}

	case models.OpDelete:
		{
			return nil, 0, &models.WsError{Code: models.WsInvalidOperation,
				Message: "Invalid operation",
				Details: "Deletion of accounts is not supported"}
		}
	}

	return nil, 0, &models.WsError{
		Code:    models.WsInvalidOperation,
		Message: "Invalid operation",
		Details: "Could not process mutation for operation " + string(mutation.Operation),
	}
}

func (s *Service) prepareBroadcast(mutation *models.Mutation, id int) ([]byte, *models.WsError) {
	var data []byte
	switch mutation.Table {
	case models.TableAccount:
		var account *models.Account
		account, err := s.dbModels.Account.GetById(id)
		if err != nil {
			wsErr := &models.WsError{
				Code:    models.WsDbQueryError,
				Message: err.Error(),
				Details: fmt.Sprintf("Could not get account with ID %d", id)}
			return nil, wsErr
		}
		data, _ = json.Marshal(account)
	}

	message, wsErr := s.marshalBroadcastWithPayload(mutation.Table, data)
	if wsErr != nil {
		return nil, wsErr
	}

	return message, nil
}
