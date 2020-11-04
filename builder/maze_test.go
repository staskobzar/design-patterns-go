package maze

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildRoom(t *testing.T) {
	m := &Maze{}
	var roomId int

	m.BuildRoom(&roomId)
	assert.Equal(t, 0, roomId)

	m.BuildRoom(&roomId)
	assert.Equal(t, 1, roomId)

	m.BuildRoom(&roomId)
	assert.Equal(t, 2, roomId)
}

func TestCreateMaze(t *testing.T) {
	maze := &Maze{}
	CreateMaze(maze)

	assert.Equal(t, 3, len(maze.room))
}

func TestCreateCountingMaze(t *testing.T) {
	maze := &CountingMaze{}
	CreateMaze(maze)
	rooms, doors := maze.GetMaze()
	assert.Equal(t, 3, rooms)
	assert.Equal(t, 2, doors)

	CreateCompexMaze(maze)
	rooms, doors = maze.GetMaze()
	assert.Equal(t, 1000, rooms)
	assert.Equal(t, 500, doors)
}

func TestCreateComplexMaze(t *testing.T) {
	maze := &Maze{}
	CreateCompexMaze(maze)

	assert.Equal(t, 1000, len(maze.room))
}
