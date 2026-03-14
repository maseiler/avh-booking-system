package ws

import (
	"encoding/json"
	"fmt"
)

// --------------------------------------------------
// Websocket errors
// --------------------------------------------------

type WSErrorCode string

const (
	WSUnknown          = "UNKNOWN"
	WSBadJSON          = "BAD_JSON"
	WSBadInterface     = "BAD_INTERFACE"
	WSBadStruct        = "BAD_STRUCT"
	WSInvalidTable     = "INVALID_TABLE"
	WSInvalidOperation = "INVALID_OPERATION"
	WSInvalidFilter    = "INVALID_FILTER"
	WSNotFound         = "NOT_FOUND"
	WSInternalError    = "INTERNAL_ERROR"
	WSDBQueryError     = "DB_QUERY_ERROR"
)

type WSError struct {
	Code    WSErrorCode `json:"code" validate:"required"`
	Message string      `json:"message" validate:"required"`
	Details string      `json:"details,omitempty"`
}

func (wse *WSError) String() string {
	msg := fmt.Sprintf("code: %s, message: '%s'", wse.Code, wse.Message)
	if wse.Details != "" {
		msg += fmt.Sprintf(", details: '%s'", wse.Details)
	}
	return msg
}

func (s *Service) sendError(c *Client, wsErr *WSError) {
	s.log.Error(wsErr.String())

	// Create error message
	msg := Message{
		Type:    MsgTypeError,
		Payload: wsErr,
	}

	// Marshal error message
	errBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	// Validate error message
	err = s.validator.validate(errBytes)
	if err != nil {
		panic(err)
	}

	c.Send <- errBytes
}
