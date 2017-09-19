package nibbles

// Location of all snakes on the board
type Board struct {
	width, height int
	snakes        []*Snake
}

// Directions enum
type Direction uint8

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
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
