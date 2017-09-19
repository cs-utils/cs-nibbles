package server

import (
	log "github.com/sirupsen/logrus"
	"time"
)

/* Manage the main game loop. */
func (s *Server) startGameLoop() {
	log.Info("Server starting main game loop")

	ticker := time.NewTicker(s.tickrate)

	for {
		select {

		// Update game logic and notify players of changes
		case <-ticker.C:
			log.Info("Got a tick")
			board := s.nibbles.Tick()
			s.hub.broadcastBoard <- board

		case shouldStop := <-s.stopGameloopChan:
			if shouldStop {
				log.Info("Stopping server game loop")
				return
			}
		}
	}
}

func (s *Server) stopGameLoop() {
	s.stopGameloopChan <- true
}
