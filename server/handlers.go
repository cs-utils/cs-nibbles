package server

import "net/http"
import (
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

func serveIndexPage(w http.ResponseWriter, r *http.Request) {
	log.WithField("url", r.URL).Info("Serving index page")
	http.ServeFile(w, r, "static/index.html")
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Handle all incoming websocket requests
func serveWebsocket(h *Hub, w http.ResponseWriter, r *http.Request) {
	log.WithField("url", r.URL).Info("Got new websocket connection")

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.WithField("err", err).Error("Error upgrading websocket request")
		return
	}

	// Successfully accepted new websocket connection, associate it with a new client and add it to the hub
	client := &Client{
		hub:  h,
		conn: conn,
	}

	// Start handling incoming websocket messages on a separate go-routine
	go client.startIncomingMessageLoop()

	log.Debug("Sending client to hub")
	h.register <- client
}
