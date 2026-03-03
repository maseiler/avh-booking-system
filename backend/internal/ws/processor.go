package ws

import (
	"encoding/json"
	"fmt"
	"github.com/av-huette/avh-booking-system/internal/models"
	"log/slog"
	"time"
)

func (s *Service) processPing(message Message) ([]byte, *WsError) {
	b, err := json.Marshal(message.Payload)
	if err != nil {
		return nil, &WsError{Code: WsBadInterface, Message: err.Error(), Details: "Marshal payload"}
	}

	var ping PingPong
	if err := json.Unmarshal(b, &ping); err != nil {
		return nil, &WsError{Code: WsBadJson, Message: err.Error(), Details: "Unmarshal to PingPong"}
	}

	s.log.Info("Received ping with timestamp", slog.String("timestamp", ping.Timestamp.String()))

	pong := PingPong{Timestamp: time.Now()}
	response := Message{Type: MsgTypePong, Payload: pong}
	b, _ = json.Marshal(response)

	err = s.validator.ValidateMessage(b)
	if err != nil {
		return nil, &WsError{Code: WsBadJson, Message: err.Error(), Details: "JSON does not comply with Message schema"}
	}

	return b, nil
}

// processQuery unmarshals the message, fetches the data from the database and returns the object as JSON
func (s *Service) processQuery(message Message) ([]byte, *WsError) {
	query, err := unmarshalInterface[models.Query](message.Payload)
	if err != nil {
		return nil, err
	}

	switch query.Table {
	case models.TableAccount:
		{
			accounts, _ := s.stores.Account.Get(query)
			var rawJsonSlice []json.RawMessage
			for _, account := range accounts {
				rawAccount, _ := json.Marshal(account)
				rawJsonSlice = append(rawJsonSlice, rawAccount)
			}

			return s.marshalQueryResultList(models.TableAccount, &rawJsonSlice)
		}

	case models.TableCategory:
		{
			categories, _ := s.stores.Category.Get(query)
			var rawJsonSlice []json.RawMessage
			for _, category := range categories {
				rawCategory, _ := json.Marshal(category)
				rawJsonSlice = append(rawJsonSlice, rawCategory)
			}

			return s.marshalQueryResultList(models.TableCategory, &rawJsonSlice)
		}

	case models.TableLocation:
		{
			locations, _ := s.stores.Location.Get(query)
			var rawJsonSlice []json.RawMessage
			for _, loc := range locations {
				rawLoc, _ := json.Marshal(loc)
				rawJsonSlice = append(rawJsonSlice, rawLoc)
			}

			return s.marshalQueryResultList(models.TableLocation, &rawJsonSlice)
		}

	case models.TableProduct:
		{
			products, _ := s.stores.Product.Get(query)
			var rawJsonSlice []json.RawMessage
			for _, product := range products {
				rawProduct, _ := json.Marshal(product)
				rawJsonSlice = append(rawJsonSlice, rawProduct)
			}

			return s.marshalQueryResultList(models.TableProduct, &rawJsonSlice)
		}

	case models.TableProductGroup:
		{
			groups, _ := s.stores.ProductGroup.Get(query)
			var rawJsonSlice []json.RawMessage
			for _, group := range groups {
				rawGroup, _ := json.Marshal(group)
				rawJsonSlice = append(rawJsonSlice, rawGroup)
			}

			return s.marshalQueryResultList(models.TableProductGroup, &rawJsonSlice)
		}

	case models.TableProductVisibility:
		{
			visibilities, _ := s.stores.ProductVisibility.Get(query)
			var rawJsonSlice []json.RawMessage
			for _, vis := range visibilities {
				rawVis, _ := json.Marshal(vis)
				rawJsonSlice = append(rawJsonSlice, rawVis)
			}

			return s.marshalQueryResultList(models.TableProductVisibility, &rawJsonSlice)
		}

	case models.TableUnit:
		{
			units, _ := s.stores.Unit.Get(query)
			var rawJsonSlice []json.RawMessage
			for _, unit := range units {
				rawUnit, _ := json.Marshal(unit)
				rawJsonSlice = append(rawJsonSlice, rawUnit)
			}

			return s.marshalQueryResultList(models.TableUnit, &rawJsonSlice)
		}

	case models.TableVat:
		{
			vats, _ := s.stores.Vat.Get(query)
			var rawJsonSlice []json.RawMessage
			for _, vat := range vats {
				rawVat, _ := json.Marshal(vat)
				rawJsonSlice = append(rawJsonSlice, rawVat)
			}

			return s.marshalQueryResultList(models.TableVat, &rawJsonSlice)
		}
	}

	return nil, &WsError{
		Code:    WsInvalidTable,
		Message: "Invalid table",
		Details: "Could not process query for table " + string(query.Table),
	}
}

// processMutation unmarshals the message and initiates a database mutation
func (s *Service) processMutation(mutation *models.Mutation) ([]byte, int, *WsError) {

	switch mutation.Operation {
	case models.OpInsert:
		{
			switch mutation.Table {
			case models.TableAccount:
				account, wsErr := unmarshalInterface[models.Account](mutation.Values)
				if wsErr != nil {
					return nil, 0, wsErr
				}

				newId, err := s.stores.Account.Insert(*account)
				if err != nil {
					return nil, 0, &WsError{Code: WsInternalError, Message: err.Error(), Details: "Could not create account"}
				}

				b, wsErr := s.marshalResultMutation(mutation.Table, mutation.Operation, newId)
				return b, newId, wsErr
			}

			return nil, 0, &WsError{
				Code:    WsInvalidTable,
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

			newId, err := s.stores.Account.Update(*account)
			if err != nil {
				return nil, 0, &WsError{Code: WsInternalError, Message: err.Error(), Details: "Could not update account"}
			}

			b, wsErr := s.marshalResultMutation(mutation.Table, mutation.Operation, newId)
			return b, newId, wsErr
		}

	case models.OpDelete:
		{
			return nil, 0, &WsError{Code: WsInvalidOperation,
				Message: "Invalid operation",
				Details: "Deletion of accounts is not supported"}
		}
	}

	return nil, 0, &WsError{
		Code:    WsInvalidOperation,
		Message: "Invalid operation",
		Details: "Could not process mutation for operation " + string(mutation.Operation),
	}
}

func (s *Service) prepareBroadcast(mutation *models.Mutation, id int) ([]byte, *WsError) {
	var data []byte
	switch mutation.Table {
	case models.TableAccount:
		var account *models.Account
		account, err := s.stores.Account.GetById(id)
		if err != nil {
			wsErr := &WsError{
				Code:    WsDbQueryError,
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
