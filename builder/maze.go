package maze

import (
	"math/rand"
	"time"
)

type MazeBuilder interface {
	BuildMaze()
	BuildRoom(room *int)
	BuildDoor(roomFrom, roomTo int)
}

type Maze struct {
	room []*Room
}

func (m *Maze) BuildMaze() {
	m.room = make([]*Room, 0)
}

func (m *Maze) BuildRoom(roomId *int) {
	room := &Room{}
	room.InitRoom()
	*roomId = len(m.room)
	m.room = append(m.room, room)
}

func (m *Maze) BuildDoor(roomFrom, roomTo int) {
	r1 := m.room[roomFrom]
	r2 := m.room[roomTo]
	r1.SetElement(Door, m.CommonWall(r1, r2))
}

func (m *Maze) CommonWall(r1, r2 *Room) Side {
	rand.Seed(time.Now().UnixNano())
	sides := []Side{North, East, South, West}
	return sides[rand.Intn(len(sides))]
}

type CountingMaze struct {
	doors int
	rooms int
}

func (m *CountingMaze) BuildMaze() {
	m.doors = 0
	m.rooms = 0
}

func (m *CountingMaze) BuildRoom(roomId *int) {
	m.rooms++
}

func (m *CountingMaze) BuildDoor(roomFrom, roomTo int) {
	m.doors++
}

func (m *CountingMaze) GetMaze() (int, int) {
	return m.rooms, m.doors
}

// ------- directors
func CreateMaze(m MazeBuilder) {
	m.BuildMaze()

	var id1, id2 int

	m.BuildRoom(&id1)
	m.BuildRoom(&id2)
	m.BuildDoor(id1, id2)

	m.BuildRoom(&id1)
	m.BuildDoor(id2, id1)
}

func CreateCompexMaze(m MazeBuilder) {
	m.BuildMaze()

	for i := 0; i < 500; i++ {
		var id1, id2 int

		m.BuildRoom(&id1)
		m.BuildRoom(&id2)
		m.BuildDoor(id1, id2)
	}
}
