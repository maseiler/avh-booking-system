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

func (s *Service) processMutation(ctx context.Context, mutation *Mutation) ([]byte, int, *WSError) {
	switch mutation.Operation {
	case repo.OpInsert:
		switch mutation.Table {
		case repo.TableAccount:
			return processInsert(ctx, s, mutation, s.stores.Account.Insert)
		case repo.TableCategory:
			return processInsert(ctx, s, mutation, s.stores.Category.Insert)
		case repo.TableFavorites:
			return processInsertNoReturn(ctx, s, mutation, 0, s.stores.Favorites.Insert)
		case repo.TableLocation:
			return processInsert(ctx, s, mutation, s.stores.Location.Insert)
		case repo.TableOrder:
			return processInsert(ctx, s, mutation, s.stores.Order.Insert)
		case repo.TableProduct:
			return processInsert(ctx, s, mutation, s.stores.Product.Insert)
		case repo.TableProductGroup:
			return processInsert(ctx, s, mutation, s.stores.ProductGroup.Insert)
		case repo.TableProductOrder:
			return processInsertNoReturn(ctx, s, mutation, 0, s.stores.ProductOrder.Insert)
		case repo.TableProductVisibility:
			return processInsert(ctx, s, mutation, s.stores.ProductVisibility.Insert)
		case repo.TableRights:
			return processInsert(ctx, s, mutation, s.stores.Rights.Insert)
		case repo.TableRole:
			return processInsert(ctx, s, mutation, s.stores.Role.Insert)
		case repo.TableServiceLink:
			return processInsertNoReturn(ctx, s, mutation, 0, s.stores.ServiceLink.Insert)
		case repo.TableServiceSewobe:
			return processInsert(ctx, s, mutation, s.stores.ServiceSewobe.Insert)
		case repo.TableSettingsFrontend:
			return processInsertNoReturn(ctx, s, mutation, 0, s.stores.SettingsFrontend.Insert)
		case repo.TableSettingsPayment:
			return processInsertNoReturn(ctx, s, mutation, 0, s.stores.SettingsPayment.Insert)
		case repo.TableSettingsEmail:
			return processInsertNoReturn(ctx, s, mutation, 0, s.stores.SettingsEmail.Insert)
		case repo.TableUnit:
			return processInsert(ctx, s, mutation, s.stores.Unit.Insert)
		case repo.TableUser:
			return processInsert(ctx, s, mutation, s.stores.User.Insert)
		case repo.TableVat:
			return processInsert(ctx, s, mutation, s.stores.Vat.Insert)
		}

	case repo.OpUpdate:
		switch mutation.Table {
		case repo.TableAccount:
			return processUpdate(ctx, s, mutation, s.stores.Account.Update)
		case repo.TableCategory:
			return processUpdate(ctx, s, mutation, s.stores.Category.Update)
		case repo.TableFavorites:
			return processUpdateNoReturn(ctx, s, mutation, 0, s.stores.Favorites.Update)
		case repo.TableProduct:
			return processUpdate(ctx, s, mutation, s.stores.Product.Update)
		case repo.TableProductGroup:
			return processUpdate(ctx, s, mutation, s.stores.ProductGroup.Update)
		case repo.TableRights:
			return processUpdate(ctx, s, mutation, s.stores.Rights.Update)
		case repo.TableRole:
			return processUpdate(ctx, s, mutation, s.stores.Role.Update)
		case repo.TableServiceLink:
			return processUpdateNoReturn(ctx, s, mutation, 0, s.stores.ServiceLink.Update)
		case repo.TableServiceSewobe:
			return processUpdate(ctx, s, mutation, s.stores.ServiceSewobe.Update)
		case repo.TableSettingsFrontend:
			return processUpdateNoReturn(ctx, s, mutation, 0, s.stores.SettingsFrontend.Update)
		case repo.TableSettingsPayment:
			return processUpdateNoReturn(ctx, s, mutation, 0, s.stores.SettingsPayment.Update)
		case repo.TableSettingsEmail:
			return processUpdateNoReturn(ctx, s, mutation, 0, s.stores.SettingsEmail.Update)
		case repo.TableUnit:
			return processUpdate(ctx, s, mutation, s.stores.Unit.Update)
		case repo.TableUser:
			return processUpdate(ctx, s, mutation, s.stores.User.Update)
		case repo.TableVat:
			return processUpdate(ctx, s, mutation, s.stores.Vat.Update)
		}

	case repo.OpDelete:
		switch mutation.Table {
		case repo.TableAccount:
			return nil, 0, &WSError{Code: WSInvalidOperation, Message: "Invalid operation", Details: "Deletion of accounts is not supported"}
		case repo.TableCategory:
			return processDeleteByIntID(ctx, s, mutation, "category_id", s.stores.Category.Delete)
		case repo.TableProductGroup:
			return processDeleteByIntID(ctx, s, mutation, "product_group_id", s.stores.ProductGroup.Delete)
		case repo.TableProductVisibility:
			return processDeleteByIntID(ctx, s, mutation, "product_visibility_id", s.stores.ProductVisibility.Delete)
		case repo.TableRights:
			return processDeleteByIntID(ctx, s, mutation, "rights_id", s.stores.Rights.Delete)
		case repo.TableRole:
			return processDeleteByIntID(ctx, s, mutation, "role_id", s.stores.Role.Delete)
		case repo.TableServiceLink:
			idStr, ok := mutation.Where["foreign_user_id"]
			if !ok {
				return nil, 0, &WSError{Code: WSBadJSON, Message: "Missing foreign_user_id in where clause", Details: "Delete requires foreign_user_id"}
			}
			fid, parseErr := strconv.Atoi(idStr)
			if parseErr != nil {
				return nil, 0, &WSError{Code: WSBadJSON, Message: parseErr.Error(), Details: "foreign_user_id must be an integer"}
			}
			if delErr := s.stores.ServiceLink.Delete(ctx, fid); delErr != nil {
				return nil, 0, &WSError{Code: WSInternalError, Message: delErr.Error(), Details: "Could not delete from service_link"}
			}
			b, wsErr := s.marshalResultMutation(mutation.Table, mutation.Operation, fid)
			return b, fid, wsErr
		case repo.TableServiceSewobe:
			return processDeleteByIntID(ctx, s, mutation, "service_sewobe_id", s.stores.ServiceSewobe.Delete)
		case repo.TableSettingsFrontend:
			return processDeleteByStringKey(ctx, s, mutation, "key", s.stores.SettingsFrontend.Delete)
		case repo.TableSettingsPayment:
			return processDeleteByStringKey(ctx, s, mutation, "key", s.stores.SettingsPayment.Delete)
		case repo.TableSettingsEmail:
			return processDeleteByStringKey(ctx, s, mutation, "key", s.stores.SettingsEmail.Delete)
		case repo.TableUnit:
			return processDeleteByIntID(ctx, s, mutation, "unit_id", s.stores.Unit.Delete)
		case repo.TableUser:
			return processDeleteByIntID(ctx, s, mutation, "user_id", s.stores.User.Delete)
		case repo.TableVat:
			return processDeleteByIntID(ctx, s, mutation, "vat_id", s.stores.Vat.Delete)
		}
	}

	return nil, 0, &WSError{
		Code:    WSInvalidOperation,
		Message: "Invalid operation",
		Details: "Could not process mutation for operation " + string(mutation.Operation) + " on table " + string(mutation.Table),
	}
}

func (s *Service) prepareBroadcast(ctx context.Context, mutation *Mutation, id int) ([]byte, *WSError) {
	switch mutation.Table {
	case repo.TableAccount:
		return broadcastByID(ctx, s, repo.TableAccount, id, s.stores.Account.GetByID)

	case repo.TableCategory:
		if mutation.Operation == repo.OpDelete {
			return broadcastList(ctx, s, repo.TableCategory, s.stores.Category.Get)
		}
		return broadcastByID(ctx, s, repo.TableCategory, id, s.stores.Category.GetByID)

	case repo.TableFavorites:
		return broadcastList(ctx, s, repo.TableFavorites, s.stores.Favorites.Get)

	case repo.TableLocation:
		return broadcastByID(ctx, s, repo.TableLocation, id, s.stores.Location.GetByID)

	case repo.TableOrder:
		return broadcastByID(ctx, s, repo.TableOrder, id, s.stores.Order.GetByID)

	case repo.TableProduct:
		return broadcastByID(ctx, s, repo.TableProduct, id, s.stores.Product.GetByID)

	case repo.TableProductGroup:
		if mutation.Operation == repo.OpDelete {
			return broadcastList(ctx, s, repo.TableProductGroup, s.stores.ProductGroup.Get)
		}
		return broadcastByID(ctx, s, repo.TableProductGroup, id, s.stores.ProductGroup.GetByID)

	case repo.TableProductOrder:
		return broadcastList(ctx, s, repo.TableProductOrder, s.stores.ProductOrder.Get)

	case repo.TableProductVisibility:
		if mutation.Operation == repo.OpDelete {
			return broadcastList(ctx, s, repo.TableProductVisibility, s.stores.ProductVisibility.Get)
		}
		return broadcastByID(ctx, s, repo.TableProductVisibility, id, s.stores.ProductVisibility.GetByID)

	case repo.TableRights:
		if mutation.Operation == repo.OpDelete {
			return broadcastList(ctx, s, repo.TableRights, s.stores.Rights.Get)
		}
		return broadcastByID(ctx, s, repo.TableRights, id, s.stores.Rights.GetByID)

	case repo.TableRole:
		if mutation.Operation == repo.OpDelete {
			return broadcastList(ctx, s, repo.TableRole, s.stores.Role.Get)
		}
		return broadcastByID(ctx, s, repo.TableRole, id, s.stores.Role.GetByID)

	case repo.TableServiceLink:
		return broadcastList(ctx, s, repo.TableServiceLink, s.stores.ServiceLink.Get)

	case repo.TableServiceSewobe:
		if mutation.Operation == repo.OpDelete {
			return broadcastList(ctx, s, repo.TableServiceSewobe, s.stores.ServiceSewobe.Get)
		}
		return broadcastByID(ctx, s, repo.TableServiceSewobe, id, s.stores.ServiceSewobe.GetByID)

	case repo.TableSettingsFrontend:
		return broadcastGetAllList(ctx, s, repo.TableSettingsFrontend, s.stores.SettingsFrontend.GetAll)

	case repo.TableSettingsPayment:
		return broadcastGetAllList(ctx, s, repo.TableSettingsPayment, s.stores.SettingsPayment.GetAll)

	case repo.TableSettingsEmail:
		return broadcastGetAllList(ctx, s, repo.TableSettingsEmail, s.stores.SettingsEmail.GetAll)

	case repo.TableUnit:
		if mutation.Operation == repo.OpDelete {
			return broadcastList(ctx, s, repo.TableUnit, s.stores.Unit.Get)
		}
		return broadcastByID(ctx, s, repo.TableUnit, id, s.stores.Unit.GetByID)

	case repo.TableUser:
		if mutation.Operation == repo.OpDelete {
			return broadcastList(ctx, s, repo.TableUser, s.stores.User.Get)
		}
		return broadcastByID(ctx, s, repo.TableUser, id, s.stores.User.GetByID)

	case repo.TableVat:
		if mutation.Operation == repo.OpDelete {
			return broadcastList(ctx, s, repo.TableVat, s.stores.Vat.Get)
		}
		return broadcastByID(ctx, s, repo.TableVat, id, s.stores.Vat.GetByID)
	}

	return nil, nil
}
