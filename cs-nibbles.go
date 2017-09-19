package main

import (
	"github.com/cs-utils/cs-nibbles/server"
	log "github.com/sirupsen/logrus"
	"time"
)

func main() {
	log.SetLevel(log.DebugLevel)

	ADDRESS := ":8080"
	BOARD_WIDTH := 100
	BOARD_HEIGHT := 100
	TICKRATE := (1 * time.Second) / 20

	log.Info("Starting server on " + ADDRESS)
	server := server.CreateServer(ADDRESS, BOARD_WIDTH, BOARD_HEIGHT, TICKRATE)
	server.Start()
}
