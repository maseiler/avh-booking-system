package ws

import "time"

// ============================================================================
// Message Types and Protocol
// ============================================================================

type MessageType string

const (
	MsgTypeAuth            MessageType = "auth"
	MsgTypeRegister        MessageType = "register"
	MsgTypeUnRegister      MessageType = "unregister"
	MsgTypeError           MessageType = "error"
	MsgTypeBroadcast       MessageType = "broadcast"
	MsgTypePing            MessageType = "ping"
	MsgTypePong            MessageType = "pong"
	MsgTypeQuery           MessageType = "query"
	MsgTypeQueryResult     MessageType = "queryResult"
	MsgTypeQueryResultList MessageType = "queryResultList"
	MsgTypeMutation        MessageType = "mutation"
	MsgTypeMutationResult  MessageType = "mutationResult"
)

func (t MessageType) String() string {
	return string(t)
}

type Message struct {
	Type    MessageType `json:"type" validate:"required"`
	Payload interface{} `json:"payload,omitempty"`
}

type PingPong struct {
	Timestamp time.Time `json:"timestamp,omitempty"`
}
