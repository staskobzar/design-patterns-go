package maze

type Maze struct {
	room []IRoom
}

func CreateMaze() *Maze {
	maze := &Maze{}
	maze.room = make([]IRoom, 0)

	doorSides := []Side{East, West}
	for i, room := range []IRoom{maze.CreateRoom("common"), maze.CreateRoom("bombed")} {
		room.MakeRoom()
		room.MakeDoor(doorSides[i])
		maze.room = append(maze.room, room)
	}
	return maze
}

func (m *Maze) CreateRoom(room string) IRoom {
	switch room {
	case "bombed":
		return NewBombedRoom()
	default:
		return NewCommonRoom()
	}
}
