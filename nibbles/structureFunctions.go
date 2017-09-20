package nibbles

import (
	"strconv"
	"bytes"
)

/* Functions that can be executed on the structures */

/*
Serializes the board as a string. The format will be:
	id,x,y,x,y...;
That is, each snake is separated by a semicolon.
Each snake's data is separated by commas. The first element in
the array will be the snake's ID
*/
func (b *Board) SerializeBoardAsString() string {
	var buffer bytes.Buffer

	for _, snake := range b.snakes {
		buffer.WriteString(strconv.Itoa(int(snake.id)))

		for _, segment := range snake.segments {
			buffer.WriteString(",")
			buffer.WriteString(strconv.Itoa(int(segment.x)))
			buffer.WriteString(",")
			buffer.WriteString(strconv.Itoa(int(segment.y)))
		}

		buffer.WriteString(";")

	}

	return buffer.String()
}
