package nibbles

import "strconv"

/* Functions that can be executed on the structures */

/*
Serializes the board as a string. The format will be:
	id,x,y,x,y...;
That is, each snake is separated by a semicolon.
Each snake's data is separated by commas. The first element in
the array will be the snake's ID
*/
func (b *Board) SerializeBoardAsString() string {
	result := ""

	for _, snake := range b.snakes {
		result += strconv.Itoa(int(snake.id))

		for _, segment := range snake.segments {
			result += "," + strconv.Itoa(int(segment.x)) + "," + strconv.Itoa(int(segment.y))
		}

		result += ";"
	}

	return result
}
