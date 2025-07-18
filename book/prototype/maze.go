package main

type Side uint8

const (
	North Side = iota + 100
	East
	South
	West
)

type ISide interface {
	IsWall() bool
	IsDoor() bool
	Clone() ISide
}

type Room struct {
	side map[Side]ISide
}

func (r *Room) Clone() *Room {
	room := &Room{}
	room.side = make(map[Side]ISide)
	for _, s := range []Side{North, East, South, West} {
		room.side[s] = r.side[s].Clone()
	}
	return room
}

type Wall struct{}

func (w *Wall) IsWall() bool { return true }
func (w *Wall) IsDoor() bool { return false }
func (w *Wall) Clone() ISide { return &Wall{} }

type Door struct{}

func (d *Door) IsWall() bool { return false }
func (d *Door) IsDoor() bool { return true }
func (w *Door) Clone() ISide { return &Door{} }

type Maze struct {
	room []*Room
}

func CreateMaze() *Maze {
	m := &Maze{}

	r := &Room{}
	w := &Wall{}
	d := &Door{}

	r.side = make(map[Side]ISide)
	r.side[North] = w
	r.side[East] = w
	r.side[South] = d
	r.side[West] = w
	m.room = append(m.room, r)

	m.room = append(m.room, r.Clone())
	m.room = append(m.room, r.Clone())
	m.room = append(m.room, r.Clone())

	return m
}
