package ws

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"strconv"
	"time"

	"github.com/av-huette/avh-booking-system/internal/models"
	"github.com/av-huette/avh-booking-system/internal/repo"
)

func queryError(table repo.TableName, err error) *WSError {
	code := WSErrorCode(WSDBQueryError)
	if errors.Is(err, repo.ErrInvalidColumn) {
		code = WSInvalidFilter
	}
	return &WSError{Code: code, Message: err.Error(), Details: "Query failed for table " + string(table)}
}

func (s *Service) processPing(message Message) ([]byte, *WSError) {
	b, err := json.Marshal(message.Payload)
	if err != nil {
		return nil, &WSError{Code: WSBadInterface, Message: err.Error(), Details: "Marshal payload"}
	}

	var ping PingPong
	if err := json.Unmarshal(b, &ping); err != nil {
		return nil, &WSError{Code: WSBadJSON, Message: err.Error(), Details: "Unmarshal to PingPong"}
	}

	s.log.Info("Received ping with timestamp", slog.String("timestamp", ping.Timestamp.String()))

	pong := PingPong{Timestamp: time.Now()}
	response := Message{Type: MsgTypePong, Payload: pong}
	b, _ = json.Marshal(response)

	err = s.validator.validate(b)
	if err != nil {
		return nil, &WSError{Code: WSBadJSON, Message: err.Error(), Details: "JSON does not comply with Message schema"}
	}

	return b, nil
}

// processQuery unmarshals the message, fetches the data from the database and returns the object as JSON
func (s *Service) processQuery(message Message) ([]byte, *WSError) {
	query, err := unmarshalInterface[repo.Query](message.Payload)
	if err != nil {
		return nil, err
	}

	switch query.Table {
	case repo.TableAccount:
		accounts, dbErr := s.stores.Account.Get(query)
		if dbErr != nil {
			return nil, queryError(query.Table, dbErr)
		}
		var rawJSONSlice []json.RawMessage
		for _, account := range accounts {
			rawAccount, _ := json.Marshal(account)
			rawJSONSlice = append(rawJSONSlice, rawAccount)
		}
		return s.marshalQueryResultList(repo.TableAccount, &rawJSONSlice)

	case repo.TableCategory:
		categories, dbErr := s.stores.Category.Get(query)
		if dbErr != nil {
			return nil, queryError(query.Table, dbErr)
		}
		var rawJSONSlice []json.RawMessage
		for _, category := range categories {
			rawCategory, _ := json.Marshal(category)
			rawJSONSlice = append(rawJSONSlice, rawCategory)
		}
		return s.marshalQueryResultList(repo.TableCategory, &rawJSONSlice)

	case repo.TableLocation:
		locations, dbErr := s.stores.Location.Get(query)
		if dbErr != nil {
			return nil, queryError(query.Table, dbErr)
		}
		var rawJSONSlice []json.RawMessage
		for _, loc := range locations {
			rawLoc, _ := json.Marshal(loc)
			rawJSONSlice = append(rawJSONSlice, rawLoc)
		}
		return s.marshalQueryResultList(repo.TableLocation, &rawJSONSlice)

	case repo.TableProduct:
		products, dbErr := s.stores.Product.Get(query)
		if dbErr != nil {
			return nil, queryError(query.Table, dbErr)
		}
		var rawJSONSlice []json.RawMessage
		for _, product := range products {
			rawProduct, _ := json.Marshal(product)
			rawJSONSlice = append(rawJSONSlice, rawProduct)
		}
		return s.marshalQueryResultList(repo.TableProduct, &rawJSONSlice)

	case repo.TableProductGroup:
		groups, dbErr := s.stores.ProductGroup.Get(query)
		if dbErr != nil {
			return nil, queryError(query.Table, dbErr)
		}
		var rawJSONSlice []json.RawMessage
		for _, group := range groups {
			rawGroup, _ := json.Marshal(group)
			rawJSONSlice = append(rawJSONSlice, rawGroup)
		}
		return s.marshalQueryResultList(repo.TableProductGroup, &rawJSONSlice)

	case repo.TableProductVisibility:
		visibilities, dbErr := s.stores.ProductVisibility.Get(query)
		if dbErr != nil {
			return nil, queryError(query.Table, dbErr)
		}
		var rawJSONSlice []json.RawMessage
		for _, vis := range visibilities {
			rawVis, _ := json.Marshal(vis)
			rawJSONSlice = append(rawJSONSlice, rawVis)
		}
		return s.marshalQueryResultList(repo.TableProductVisibility, &rawJSONSlice)

	case repo.TableUnit:
		units, dbErr := s.stores.Unit.Get(query)
		if dbErr != nil {
			return nil, queryError(query.Table, dbErr)
		}
		var rawJSONSlice []json.RawMessage
		for _, unit := range units {
			rawUnit, _ := json.Marshal(unit)
			rawJSONSlice = append(rawJSONSlice, rawUnit)
		}
		return s.marshalQueryResultList(repo.TableUnit, &rawJSONSlice)

	case repo.TableVat:
		vats, dbErr := s.stores.Vat.Get(query)
		if dbErr != nil {
			return nil, queryError(query.Table, dbErr)
		}
		var rawJSONSlice []json.RawMessage
		for _, vat := range vats {
			rawVat, _ := json.Marshal(vat)
			rawJSONSlice = append(rawJSONSlice, rawVat)
		}
		return s.marshalQueryResultList(repo.TableVat, &rawJSONSlice)
	}

	return nil, &WSError{
		Code:    WSInvalidTable,
		Message: "Invalid table",
		Details: "Could not process query for table " + string(query.Table),
	}
}

// processMutation unmarshals the message and initiates a database mutation
func (s *Service) processMutation(mutation *Mutation) ([]byte, int, *WSError) {

	switch mutation.Operation {
	case repo.OpInsert:
		{
			switch mutation.Table {
			case repo.TableAccount:
				account, wsErr := unmarshalInterface[models.Account](mutation.Values)
				if wsErr != nil {
					return nil, 0, wsErr
				}

				newID, err := s.stores.Account.Insert(*account)
				if err != nil {
					return nil, 0, &WSError{Code: WSInternalError, Message: err.Error(), Details: "Could not create account"}
				}

				b, wsErr := s.marshalResultMutation(mutation.Table, mutation.Operation, newID)
				return b, newID, wsErr

			case repo.TableProductVisibility:
				visibility, wsErr := unmarshalInterface[models.ProductVisibility](mutation.Values)
				if wsErr != nil {
					return nil, 0, wsErr
				}

				newID, err := s.stores.ProductVisibility.Insert(*visibility)
				if err != nil {
					return nil, 0, &WSError{Code: WSInternalError, Message: err.Error(), Details: "Could not create product visibility"}
				}

				b, wsErr := s.marshalResultMutation(mutation.Table, mutation.Operation, newID)
				return b, newID, wsErr
			}

			return nil, 0, &WSError{
				Code:    WSInvalidTable,
				Message: "Invalid table",
				Details: "Could not process mutation for table " + string(mutation.Table),
			}
		}
	case repo.OpUpdate:
		{
			account, wsErr := unmarshalInterface[models.Account](mutation.Values)
			if wsErr != nil {
				return nil, 0, wsErr
			}

			newID, err := s.stores.Account.Update(*account)
			if err != nil {
				return nil, 0, &WSError{Code: WSInternalError, Message: err.Error(), Details: "Could not update account"}
			}

			b, wsErr := s.marshalResultMutation(mutation.Table, mutation.Operation, newID)
			return b, newID, wsErr
		}

	case repo.OpDelete:
		{
			switch mutation.Table {
			case repo.TableAccount:
				{
					return nil, 0, &WSError{Code: WSInvalidOperation,
						Message: "Invalid operation",
						Details: "Deletion of accounts is not supported"}
				}

			case repo.TableProductVisibility:
				{
					idStr, ok := mutation.Where["product_visibility_id"]
					if !ok {
						return nil, 0, &WSError{
							Code:    WSBadJSON,
							Message: "Missing product_visibility_id in where clause",
							Details: "Delete requires product_visibility_id",
						}
					}

					id, err := strconv.Atoi(idStr)
					if err != nil {
						return nil, 0, &WSError{
							Code:    WSBadJSON,
							Message: err.Error(),
							Details: "product_visibility_id must be an integer",
						}
					}

					oldID, err := s.stores.ProductVisibility.Delete(id)
					if err != nil {
						return nil, 0, &WSError{Code: WSInternalError, Message: err.Error(), Details: "Could not delete visibility"}
					}

					b, wsErr := s.marshalResultMutation(mutation.Table, mutation.Operation, oldID)
					return b, oldID, wsErr

				}
			}

			return nil, 0, &WSError{
				Code:    WSInvalidTable,
				Message: "Invalid table",
				Details: "Could not process delete for table " + string(mutation.Table),
			}
		}
	}

	return nil, 0, &WSError{
		Code:    WSInvalidOperation,
		Message: "Invalid operation",
		Details: "Could not process mutation for operation " + string(mutation.Operation),
	}
}

func (s *Service) prepareBroadcast(mutation *Mutation, id int) ([]byte, *WSError) {
	var queryResultList []json.RawMessage
	switch mutation.Table {
	case repo.TableAccount:
		account, err := s.stores.Account.GetByID(id)
		if err != nil {
			return nil, &WSError{
				Code:    WSDBQueryError,
				Message: err.Error(),
				Details: fmt.Sprintf("Could not get account with ID %d", id),
			}
		}
		rawAcc, _ := json.Marshal(account)
		queryResultList = append(queryResultList, rawAcc)

	case repo.TableProductVisibility:
		query := repo.Query{
			Table: repo.TableProductVisibility,
		}
		visibilities, err := s.stores.ProductVisibility.Get(&query)
		if err != nil {
			return nil, &WSError{
				Code:    WSDBQueryError,
				Message: err.Error(),
				Details: "Could not reload product visibilities after delete",
			}
		}
		for _, vis := range visibilities {
			rawVis, _ := json.Marshal(vis)
			queryResultList = append(queryResultList, rawVis)
		}

	} // end of switch

	if len(queryResultList) == 1 {
		message, wsErr := s.marshalBroadcastQueryResult(mutation.Table, queryResultList[0])
		if wsErr != nil {
			return nil, wsErr
		}
		return message, nil
	} else if len(queryResultList) > 1 {
		message, wsErr := s.marshalBroadcastQueryResultList(mutation.Table, &queryResultList)
		if wsErr != nil {
			return nil, wsErr
		}
		return message, nil
	}

	return nil, &WSError{Code: WSUnknown, Message: "Broadcast data is empty"}
}
