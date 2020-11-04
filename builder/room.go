package maze

type MazeElement int

const (
	Wall MazeElement = 1 + iota
	Door
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

func (r *Room) InitRoom() {
	r.side = make(map[Side]MazeElement, 4)
	for _, s := range []Side{North, East, South, West} {
		r.SetElement(Wall, s)
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
