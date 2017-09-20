package nibbles

import log "github.com/sirupsen/logrus"

/* All nibble-related game logic */

type NibbleGame struct {
	board Board
}

/*
Initializes nibble board and game.
*/
func CreateGame(width, height int) *NibbleGame {
	log.WithFields(log.Fields{
		"width":  width,
		"height": height,
	}).Info("Creating new board")

	game := &NibbleGame{
		board: Board{
			width:  width,
			height: height,
		},
	}

	//TODO
	//TEMP Init board with some snakes for testing

	for i := 0; i < 1; i++ {
		snake := &Snake{
			id:        uint8(i),
			direction: UP,
			segments: []SnakeSegment{
				SnakeSegment{x: int8(50), y: 25},
			},
		}

		for j := 0; j < 50; j++ {
			snake.segments = append(snake.segments, SnakeSegment{x: 50, y: 50})
		}

		game.board.snakes = append(game.board.snakes, snake)
	}

	return game
}

// Simulate the game for a single tick
func (nibble *NibbleGame) Tick() *Board {
	for _, snake := range nibble.board.snakes {
		newHead := SnakeSegment{
			x: snake.segments[0].x,
			y: snake.segments[0].y,
		}

		// Move new head in correct direction
		switch snake.direction {
		case UP:
			newHead.y--
		case DOWN:
			newHead.y++
		case LEFT:
			newHead.x--
		case RIGHT:
			newHead.x++
		}

		// Does it wrap? If so, wrap
		if int(newHead.x) >= nibble.board.width {
			newHead.x = 0
		}

		if newHead.x < 0 {
			newHead.x = int8(nibble.board.width - 1)
		}

		if int(newHead.y) >= nibble.board.height {
			newHead.y = 0
		}

		if newHead.y < 0 {
			newHead.y = int8(nibble.board.height - 1)
		}

		// Append head to snake and remove last snake segment
		snake.segments = append([]SnakeSegment{newHead}, snake.segments[:len(snake.segments)-1]...)
	}

	return &nibble.board
}

// Change direction of snake on board
func (nibble *NibbleGame) ChangeSnakeDirection(snakeID int, direction Direction) {
	nibble.board.snakes[0].direction = direction
}
