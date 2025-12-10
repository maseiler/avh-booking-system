package models

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
