package ws

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
)

// logLevel returns the appropriate log level for this error.
// Client errors (bad input) are logged at Warn; server errors at Error.
func (wse *WSError) logLevel() slog.Level {
	switch wse.Code {
	case WSBadJSON, WSBadInterface, WSBadStruct,
		WSInvalidTable, WSInvalidOperation, WSInvalidFilter,
		WSNotFound:
		return slog.LevelWarn
	default:
		return slog.LevelError
	}
}

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
	s.log.Log(context.Background(), wsErr.logLevel(), wsErr.String())

	// Create error message
	msg := Message{
		Type:    MsgTypeError,
		Payload: wsErr,
	}

	errBytes, err := json.Marshal(msg)
	if err != nil {
		s.log.Error("failed to marshal error message", slog.String("error", err.Error()))
		return
	}

	err = s.validator.validate(errBytes)
	if err != nil {
		s.log.Error("error message failed schema validation", slog.String("error", err.Error()))
		return
	}

	c.Send <- errBytes
}
