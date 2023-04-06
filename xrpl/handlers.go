package xrpl

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

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
	case "ledgerClosed":
		c.LedgerStream <- message
	case "validationReceived":
		c.ValidationStream <- message
	case "transaction":
		c.TransactionStream <- message
	case "peerStatusChange":
		c.PeerStatusStream <- message
	case "consensusPhase":
		c.ConsensusStream <- message
	case "path_find":
		c.PathFindStream <- message
	default:
	}
}
