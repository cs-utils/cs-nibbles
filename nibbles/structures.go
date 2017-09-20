package nibbles

// Location of all snakes on the board
type Board struct {
	width, height int
	snakes        []*Snake
}

// Directions enum
type Direction uint8

const (
	UP    Direction = 0
	DOWN  Direction = 1
	LEFT  Direction = 2
	RIGHT Direction = 3
)

type Snake struct {
	id        uint8
	direction Direction

	// Body of the snake. Element 0 is the head
	segments []SnakeSegment
}

type SnakeSegment struct {
	x, y int8
}
