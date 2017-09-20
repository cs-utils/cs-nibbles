package server

import "github.com/cs-utils/cs-nibbles/nibbles"

// All messages communicated between clients and the server will be WebsocketMessages
// serialized as JSON.
// 	Type: Purpose of message. Can be any of:
//		0:	Board Update. Sent from server to clients, represents current state of board.
const (
	MESSAGE_BOARD_UPDATE     int = 0
	MESSAGE_CHANGE_DIRECTION int = 1
)

type WebsocketMessage struct {
	Type int    `json:"type"`
	Data string `json:"data"`
}

// Internal message sent from client to server to update snake direction
type requestChangeSnakeDirection struct {
	c         *Client
	direction nibbles.Direction
}
