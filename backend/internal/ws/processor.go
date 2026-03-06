package ws

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"strconv"
	"time"

	"github.com/av-huette/avh-booking-system/internal/models"
	"github.com/av-huette/avh-booking-system/internal/repo"
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
	query, err := unmarshalInterface[repo.Query](message.Payload)
	if err != nil {
		return nil, err
	}

	switch query.Table {
	case repo.TableAccount:
		accounts, dbErr := s.stores.Account.Get(query)
		if dbErr != nil {
			return nil, &WsError{Code: WsDbQueryError, Message: dbErr.Error(), Details: "Query failed for table " + string(query.Table)}
		}
		var rawJsonSlice []json.RawMessage
		for _, account := range accounts {
			rawAccount, _ := json.Marshal(account)
			rawJsonSlice = append(rawJsonSlice, rawAccount)
		}
		return s.marshalQueryResultList(repo.TableAccount, &rawJsonSlice)

	case repo.TableCategory:
		categories, dbErr := s.stores.Category.Get(query)
		if dbErr != nil {
			return nil, &WsError{Code: WsDbQueryError, Message: dbErr.Error(), Details: "Query failed for table " + string(query.Table)}
		}
		var rawJsonSlice []json.RawMessage
		for _, category := range categories {
			rawCategory, _ := json.Marshal(category)
			rawJsonSlice = append(rawJsonSlice, rawCategory)
		}
		return s.marshalQueryResultList(repo.TableCategory, &rawJsonSlice)

	case repo.TableLocation:
		locations, dbErr := s.stores.Location.Get(query)
		if dbErr != nil {
			return nil, &WsError{Code: WsDbQueryError, Message: dbErr.Error(), Details: "Query failed for table " + string(query.Table)}
		}
		var rawJsonSlice []json.RawMessage
		for _, loc := range locations {
			rawLoc, _ := json.Marshal(loc)
			rawJsonSlice = append(rawJsonSlice, rawLoc)
		}
		return s.marshalQueryResultList(repo.TableLocation, &rawJsonSlice)

	case repo.TableProduct:
		products, dbErr := s.stores.Product.Get(query)
		if dbErr != nil {
			return nil, &WsError{Code: WsDbQueryError, Message: dbErr.Error(), Details: "Query failed for table " + string(query.Table)}
		}
		var rawJsonSlice []json.RawMessage
		for _, product := range products {
			rawProduct, _ := json.Marshal(product)
			rawJsonSlice = append(rawJsonSlice, rawProduct)
		}
		return s.marshalQueryResultList(repo.TableProduct, &rawJsonSlice)

	case repo.TableProductGroup:
		groups, dbErr := s.stores.ProductGroup.Get(query)
		if dbErr != nil {
			return nil, &WsError{Code: WsDbQueryError, Message: dbErr.Error(), Details: "Query failed for table " + string(query.Table)}
		}
		var rawJsonSlice []json.RawMessage
		for _, group := range groups {
			rawGroup, _ := json.Marshal(group)
			rawJsonSlice = append(rawJsonSlice, rawGroup)
		}
		return s.marshalQueryResultList(repo.TableProductGroup, &rawJsonSlice)

	case repo.TableProductVisibility:
		visibilities, dbErr := s.stores.ProductVisibility.Get(query)
		if dbErr != nil {
			return nil, &WsError{Code: WsDbQueryError, Message: dbErr.Error(), Details: "Query failed for table " + string(query.Table)}
		}
		var rawJsonSlice []json.RawMessage
		for _, vis := range visibilities {
			rawVis, _ := json.Marshal(vis)
			rawJsonSlice = append(rawJsonSlice, rawVis)
		}
		return s.marshalQueryResultList(repo.TableProductVisibility, &rawJsonSlice)

	case repo.TableUnit:
		units, dbErr := s.stores.Unit.Get(query)
		if dbErr != nil {
			return nil, &WsError{Code: WsDbQueryError, Message: dbErr.Error(), Details: "Query failed for table " + string(query.Table)}
		}
		var rawJsonSlice []json.RawMessage
		for _, unit := range units {
			rawUnit, _ := json.Marshal(unit)
			rawJsonSlice = append(rawJsonSlice, rawUnit)
		}
		return s.marshalQueryResultList(repo.TableUnit, &rawJsonSlice)

	case repo.TableVat:
		vats, dbErr := s.stores.Vat.Get(query)
		if dbErr != nil {
			return nil, &WsError{Code: WsDbQueryError, Message: dbErr.Error(), Details: "Query failed for table " + string(query.Table)}
		}
		var rawJsonSlice []json.RawMessage
		for _, vat := range vats {
			rawVat, _ := json.Marshal(vat)
			rawJsonSlice = append(rawJsonSlice, rawVat)
		}
		return s.marshalQueryResultList(repo.TableVat, &rawJsonSlice)
	}

	return nil, &WsError{
		Code:    WsInvalidTable,
		Message: "Invalid table",
		Details: "Could not process query for table " + string(query.Table),
	}
}

// processMutation unmarshals the message and initiates a database mutation
func (s *Service) processMutation(mutation *Mutation) ([]byte, int, *WsError) {

	switch mutation.Operation {
	case repo.OpInsert:
		{
			switch mutation.Table {
			case repo.TableAccount:
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

			case repo.TableProductVisibility:
				visibility, wsErr := unmarshalInterface[models.ProductVisibility](mutation.Values)
				if wsErr != nil {
					return nil, 0, wsErr
				}

				newId, err := s.stores.ProductVisibility.Insert(*visibility)
				if err != nil {
					return nil, 0, &WsError{Code: WsInternalError, Message: err.Error(), Details: "Could not create product visibility"}
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
	case repo.OpUpdate:
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

	case repo.OpDelete:
		{
			switch mutation.Table {
			case repo.TableAccount:
				{
					return nil, 0, &WsError{Code: WsInvalidOperation,
						Message: "Invalid operation",
						Details: "Deletion of accounts is not supported"}
				}

			case repo.TableProductVisibility:
				{
					idStr, ok := mutation.Where["product_visibility_id"]
					if !ok {
						return nil, 0, &WsError{
							Code:    WsBadJson,
							Message: "Missing product_visibility_id in where clause",
							Details: "Delete requires product_visibility_id",
						}
					}

					id, err := strconv.Atoi(idStr)
					if err != nil {
						return nil, 0, &WsError{
							Code:    WsBadJson,
							Message: err.Error(),
							Details: "product_visibility_id must be an integer",
						}
					}

					oldId, err := s.stores.ProductVisibility.Delete(id)
					if err != nil {
						return nil, 0, &WsError{Code: WsInternalError, Message: err.Error(), Details: "Could not delete visibility"}
					}

					b, wsErr := s.marshalResultMutation(mutation.Table, mutation.Operation, oldId)
					return b, oldId, wsErr

				}
			}

			return nil, 0, &WsError{
				Code:    WsInvalidTable,
				Message: "Invalid table",
				Details: "Could not process delete for table " + string(mutation.Table),
			}
		}
	}

	return nil, 0, &WsError{
		Code:    WsInvalidOperation,
		Message: "Invalid operation",
		Details: "Could not process mutation for operation " + string(mutation.Operation),
	}
}

func (s *Service) prepareBroadcast(mutation *Mutation, id int) ([]byte, *WsError) {
	var queryResultList []json.RawMessage
	switch mutation.Table {
	case repo.TableAccount:
		account, err := s.stores.Account.GetById(id)
		if err != nil {
			return nil, &WsError{
				Code:    WsDbQueryError,
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
			return nil, &WsError{
				Code:    WsDbQueryError,
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

	return nil, &WsError{Code: WsUnknown, Message: "Broadcast data is empty"}
}
