package server

import log "github.com/sirupsen/logrus"

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

	// Stop the server
	shutdown 	chan bool
}

func newHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		deregister: make(chan *Client),
	}
}

func (h *Hub) stop() {
	log.Info("Sending request to stop hub...")
	h.shutdown <- true
}


func (h *Hub) run() {
	for {
		select {

		// Stop hub's run loop
		case shouldShutdown := <-h.shutdown:
			if shouldShutdown {
				log.Info("Stopping hub")
				return
			}

		// New client, add to client list
		case client := <-h.register:
			h.clients[client] = true
			log.Info("New client added to hub")
		}
	}
}
