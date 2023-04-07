package xrpl

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

func (c *Client) handlePong(message string) error {
	fmt.Println("PONG response:", message)
	return nil
}

func (c *Client) handleResponse() error {
	go func() {
		for {
			if c.closed {
				break
			}
			messageType, message, err := c.connection.ReadMessage()
			if err != nil && websocket.IsCloseError(err) {
				log.Println("XRPL read error: ", err)
			}

			switch messageType {
			case websocket.CloseMessage:
				return
			case websocket.TextMessage:
				c.resolveStream(message)
			case websocket.BinaryMessage:
			default:
			}
		}
	}()
	return nil
}

func (c *Client) resolveStream(message []byte) {
	var m map[string]interface{}
	if err := json.Unmarshal(message, &m); err != nil {
		fmt.Println("json.Unmarshal error: ", err)
	}

	switch m["type"] {
	case StreamResponseType(StreamTypeLedger):
		c.StreamLedger <- message

	case StreamResponseType(StreamTypeTransaction):
		c.StreamTransaction <- message

	case StreamResponseType(StreamTypeValidations):
		c.StreamValidation <- message

	case StreamResponseType(StreamTypeManifests):
		c.StreamManifest <- message

	case StreamResponseType(StreamTypePeerStatus):
		c.StreamPeerStatus <- message

	case StreamResponseType(StreamTypeConsensus):
		c.StreamConsensus <- message

	case StreamResponseType(StreamTypePathFind):
		c.StreamPathFind <- message

	case StreamResponseType(StreamTypeServer):
		c.StreamServer <- message

	default:
		c.StreamDefault <- message
	}
}
