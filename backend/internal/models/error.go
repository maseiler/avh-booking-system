package models

import "fmt"

type WsErrorCode string

const (
	WsUnknown          = "UNKNOWN"
	WsBadJson          = "BAD_JSON"
	WsBadInterface     = "BAD_INTERFACE"
	WsBadStruct        = "BAD_STRUCT"
	WsInvalidTable     = "INVALID_TABLE"
	WsInvalidOperation = "INVALID_OPERATION"
	WsInvalidFilter    = "INVALID_FILTER"
	WsNotFound         = "NOT_FOUND"
	WsInternalError    = "INTERNAL_ERROR"
)

type WsError struct {
	Code    WsErrorCode `json:"code" validate:"required"`
	Message string      `json:"message" validate:"required"`
	Details string      `json:"details,omitempty"`
}

func (wse *WsError) String() string {
	msg := fmt.Sprintf("code: %s, message: '%s'", wse.Code, wse.Message)
	if wse.Details != "" {
		msg += fmt.Sprintf(", details: '%s'", wse.Details)
	}
	return msg
}
