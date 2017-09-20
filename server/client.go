package server

import "github.com/gorilla/websocket"
import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/cs-utils/cs-nibbles/nibbles"
)

/* Bridge between websocket and hub */

type Client struct {
	hub *Hub

	conn *websocket.Conn
}

func (c *Client) startIncomingMessageLoop() {
	log.Info("New client started incoming message loop")

	// Cleanup
	defer func() {
		c.hub.deregister <- c
		c.conn.Close()
	}()

	for {
		_, message, err := c.conn.ReadMessage()

		if err != nil {
			log.WithField("err", err).Info("Client received non-nil error. Stopping.")
			return
		}

		// Parse message
		msg := WebsocketMessage{}
		err = json.Unmarshal(message, &msg)
		if err != nil {
			log.WithField("err", err).Error("Error de-serializing message from client")
			continue
		}

		log.WithField("msg", msg).Debug("Got new message from client")

		// Handle message
		if msg.Type == MESSAGE_CHANGE_DIRECTION {
			switch msg.Data {
			case "0":
				c.hub.updateDirection <-requestChangeSnakeDirection{c, nibbles.UP}
			case "1":
				c.hub.updateDirection <-requestChangeSnakeDirection{c, nibbles.DOWN}
			case "2":
				c.hub.updateDirection <-requestChangeSnakeDirection{c, nibbles.LEFT}
			case "3":
				c.hub.updateDirection <-requestChangeSnakeDirection{c, nibbles.RIGHT}
			}
		}
	}
}
