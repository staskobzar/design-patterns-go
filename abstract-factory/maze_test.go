package maze

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeGenericRoom(t *testing.T) {
	r := &GenericRoom{}
	r.MakeRoom()
	assert.Equal(t, Wall, r.North())
	assert.Equal(t, Wall, r.East())
	assert.Equal(t, Wall, r.South())
	assert.Equal(t, Wall, r.West())

	r.MakeDoor(South)
	assert.Equal(t, Door, r.South())

	r.MakeWall(South)
	assert.Equal(t, Wall, r.South())
}

func TestMakeEnchantedRoom(t *testing.T) {
	r := &EnchantedRoom{}
	r.MakeRoom()
	assert.Equal(t, WallEnchanted, r.North())
	assert.Equal(t, WallEnchanted, r.East())
	assert.Equal(t, WallEnchanted, r.South())
	assert.Equal(t, WallEnchanted, r.West())

	r.MakeDoor(South)
	assert.Equal(t, DoorNeedingSpell, r.South())

	r.MakeWall(South)
	assert.Equal(t, WallEnchanted, r.South())
}

func TestMakeBombedRoom(t *testing.T) {
	r := &BombedRoom{}
	r.MakeRoom()
	assert.Equal(t, WallBobmed, r.North())
	assert.Equal(t, WallBobmed, r.East())
	assert.Equal(t, WallBobmed, r.South())
	assert.Equal(t, WallBobmed, r.West())

	r.MakeDoor(South)
	assert.Equal(t, DoorWithBomb, r.South())

	r.MakeWall(South)
	assert.Equal(t, WallBobmed, r.South())
}

func TestMakeMaze(t *testing.T) {

	maze := MakeMaze()
	maze.AddRoom(&GenericRoom{}, North)
	maze.AddRoom(&EnchantedRoom{}, West)
	maze.AddRoom(&BombedRoom{}, South)

	assert.Equal(t, 3, len(maze.rooms))

	assert.Equal(t, Door, maze.rooms[0].North())
	assert.Equal(t, Wall, maze.rooms[0].East())
	assert.Equal(t, Wall, maze.rooms[0].South())
	assert.Equal(t, Wall, maze.rooms[0].West())

	assert.Equal(t, WallEnchanted, maze.rooms[1].North())
	assert.Equal(t, WallEnchanted, maze.rooms[1].East())
	assert.Equal(t, WallEnchanted, maze.rooms[1].South())
	assert.Equal(t, DoorNeedingSpell, maze.rooms[1].West())

	assert.Equal(t, WallBobmed, maze.rooms[2].North())
	assert.Equal(t, WallBobmed, maze.rooms[2].East())
	assert.Equal(t, DoorWithBomb, maze.rooms[2].South())
	assert.Equal(t, WallBobmed, maze.rooms[2].West())
}
