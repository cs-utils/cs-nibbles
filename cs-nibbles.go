package main

import (
	"github.com/cs-utils/cs-nibbles/server"
	log "github.com/sirupsen/logrus"
	"time"
)

func main() {
	log.SetLevel(log.DebugLevel)

	ADDRESS := ":8080"
	BOARD_WIDTH := 50
	BOARD_HEIGHT := 50
	TICKRATE := (1 * time.Second) / 2

	log.Info("Starting server on " + ADDRESS)
	server := server.CreateServer(ADDRESS, BOARD_WIDTH, BOARD_HEIGHT, TICKRATE)
	server.Start()
}
