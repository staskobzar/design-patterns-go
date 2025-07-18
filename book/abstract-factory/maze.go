package maze

type MazeElement int

const (
	Wall MazeElement = 1 + iota
	WallEnchanted
	WallBobmed
	Door
	DoorNeedingSpell
	DoorWithBomb
)

type Side int

const (
	North Side = 100 + iota
	East
	South
	West
)

type Room struct {
	side map[Side]MazeElement
}

func (r *Room) InitRoom(wall MazeElement) {
	r.side = make(map[Side]MazeElement, 4)
	for _, s := range []Side{North, East, South, West} {
		r.SetElement(wall, s)
	}
}

func (r *Room) North() MazeElement {
	return r.side[North]
}

func (r *Room) East() MazeElement {
	return r.side[East]
}

func (r *Room) South() MazeElement {
	return r.side[South]
}

func (r *Room) West() MazeElement {
	return r.side[West]
}

func (r *Room) SetElement(el MazeElement, side Side) {
	r.side[side] = el
}

type MazeFactory interface {
	MakeRoom()
	MakeWall(side Side)
	MakeDoor(side Side)
	North() MazeElement
	East() MazeElement
	South() MazeElement
	West() MazeElement
}

type GenericRoom struct {
	*Room
}

func (r *GenericRoom) MakeRoom() {
	r.Room = &Room{}
	r.InitRoom(Wall)
}

func (r *GenericRoom) MakeDoor(side Side) {
	r.SetElement(Door, side)
}

func (r *GenericRoom) MakeWall(side Side) {
	r.SetElement(Wall, side)
}

type EnchantedRoom struct {
	*Room
}

func (r *EnchantedRoom) MakeRoom() {
	r.Room = &Room{}
	r.InitRoom(WallEnchanted)
}

func (r *EnchantedRoom) MakeDoor(side Side) {
	r.SetElement(DoorNeedingSpell, side)
}

func (r *EnchantedRoom) MakeWall(side Side) {
	r.SetElement(WallEnchanted, side)
}

type BombedRoom struct {
	*Room
}

func (r *BombedRoom) MakeRoom() {
	r.Room = &Room{}
	r.InitRoom(WallBobmed)
}

func (r *BombedRoom) MakeDoor(side Side) {
	r.SetElement(DoorWithBomb, side)
}

func (r *BombedRoom) MakeWall(side Side) {
	r.SetElement(WallBobmed, side)
}

type Maze struct {
	rooms []MazeFactory
}

func MakeMaze() *Maze {
	maze := &Maze{}
	maze.rooms = make([]MazeFactory, 0)
	return maze
}

func (m *Maze) AddRoom(factory MazeFactory, side Side) {
	factory.MakeRoom()
	factory.MakeDoor(side)
	m.rooms = append(m.rooms, factory)
}
