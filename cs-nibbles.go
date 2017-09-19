package main

import (
	"github.com/cs-utils/cs-nibbles/server"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel)

	ADDRESS := ":8080"

	log.Info("Starting server on " + ADDRESS)
	server := server.CreateServer(ADDRESS)
	server.Start()
}
