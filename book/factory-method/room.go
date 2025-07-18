package maze

type MazeElement int

const (
	Wall MazeElement = 1 + iota
	Door
	WallBombed
	DoorBombed
)

type Side int

const (
	North Side = 100 + iota
	East
	South
	West
)

type IRoom interface {
	MakeRoom()
	MakeDoor(side Side)
	GetRoomSide(side Side) MazeElement
}

type Room struct {
	side     map[Side]MazeElement
	wallType MazeElement
	doorType MazeElement
}

func (r *Room) MakeRoom() {
	r.side = make(map[Side]MazeElement, 4)
	for _, s := range []Side{North, East, South, West} {
		r.side[s] = r.wallType
	}
}

func (r *Room) MakeDoor(side Side) {
	r.side[side] = r.doorType
}

func (r *Room) GetRoomSide(side Side) MazeElement {
	return r.side[side]
}

type CommonRoom struct {
	*Room
}

type BombedRoom struct {
	*Room
}

func NewCommonRoom() IRoom {
	return &CommonRoom{
		Room: &Room{
			wallType: Wall,
			doorType: Door,
		},
	}
}

func NewBombedRoom() IRoom {
	return &BombedRoom{
		Room: &Room{
			wallType: WallBombed,
			doorType: DoorBombed,
		},
	}
}
