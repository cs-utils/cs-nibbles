package server

import (
	"github.com/cs-utils/cs-nibbles/nibbles"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

/* Server responsible for handling all http and websocket requests */

type Server struct {
	address string
	hub     *Hub
	nibbles *nibbles.NibbleGame

	// Duration between game logic updates
	tickrate time.Duration

	// Method to stop game update loop
	stopGameloopChan chan bool
}

/*
Create and initialize the server.
Parameters:
	address: Address and port of the server. Same as http.ListenAndServe's address
*/
func CreateServer(address string, boardWidth, boardHeight int, tickrate time.Duration) *Server {
	serve := Server{
		address:  address,
		nibbles:  nibbles.CreateGame(boardWidth, boardHeight),
		tickrate: tickrate,
	}

	serve.hub = newHub()

	return &serve
}

func (s *Server) Start() error {
	// Start hub
	defer s.hub.stop()
	go s.hub.run()

	// Start main game loop
	defer s.stopGameLoop()
	go s.startGameLoop()

	// Mux router
	r := mux.NewRouter()
	r.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWebsocket(s.hub, w, r)
	})
	r.HandleFunc("/", serveIndexPage)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	http.Handle("/", r)
	return http.ListenAndServe(s.address, nil)
}
