package ws

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"strconv"
	"time"

	"github.com/av-huette/avh-booking-system/internal/models"
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

	case models.TableLocation:
		{
			locations, _ := s.dbModels.Location.Get(query)
			var rawJsonSlice []json.RawMessage
			for _, loc := range locations {
				rawLoc, _ := json.Marshal(loc)
				rawJsonSlice = append(rawJsonSlice, rawLoc)
			}

			return s.marshalQueryResultList(models.TableLocation, &rawJsonSlice)
		}

	case models.TableProduct:
		{
			products, _ := s.dbModels.Product.Get(query)
			var rawJsonSlice []json.RawMessage
			for _, product := range products {
				rawProduct, _ := json.Marshal(product)
				rawJsonSlice = append(rawJsonSlice, rawProduct)
			}

			return s.marshalQueryResultList(models.TableProduct, &rawJsonSlice)
		}

	case models.TableProductGroup:
		{
			groups, _ := s.dbModels.ProductGroup.Get(query)
			var rawJsonSlice []json.RawMessage
			for _, group := range groups {
				rawGroup, _ := json.Marshal(group)
				rawJsonSlice = append(rawJsonSlice, rawGroup)
			}

			return s.marshalQueryResultList(models.TableProductGroup, &rawJsonSlice)
		}

	case models.TableProductVisibility:
		{
			visibilities, _ := s.dbModels.ProductVisibility.Get(query)
			var rawJsonSlice []json.RawMessage
			for _, vis := range visibilities {
				rawVis, _ := json.Marshal(vis)
				rawJsonSlice = append(rawJsonSlice, rawVis)
			}

			return s.marshalQueryResultList(models.TableProductVisibility, &rawJsonSlice)
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

			case models.TableProductVisibility:
				visibility, wsErr := unmarshalInterface[models.ProductVisibility](mutation.Values)
				if wsErr != nil {
					return nil, 0, wsErr
				}

				newId, err := s.dbModels.ProductVisibility.Insert(*visibility)
				if err != nil {
					return nil, 0, &models.WsError{Code: models.WsInternalError, Message: err.Error(), Details: "Could not create product visibility"}
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
			switch mutation.Table {
			case models.TableAccount:
				{
					return nil, 0, &models.WsError{Code: models.WsInvalidOperation,
						Message: "Invalid operation",
						Details: "Deletion of accounts is not supported"}
				}

			case models.TableProductVisibility:
				{
					idStr, ok := mutation.Where["product_visibility_id"]
					if !ok {
						return nil, 0, &models.WsError{
							Code:    models.WsBadJson,
							Message: "Missing product_visibility_id in where clause",
							Details: "Delete requires product_visibility_id",
						}
					}

					id, err := strconv.Atoi(idStr)
					if err != nil {
						return nil, 0, &models.WsError{
							Code:    models.WsBadJson,
							Message: err.Error(),
							Details: "product_visibility_id must be an integer",
						}
					}

					oldId, err := s.dbModels.ProductVisibility.Delete(id)
					if err != nil {
						return nil, 0, &models.WsError{Code: models.WsInternalError, Message: err.Error(), Details: "Could not delete visibility"}
					}

					b, wsErr := s.marshalResultMutation(mutation.Table, mutation.Operation, oldId)
					return b, oldId, wsErr

				}
			}

			return nil, 0, &models.WsError{
				Code:    models.WsInvalidTable,
				Message: "Invalid table",
				Details: "Could not process delete for table " + string(mutation.Table),
			}
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

	case models.TableProductVisibility:
		query := models.Query{
			Table: models.TableProductVisibility,
		}
		visibilities, err := s.dbModels.ProductVisibility.Get(&query)
		if err != nil {
			return nil, &models.WsError{
				Code:    models.WsDbQueryError,
				Message: err.Error(),
				Details: "Could not reload product visibilities after delete",
			}
		}
		data, _ = json.Marshal(visibilities)
	}

	message, wsErr := s.marshalBroadcastWithPayload(mutation.Table, data)
	if wsErr != nil {
		return nil, wsErr
	}

	return message, nil
}
