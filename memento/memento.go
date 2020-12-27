package memento

type Point struct {
	top  int
	left int
}

type Graphic struct {
	pos *Point
}

func NewGraphic(top, left int) *Graphic {
	return &Graphic{pos: &Point{top, left}}
}

type Memento struct {
	state *Point
}

func NewMemento(p *Point) *Memento {
	return &Memento{state: p}
}

func (m *Memento) GetState() *Point {
	return m.state
}

type MoveCommand struct {
	state  *Memento
	target *Graphic
	dest   *Point
}

func NewMoveCommand(g *Graphic) *MoveCommand {
	c := &MoveCommand{target: g}
	return c
}

func (c *MoveCommand) Move(top, left int) {
	c.dest = &Point{top, left}
}

func (c *MoveCommand) Execute() {
	c.state = NewMemento(c.target.pos)
	c.target.pos = c.dest
}

func (c *MoveCommand) Unexecute() {
	c.target.pos = c.state.GetState()
}
