package ws

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/av-huette/avh-booking-system/internal/repo"
)

func processGetQuery[T any](ctx context.Context, s *Service, query *repo.Query, get func(context.Context, *repo.Query) ([]T, error)) ([]byte, *WSError) {
	items, err := get(ctx, query)
	if err != nil {
		return nil, queryError(query.Table, err)
	}
	raw, wsErr := marshalSlice(items)
	if wsErr != nil {
		return nil, wsErr
	}
	return s.marshalQueryResultList(query.Table, &raw)
}

func processGetAllQuery[T any](ctx context.Context, s *Service, table repo.TableName, getAll func(context.Context) ([]T, error)) ([]byte, *WSError) {
	items, err := getAll(ctx)
	if err != nil {
		return nil, queryError(table, err)
	}
	raw, wsErr := marshalSlice(items)
	if wsErr != nil {
		return nil, wsErr
	}
	return s.marshalQueryResultList(table, &raw)
}

func processInsert[T any](ctx context.Context, s *Service, mutation *Mutation, insert func(context.Context, T) (int, error)) ([]byte, int, *WSError) {
	item, wsErr := unmarshalInterface[T](mutation.Values)
	if wsErr != nil {
		return nil, 0, wsErr
	}
	newID, err := insert(ctx, *item)
	if err != nil {
		return nil, 0, &WSError{Code: WSInternalError, Message: err.Error(), Details: "Could not insert into " + string(mutation.Table)}
	}
	b, wsErr := s.marshalResultMutation(mutation.Table, mutation.Operation, newID)
	return b, newID, wsErr
}

// processInsertNoReturn handles inserts for stores that return only error (composite-PK tables).
// id is provided by the caller (e.g. the parent entity's ID).
func processInsertNoReturn[T any](ctx context.Context, s *Service, mutation *Mutation, id int, insert func(context.Context, T) error) ([]byte, int, *WSError) {
	item, wsErr := unmarshalInterface[T](mutation.Values)
	if wsErr != nil {
		return nil, 0, wsErr
	}
	if err := insert(ctx, *item); err != nil {
		return nil, 0, &WSError{Code: WSInternalError, Message: err.Error(), Details: "Could not insert into " + string(mutation.Table)}
	}
	b, wsErr := s.marshalResultMutation(mutation.Table, mutation.Operation, id)
	return b, id, wsErr
}

func processUpdate[T any](ctx context.Context, s *Service, mutation *Mutation, update func(context.Context, T) (int, error)) ([]byte, int, *WSError) {
	item, wsErr := unmarshalInterface[T](mutation.Values)
	if wsErr != nil {
		return nil, 0, wsErr
	}
	id, err := update(ctx, *item)
	if err != nil {
		return nil, 0, &WSError{Code: WSInternalError, Message: err.Error(), Details: "Could not update in " + string(mutation.Table)}
	}
	b, wsErr := s.marshalResultMutation(mutation.Table, mutation.Operation, id)
	return b, id, wsErr
}

// processUpdateNoReturn handles updates for stores that return only error (composite-PK tables).
// id is provided by the caller.
func processUpdateNoReturn[T any](ctx context.Context, s *Service, mutation *Mutation, id int, update func(context.Context, T) error) ([]byte, int, *WSError) {
	item, wsErr := unmarshalInterface[T](mutation.Values)
	if wsErr != nil {
		return nil, 0, wsErr
	}
	if err := update(ctx, *item); err != nil {
		return nil, 0, &WSError{Code: WSInternalError, Message: err.Error(), Details: "Could not update in " + string(mutation.Table)}
	}
	b, wsErr := s.marshalResultMutation(mutation.Table, mutation.Operation, id)
	return b, id, wsErr
}

// processDeleteByIntID handles deletes where the PK is a single int.
// idField is the key in mutation.Where (e.g. "product_visibility_id").
func processDeleteByIntID(ctx context.Context, s *Service, mutation *Mutation, idField string, del func(context.Context, int) (int, error)) ([]byte, int, *WSError) {
	idStr, ok := mutation.Where[idField]
	if !ok {
		return nil, 0, &WSError{Code: WSBadJSON, Message: "Missing " + idField + " in where clause", Details: "Delete requires " + idField}
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, 0, &WSError{Code: WSBadJSON, Message: err.Error(), Details: idField + " must be an integer"}
	}
	oldID, err := del(ctx, id)
	if err != nil {
		return nil, 0, &WSError{Code: WSInternalError, Message: err.Error(), Details: "Could not delete from " + string(mutation.Table)}
	}
	b, wsErr := s.marshalResultMutation(mutation.Table, mutation.Operation, oldID)
	return b, oldID, wsErr
}

// processDeleteByStringKey handles deletes where the PK is a string (Settings tables).
func processDeleteByStringKey(ctx context.Context, s *Service, mutation *Mutation, keyField string, del func(context.Context, string) error) ([]byte, int, *WSError) {
	key, ok := mutation.Where[keyField]
	if !ok {
		return nil, 0, &WSError{Code: WSBadJSON, Message: "Missing " + keyField + " in where clause", Details: "Delete requires " + keyField}
	}
	if err := del(ctx, key); err != nil {
		return nil, 0, &WSError{Code: WSInternalError, Message: err.Error(), Details: "Could not delete from " + string(mutation.Table)}
	}
	b, wsErr := s.marshalResultMutation(mutation.Table, mutation.Operation, 0)
	return b, 0, wsErr
}

func broadcastByID[T any](ctx context.Context, s *Service, table repo.TableName, id int, getByID func(context.Context, int) (*T, error)) ([]byte, *WSError) {
	item, err := getByID(ctx, id)
	if err != nil {
		return nil, &WSError{Code: WSDBQueryError, Message: err.Error(), Details: fmt.Sprintf("Could not get %s with ID %d", string(table), id)}
	}
	raw, jsonErr := json.Marshal(item)
	if jsonErr != nil {
		return nil, &WSError{Code: WSInternalError, Message: jsonErr.Error(), Details: "Failed to encode " + string(table) + " for broadcast"}
	}
	return s.marshalBroadcastQueryResult(table, raw)
}

func broadcastList[T any](ctx context.Context, s *Service, table repo.TableName, get func(context.Context, *repo.Query) ([]T, error)) ([]byte, *WSError) {
	items, err := get(ctx, &repo.Query{Table: table})
	if err != nil {
		return nil, &WSError{Code: WSDBQueryError, Message: err.Error(), Details: "Could not reload " + string(table) + " after mutation"}
	}
	raw, wsErr := marshalSlice(items)
	if wsErr != nil {
		return nil, wsErr
	}
	return s.marshalBroadcastQueryResultList(table, &raw)
}

func broadcastGetAllList[T any](ctx context.Context, s *Service, table repo.TableName, getAll func(context.Context) ([]T, error)) ([]byte, *WSError) {
	items, err := getAll(ctx)
	if err != nil {
		return nil, &WSError{Code: WSDBQueryError, Message: err.Error(), Details: "Could not reload " + string(table) + " after mutation"}
	}
	raw, wsErr := marshalSlice(items)
	if wsErr != nil {
		return nil, wsErr
	}
	return s.marshalBroadcastQueryResultList(table, &raw)
}
