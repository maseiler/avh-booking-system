package models

import (
	"encoding/json"
	"fmt"
)

type WsErrorCode uint

const (
	WsUnknown WsErrorCode = iota + 1
	WsBadJson
	WsBadInterface
	WsBadStruct
	WsInvalidFilter
	WsNotFound
	WsInternalError
)

const (
	WsUnknownStr       = "UNKNOWN"
	WsBadJsonStr       = "BAD_JSON"
	WsBadInterfaceStr  = "BAD_INTERFACE"
	WsBadStructStr     = "BAD_STRUCT"
	WsInvalidFilterStr = "INVALID_FILTER"
	WsNotFoundStr      = "NOT_FOUND"
	WsInternalErrorStr = "INTERNAL_ERROR"
)

func (c WsErrorCode) String() string {
	switch c {
	case WsBadJson:
		return WsBadJsonStr
	case WsBadInterface:
		return WsBadInterfaceStr
	case WsBadStruct:
		return WsBadStructStr
	case WsInvalidFilter:
		return WsInvalidFilterStr
	case WsNotFound:
		return WsNotFoundStr
	case WsInternalError:
		return WsInternalErrorStr
	default:
		return WsUnknownStr
	}
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (c *WsErrorCode) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	switch s {
	case WsBadJsonStr:
		*c = WsBadJson
	case WsBadInterfaceStr:
		*c = WsBadInterface
	case WsBadStructStr:
		*c = WsBadStruct
	case WsInvalidFilterStr:
		*c = WsInvalidFilter
	case WsNotFoundStr:
		*c = WsNotFound
	case WsInternalErrorStr:
		*c = WsInternalError
	default:
		return fmt.Errorf("invalid error: %s", s)
	}

	return nil
}

// MarshalJSON implements the json.Marshaler interface
func (c WsErrorCode) MarshalJSON() ([]byte, error) {
	var s string
	switch c {
	case WsBadJson:
		s = WsBadJsonStr
	case WsBadInterface:
		s = WsBadInterfaceStr
	case WsBadStruct:
		s = WsBadStructStr
	case WsInvalidFilter:
		s = WsInvalidFilterStr
	case WsNotFound:
		s = WsNotFoundStr
	case WsInternalError:
		s = WsInternalErrorStr
	default:
		return nil, fmt.Errorf("invalid error value: %d", c)
	}
	return json.Marshal(s)
}

type WsError struct {
	Code    WsErrorCode `json:"code" validate:"required"`
	Message string      `json:"message" validate:"required"`
	Details string      `json:"details,omitempty"`
}
