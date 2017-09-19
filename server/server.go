package server

import (
	"net/http"
	"github.com/gorilla/mux"
)

/* Server responsible for handling all http and websocket requests */

type Server struct {
	address string

	hub *Hub
}

/*
Create and initialize the server.
Parameters:
	address: Address and port of the server. Same as http.ListenAndServe's address
*/
func CreateServer(address string) *Server {
	serve := Server{
		address: address,
	}

	serve.hub = newHub()

	return &serve
}

func (s *Server) Start() error {
	// Start hub
	defer s.hub.stop()
	go s.hub.run()

	// Mux router
	r := mux.NewRouter()
	r.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request){
		serveWebsocket(s.hub, w, r)
	})
	r.HandleFunc("/", serveIndexPage)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	http.Handle("/", r)
	return http.ListenAndServe(s.address, nil)
}
