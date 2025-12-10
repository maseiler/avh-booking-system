package ws

import (
	"encoding/json"
	"github.com/av-huette/avh-booking-system/internal/models"
	"log"
)

func sendError(c *models.Client, wsErr *models.WsError) {
	log.Printf("%v", wsErr)
	msg := models.Message{
		Type: models.MsgTypeError,
		Payload: map[string]interface{}{
			"error": wsErr},
	}
	errBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	c.Send <- errBytes
}
