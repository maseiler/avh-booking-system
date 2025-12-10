package ws

import (
	"encoding/json"
	"github.com/av-huette/avh-booking-system/internal/models"
	"log"
)

func (s *Service) sendError(c *models.Client, wsErr *models.WsError) {
	log.Printf("%v", wsErr)

	// Create error message
	msg := models.Message{
		Type:    models.MsgTypeError,
		Payload: wsErr,
	}

	// Marshal error message
	errBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	// Validate error message
	err = s.validator.ValidateMessage(errBytes)
	if err != nil {
		panic(err)
	}

	c.Send <- errBytes
}
