package maze

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaze(t *testing.T) {
	maze := CreateMaze()

	assert.Equal(t, 2, len(maze.room))
	assert.Equal(t, Wall, maze.room[0].GetRoomSide(North))
	assert.Equal(t, WallBombed, maze.room[1].GetRoomSide(North))
}
