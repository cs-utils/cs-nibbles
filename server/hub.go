package server

import (
	"encoding/json"
	"github.com/cs-utils/cs-nibbles/nibbles"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	//"time"
)

/* 	Container of all clients connected to the websocket server.
Inspired by gorilla's websocket example
*/

type Hub struct {
	// All connected clients
	clients map[*Client]bool

	// Messages from clients to server
	broadcast chan []byte

	// Requests from clients to register or deregister
	register   chan *Client
	deregister chan *Client

	// Broadcast board
	broadcastBoard chan *nibbles.Board

	// Client requests to update direction
	updateDirection chan requestChangeSnakeDirection

	// Stop the server
	shutdown chan bool
}

func newHub() *Hub {
	return &Hub{
		clients:         make(map[*Client]bool),
		broadcast:       make(chan []byte),
		register:        make(chan *Client),
		deregister:      make(chan *Client),
		shutdown:        make(chan bool),
		broadcastBoard:  make(chan *nibbles.Board),
		updateDirection: make(chan requestChangeSnakeDirection),
	}
}

func (h *Hub) stop() {
	log.Info("Sending request to stop hub...")
	h.shutdown <- true
}

func (h *Hub) run() {
	for {
		select {

		// Broadcast board to all players
		case board := <-h.broadcastBoard:
			//prev := time.Now()

			updateMessage := WebsocketMessage{
				Type: MESSAGE_BOARD_UPDATE,
				Data: board.SerializeBoardAsString(),
			}

			rawData, err := json.Marshal(&updateMessage)
			if err != nil {
				log.WithField("err", err).Error("Could not serialize board -> json to clients.")
				continue
			}

			for client, registered := range h.clients {
				if !registered {
					continue
				}

				w, err := client.conn.NextWriter(websocket.TextMessage)
				if err != nil {
					log.WithField("err", err).Error("Could not create writer for client")
					continue
				}

				w.Write(rawData)
				if err := w.Close(); err != nil {
					log.WithField("err", err).Error("Could not send board for client")
					continue
				}
			}

			//delta := time.Now().Sub(prev)
			//log.WithField("timeTaken", delta).Debug("Sent board to clients")

		// New client, add to client list
		case client := <-h.register:
			_, exists := h.clients[client]
			if exists {
				log.Warning("Client tried to register while in hub already?")
				continue
			}

			h.clients[client] = true
			log.Info("New client added to hub")

			// Stop hub's run loop
		case shouldShutdown := <-h.shutdown:
			if shouldShutdown {
				log.Info("Stopping hub")
				return
			}
		}
	}
}
