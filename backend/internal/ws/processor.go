package ws

import (
	"context"
	"encoding/json"
	"log/slog"
	"strconv"
	"time"

	"github.com/av-huette/avh-booking-system/internal/repo"
)

func queryError(table repo.TableName, err error) *WSError {
	return &WSError{Code: WSDBQueryError, Message: err.Error(), Details: "Query failed for table " + string(table)}
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

	s.log.Debug("Received ping with timestamp", slog.String("timestamp", ping.Timestamp.String()))

	pong := PingPong{Timestamp: time.Now()}
	response := Message{Type: MsgTypePong, Payload: pong}
	b, err = json.Marshal(response)
	if err != nil {
		return nil, &WSError{Code: WSInternalError, Message: err.Error(), Details: "Failed to encode pong response"}
	}

	err = s.validator.validate(b)
	if err != nil {
		return nil, &WSError{Code: WSBadJSON, Message: err.Error(), Details: "JSON does not comply with Message schema"}
	}

	return b, nil
}

func (s *Service) processQuery(ctx context.Context, message Message) ([]byte, *WSError) {
	query, err := unmarshalInterface[repo.Query](message.Payload)
	if err != nil {
		return nil, err
	}

	switch query.Table {
	case repo.TableAccount:
		return processGetQuery(ctx, s, query, s.stores.Account.Get)
	case repo.TableCategory:
		return processGetQuery(ctx, s, query, s.stores.Category.Get)
	case repo.TableFavorites:
		return processGetQuery(ctx, s, query, s.stores.Favorites.Get)
	case repo.TableLocation:
		return processGetQuery(ctx, s, query, s.stores.Location.Get)
	case repo.TableOrder:
		return processGetQuery(ctx, s, query, s.stores.Order.Get)
	case repo.TableProduct:
		return processGetQuery(ctx, s, query, s.stores.Product.Get)
	case repo.TableProductGroup:
		return processGetQuery(ctx, s, query, s.stores.ProductGroup.Get)
	case repo.TableProductOrder:
		return processGetQuery(ctx, s, query, s.stores.ProductOrder.Get)
	case repo.TableProductVisibility:
		return processGetQuery(ctx, s, query, s.stores.ProductVisibility.Get)
	case repo.TableRights:
		return processGetQuery(ctx, s, query, s.stores.Rights.Get)
	case repo.TableRole:
		return processGetQuery(ctx, s, query, s.stores.Role.Get)
	case repo.TableServiceLink:
		return processGetQuery(ctx, s, query, s.stores.ServiceLink.Get)
	case repo.TableServiceSewobe:
		return processGetQuery(ctx, s, query, s.stores.ServiceSewobe.Get)
	case repo.TableSettingsFrontend:
		return processGetAllQuery(ctx, s, repo.TableSettingsFrontend, s.stores.SettingsFrontend.GetAll)
	case repo.TableSettingsPayment:
		return processGetAllQuery(ctx, s, repo.TableSettingsPayment, s.stores.SettingsPayment.GetAll)
	case repo.TableSettingsEmail:
		return processGetAllQuery(ctx, s, repo.TableSettingsEmail, s.stores.SettingsEmail.GetAll)
	case repo.TableUnit:
		return processGetQuery(ctx, s, query, s.stores.Unit.Get)
	case repo.TableUser:
		return processGetQuery(ctx, s, query, s.stores.User.Get)
	case repo.TableVat:
		return processGetQuery(ctx, s, query, s.stores.Vat.Get)
	}

	return nil, &WSError{
		Code:    WSInvalidTable,
		Message: "Invalid table",
		Details: "Could not process query for table " + string(query.Table),
	}
}

func (s *Service) processMutation(ctx context.Context, mutation *Mutation) ([]byte, func() ([]byte, *WSError), *WSError) {
	switch mutation.Operation {
	case repo.OpInsert:
		switch mutation.Table {
		case repo.TableAccount:
			b, id, wsErr := processInsert(ctx, s, mutation, s.stores.Account.Insert)
			if wsErr != nil {
				return nil, nil, wsErr
			}
			return b, func() ([]byte, *WSError) {
				return broadcastByID(ctx, s, repo.TableAccount, id, s.stores.Account.GetByID)
			}, nil
		case repo.TableCategory:
			b, id, wsErr := processInsert(ctx, s, mutation, s.stores.Category.Insert)
			if wsErr != nil {
				return nil, nil, wsErr
			}
			return b, func() ([]byte, *WSError) {
				return broadcastByID(ctx, s, repo.TableCategory, id, s.stores.Category.GetByID)
			}, nil
		case repo.TableFavorites:
			b, _, wsErr := processInsertNoReturn(ctx, s, mutation, 0, s.stores.Favorites.Insert)
			if wsErr != nil {
				return nil, nil, wsErr
			}
			return b, func() ([]byte, *WSError) { return broadcastList(ctx, s, repo.TableFavorites, s.stores.Favorites.Get) }, nil
		case repo.TableLocation:
			b, id, wsErr := processInsert(ctx, s, mutation, s.stores.Location.Insert)
			if wsErr != nil {
				return nil, nil, wsErr
			}
			return b, func() ([]byte, *WSError) {
				return broadcastByID(ctx, s, repo.TableLocation, id, s.stores.Location.GetByID)
			}, nil
		case repo.TableOrder:
			b, id, wsErr := processInsert(ctx, s, mutation, s.stores.Order.Insert)
			if wsErr != nil {
				return nil, nil, wsErr
			}
			return b, func() ([]byte, *WSError) { return broadcastByID(ctx, s, repo.TableOrder, id, s.stores.Order.GetByID) }, nil
		case repo.TableProduct:
			b, id, wsErr := processInsert(ctx, s, mutation, s.stores.Product.Insert)
			if wsErr != nil {
				return nil, nil, wsErr
			}
			return b, func() ([]byte, *WSError) {
				return broadcastByID(ctx, s, repo.TableProduct, id, s.stores.Product.GetByID)
			}, nil
		case repo.TableProductGroup:
			b, id, wsErr := processInsert(ctx, s, mutation, s.stores.ProductGroup.Insert)
			if wsErr != nil {
				return nil, nil, wsErr
			}
			return b, func() ([]byte, *WSError) {
				return broadcastByID(ctx, s, repo.TableProductGroup, id, s.stores.ProductGroup.GetByID)
			}, nil
		case repo.TableProductOrder:
			b, _, wsErr := processInsertNoReturn(ctx, s, mutation, 0, s.stores.ProductOrder.Insert)
			if wsErr != nil {
				return nil, nil, wsErr
			}
			return b, func() ([]byte, *WSError) {
				return broadcastList(ctx, s, repo.TableProductOrder, s.stores.ProductOrder.Get)
			}, nil
		case repo.TableProductVisibility:
			b, id, wsErr := processInsert(ctx, s, mutation, s.stores.ProductVisibility.Insert)
			if wsErr != nil {
				return nil, nil, wsErr
			}
			return b, func() ([]byte, *WSError) {
				return broadcastByID(ctx, s, repo.TableProductVisibility, id, s.stores.ProductVisibility.GetByID)
			}, nil
		case repo.TableRights:
			b, id, wsErr := processInsert(ctx, s, mutation, s.stores.Rights.Insert)
			if wsErr != nil {
				return nil, nil, wsErr
			}
			return b, func() ([]byte, *WSError) { return broadcastByID(ctx, s, repo.TableRights, id, s.stores.Rights.GetByID) }, nil
		case repo.TableRole:
			b, id, wsErr := processInsert(ctx, s, mutation, s.stores.Role.Insert)
			if wsErr != nil {
				return nil, nil, wsErr
			}
			return b, func() ([]byte, *WSError) { return broadcastByID(ctx, s, repo.TableRole, id, s.stores.Role.GetByID) }, nil
		case repo.TableServiceLink:
			b, _, wsErr := processInsertNoReturn(ctx, s, mutation, 0, s.stores.ServiceLink.Insert)
			if wsErr != nil {
				return nil, nil, wsErr
			}
			return b, func() ([]byte, *WSError) {
				return broadcastList(ctx, s, repo.TableServiceLink, s.stores.ServiceLink.Get)
			}, nil
		case repo.TableServiceSewobe:
			b, id, wsErr := processInsert(ctx, s, mutation, s.stores.ServiceSewobe.Insert)
			if wsErr != nil {
				return nil, nil, wsErr
			}
			return b, func() ([]byte, *WSError) {
				return broadcastByID(ctx, s, repo.TableServiceSewobe, id, s.stores.ServiceSewobe.GetByID)
			}, nil
		case repo.TableSettingsFrontend:
			b, _, wsErr := processInsertNoReturn(ctx, s, mutation, 0, s.stores.SettingsFrontend.Insert)
			if wsErr != nil {
				return nil, nil, wsErr
			}
			return b, func() ([]byte, *WSError) {
				return broadcastGetAllList(ctx, s, repo.TableSettingsFrontend, s.stores.SettingsFrontend.GetAll)
			}, nil
		case repo.TableSettingsPayment:
			b, _, wsErr := processInsertNoReturn(ctx, s, mutation, 0, s.stores.SettingsPayment.Insert)
			if wsErr != nil {
				return nil, nil, wsErr
			}
			return b, func() ([]byte, *WSError) {
				return broadcastGetAllList(ctx, s, repo.TableSettingsPayment, s.stores.SettingsPayment.GetAll)
			}, nil
		case repo.TableSettingsEmail:
			b, _, wsErr := processInsertNoReturn(ctx, s, mutation, 0, s.stores.SettingsEmail.Insert)
			if wsErr != nil {
				return nil, nil, wsErr
			}
			return b, func() ([]byte, *WSError) {
				return broadcastGetAllList(ctx, s, repo.TableSettingsEmail, s.stores.SettingsEmail.GetAll)
			}, nil
		case repo.TableUnit:
			b, id, wsErr := processInsert(ctx, s, mutation, s.stores.Unit.Insert)
			if wsErr != nil {
				return nil, nil, wsErr
			}
			return b, func() ([]byte, *WSError) { return broadcastByID(ctx, s, repo.TableUnit, id, s.stores.Unit.GetByID) }, nil
		case repo.TableUser:
			b, id, wsErr := processInsert(ctx, s, mutation, s.stores.User.Insert)
			if wsErr != nil {
				return nil, nil, wsErr
			}
			return b, func() ([]byte, *WSError) { return broadcastByID(ctx, s, repo.TableUser, id, s.stores.User.GetByID) }, nil
		case repo.TableVat:
			b, id, wsErr := processInsert(ctx, s, mutation, s.stores.Vat.Insert)
			if wsErr != nil {
				return nil, nil, wsErr
			}
			return b, func() ([]byte, *WSError) { return broadcastByID(ctx, s, repo.TableVat, id, s.stores.Vat.GetByID) }, nil
		}

	case repo.OpUpdate:
		switch mutation.Table {
		case repo.TableAccount:
			b, id, wsErr := processUpdate(ctx, s, mutation, s.stores.Account.Update)
			if wsErr != nil {
				return nil, nil, wsErr
			}
			return b, func() ([]byte, *WSError) {
				return broadcastByID(ctx, s, repo.TableAccount, id, s.stores.Account.GetByID)
			}, nil
		case repo.TableCategory:
			b, id, wsErr := processUpdate(ctx, s, mutation, s.stores.Category.Update)
			if wsErr != nil {
				return nil, nil, wsErr
			}
			return b, func() ([]byte, *WSError) {
				return broadcastByID(ctx, s, repo.TableCategory, id, s.stores.Category.GetByID)
			}, nil
		case repo.TableFavorites:
			b, _, wsErr := processUpdateNoReturn(ctx, s, mutation, 0, s.stores.Favorites.Update)
			if wsErr != nil {
				return nil, nil, wsErr
			}
			return b, func() ([]byte, *WSError) { return broadcastList(ctx, s, repo.TableFavorites, s.stores.Favorites.Get) }, nil
		case repo.TableProduct:
			b, id, wsErr := processUpdate(ctx, s, mutation, s.stores.Product.Update)
			if wsErr != nil {
				return nil, nil, wsErr
			}
			return b, func() ([]byte, *WSError) {
				return broadcastByID(ctx, s, repo.TableProduct, id, s.stores.Product.GetByID)
			}, nil
		case repo.TableProductGroup:
			b, id, wsErr := processUpdate(ctx, s, mutation, s.stores.ProductGroup.Update)
			if wsErr != nil {
				return nil, nil, wsErr
			}
			return b, func() ([]byte, *WSError) {
				return broadcastByID(ctx, s, repo.TableProductGroup, id, s.stores.ProductGroup.GetByID)
			}, nil
		case repo.TableRights:
			b, id, wsErr := processUpdate(ctx, s, mutation, s.stores.Rights.Update)
			if wsErr != nil {
				return nil, nil, wsErr
			}
			return b, func() ([]byte, *WSError) { return broadcastByID(ctx, s, repo.TableRights, id, s.stores.Rights.GetByID) }, nil
		case repo.TableRole:
			b, id, wsErr := processUpdate(ctx, s, mutation, s.stores.Role.Update)
			if wsErr != nil {
				return nil, nil, wsErr
			}
			return b, func() ([]byte, *WSError) { return broadcastByID(ctx, s, repo.TableRole, id, s.stores.Role.GetByID) }, nil
		case repo.TableServiceLink:
			b, _, wsErr := processUpdateNoReturn(ctx, s, mutation, 0, s.stores.ServiceLink.Update)
			if wsErr != nil {
				return nil, nil, wsErr
			}
			return b, func() ([]byte, *WSError) {
				return broadcastList(ctx, s, repo.TableServiceLink, s.stores.ServiceLink.Get)
			}, nil
		case repo.TableServiceSewobe:
			b, id, wsErr := processUpdate(ctx, s, mutation, s.stores.ServiceSewobe.Update)
			if wsErr != nil {
				return nil, nil, wsErr
			}
			return b, func() ([]byte, *WSError) {
				return broadcastByID(ctx, s, repo.TableServiceSewobe, id, s.stores.ServiceSewobe.GetByID)
			}, nil
		case repo.TableSettingsFrontend:
			b, _, wsErr := processUpdateNoReturn(ctx, s, mutation, 0, s.stores.SettingsFrontend.Update)
			if wsErr != nil {
				return nil, nil, wsErr
			}
			return b, func() ([]byte, *WSError) {
				return broadcastGetAllList(ctx, s, repo.TableSettingsFrontend, s.stores.SettingsFrontend.GetAll)
			}, nil
		case repo.TableSettingsPayment:
			b, _, wsErr := processUpdateNoReturn(ctx, s, mutation, 0, s.stores.SettingsPayment.Update)
			if wsErr != nil {
				return nil, nil, wsErr
			}
			return b, func() ([]byte, *WSError) {
				return broadcastGetAllList(ctx, s, repo.TableSettingsPayment, s.stores.SettingsPayment.GetAll)
			}, nil
		case repo.TableSettingsEmail:
			b, _, wsErr := processUpdateNoReturn(ctx, s, mutation, 0, s.stores.SettingsEmail.Update)
			if wsErr != nil {
				return nil, nil, wsErr
			}
			return b, func() ([]byte, *WSError) {
				return broadcastGetAllList(ctx, s, repo.TableSettingsEmail, s.stores.SettingsEmail.GetAll)
			}, nil
		case repo.TableUnit:
			b, id, wsErr := processUpdate(ctx, s, mutation, s.stores.Unit.Update)
			if wsErr != nil {
				return nil, nil, wsErr
			}
			return b, func() ([]byte, *WSError) { return broadcastByID(ctx, s, repo.TableUnit, id, s.stores.Unit.GetByID) }, nil
		case repo.TableUser:
			b, id, wsErr := processUpdate(ctx, s, mutation, s.stores.User.Update)
			if wsErr != nil {
				return nil, nil, wsErr
			}
			return b, func() ([]byte, *WSError) { return broadcastByID(ctx, s, repo.TableUser, id, s.stores.User.GetByID) }, nil
		case repo.TableVat:
			b, id, wsErr := processUpdate(ctx, s, mutation, s.stores.Vat.Update)
			if wsErr != nil {
				return nil, nil, wsErr
			}
			return b, func() ([]byte, *WSError) { return broadcastByID(ctx, s, repo.TableVat, id, s.stores.Vat.GetByID) }, nil
		}

	case repo.OpDelete:
		switch mutation.Table {
		case repo.TableAccount:
			return nil, nil, &WSError{Code: WSInvalidOperation, Message: "Invalid operation", Details: "Deletion of accounts is not supported"}
		case repo.TableCategory:
			b, _, wsErr := processDeleteByIntID(ctx, s, mutation, "category_id", s.stores.Category.Delete)
			if wsErr != nil {
				return nil, nil, wsErr
			}
			return b, func() ([]byte, *WSError) { return broadcastList(ctx, s, repo.TableCategory, s.stores.Category.Get) }, nil
		case repo.TableProductGroup:
			b, _, wsErr := processDeleteByIntID(ctx, s, mutation, "product_group_id", s.stores.ProductGroup.Delete)
			if wsErr != nil {
				return nil, nil, wsErr
			}
			return b, func() ([]byte, *WSError) {
				return broadcastList(ctx, s, repo.TableProductGroup, s.stores.ProductGroup.Get)
			}, nil
		case repo.TableProductVisibility:
			b, _, wsErr := processDeleteByIntID(ctx, s, mutation, "product_visibility_id", s.stores.ProductVisibility.Delete)
			if wsErr != nil {
				return nil, nil, wsErr
			}
			return b, func() ([]byte, *WSError) {
				return broadcastList(ctx, s, repo.TableProductVisibility, s.stores.ProductVisibility.Get)
			}, nil
		case repo.TableRights:
			b, _, wsErr := processDeleteByIntID(ctx, s, mutation, "rights_id", s.stores.Rights.Delete)
			if wsErr != nil {
				return nil, nil, wsErr
			}
			return b, func() ([]byte, *WSError) { return broadcastList(ctx, s, repo.TableRights, s.stores.Rights.Get) }, nil
		case repo.TableRole:
			b, _, wsErr := processDeleteByIntID(ctx, s, mutation, "role_id", s.stores.Role.Delete)
			if wsErr != nil {
				return nil, nil, wsErr
			}
			return b, func() ([]byte, *WSError) { return broadcastList(ctx, s, repo.TableRole, s.stores.Role.Get) }, nil
		case repo.TableServiceLink:
			idStr, ok := mutation.Where["foreign_user_id"]
			if !ok {
				return nil, nil, &WSError{Code: WSBadJSON, Message: "Missing foreign_user_id in where clause", Details: "Delete requires foreign_user_id"}
			}
			fid, parseErr := strconv.Atoi(idStr)
			if parseErr != nil {
				return nil, nil, &WSError{Code: WSBadJSON, Message: parseErr.Error(), Details: "foreign_user_id must be an integer"}
			}
			if delErr := s.stores.ServiceLink.Delete(ctx, fid); delErr != nil {
				return nil, nil, &WSError{Code: WSInternalError, Message: delErr.Error(), Details: "Could not delete from service_link"}
			}
			b, wsErr := s.marshalResultMutation(mutation.Table, mutation.Operation, fid)
			if wsErr != nil {
				return nil, nil, wsErr
			}
			return b, func() ([]byte, *WSError) {
				return broadcastList(ctx, s, repo.TableServiceLink, s.stores.ServiceLink.Get)
			}, nil
		case repo.TableServiceSewobe:
			b, _, wsErr := processDeleteByIntID(ctx, s, mutation, "service_sewobe_id", s.stores.ServiceSewobe.Delete)
			if wsErr != nil {
				return nil, nil, wsErr
			}
			return b, func() ([]byte, *WSError) {
				return broadcastList(ctx, s, repo.TableServiceSewobe, s.stores.ServiceSewobe.Get)
			}, nil
		case repo.TableSettingsFrontend:
			b, _, wsErr := processDeleteByStringKey(ctx, s, mutation, "key", s.stores.SettingsFrontend.Delete)
			if wsErr != nil {
				return nil, nil, wsErr
			}
			return b, func() ([]byte, *WSError) {
				return broadcastGetAllList(ctx, s, repo.TableSettingsFrontend, s.stores.SettingsFrontend.GetAll)
			}, nil
		case repo.TableSettingsPayment:
			b, _, wsErr := processDeleteByStringKey(ctx, s, mutation, "key", s.stores.SettingsPayment.Delete)
			if wsErr != nil {
				return nil, nil, wsErr
			}
			return b, func() ([]byte, *WSError) {
				return broadcastGetAllList(ctx, s, repo.TableSettingsPayment, s.stores.SettingsPayment.GetAll)
			}, nil
		case repo.TableSettingsEmail:
			b, _, wsErr := processDeleteByStringKey(ctx, s, mutation, "key", s.stores.SettingsEmail.Delete)
			if wsErr != nil {
				return nil, nil, wsErr
			}
			return b, func() ([]byte, *WSError) {
				return broadcastGetAllList(ctx, s, repo.TableSettingsEmail, s.stores.SettingsEmail.GetAll)
			}, nil
		case repo.TableUnit:
			b, _, wsErr := processDeleteByIntID(ctx, s, mutation, "unit_id", s.stores.Unit.Delete)
			if wsErr != nil {
				return nil, nil, wsErr
			}
			return b, func() ([]byte, *WSError) { return broadcastList(ctx, s, repo.TableUnit, s.stores.Unit.Get) }, nil
		case repo.TableUser:
			b, _, wsErr := processDeleteByIntID(ctx, s, mutation, "user_id", s.stores.User.Delete)
			if wsErr != nil {
				return nil, nil, wsErr
			}
			return b, func() ([]byte, *WSError) { return broadcastList(ctx, s, repo.TableUser, s.stores.User.Get) }, nil
		case repo.TableVat:
			b, _, wsErr := processDeleteByIntID(ctx, s, mutation, "vat_id", s.stores.Vat.Delete)
			if wsErr != nil {
				return nil, nil, wsErr
			}
			return b, func() ([]byte, *WSError) { return broadcastList(ctx, s, repo.TableVat, s.stores.Vat.Get) }, nil
		}
	}

	return nil, nil, &WSError{
		Code:    WSInvalidOperation,
		Message: "Invalid operation",
		Details: "Could not process mutation for operation " + string(mutation.Operation) + " on table " + string(mutation.Table),
	}
}
