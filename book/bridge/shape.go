package shape

import "fmt"

type Shape interface {
	SetCoord(x, y int)
	Draw() string
}

type Triangle struct {
	x, y int
}

func (t *Triangle) SetCoord(x, y int) {
	t.x = x
	t.y = y
}

func (t *Triangle) Draw() string {
	return fmt.Sprintf("Draw Triangle with top (%d,%d)", t.x, t.y)
}

type Circle struct {
	x, y int
}

func (c *Circle) SetCoord(x, y int) {
	c.x = x
	c.y = y
}

func (c *Circle) Draw() string {
	return fmt.Sprintf("Draw Circle with center (%d,%d)", c.x, c.y)
}

type ShapeBuilder struct {
	shape Shape
	x, y  int
}

func NewBuilder(x, y int, shape Shape) *ShapeBuilder {
	builder := &ShapeBuilder{}
	builder.x = x
	builder.y = y
	builder.shape = shape
	return builder
}

func (s *ShapeBuilder) draw() string {
	s.shape.SetCoord(s.x, s.y)
	return s.shape.Draw()
}
