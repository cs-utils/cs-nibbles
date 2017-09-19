package server

import "github.com/gorilla/websocket"

/* Bridge between websocket and hub */

type Client struct {
	hub *Hub

	conn *websocket.Conn
}
