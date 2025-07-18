package interpreter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContext(t *testing.T) {
	c := NewContext()
	c.Assign("A", true)
	c.Assign("B", false)

	assert.True(t, c.Lookup("A"))
	assert.False(t, c.Lookup("B"))
	assert.False(t, c.Lookup("FOO"))
}

func TestVariable(t *testing.T) {
	c := NewContext()
	x := NewVariable("X", c)
	y := NewVariable("Y", c)

	c.Assign("X", true)
	c.Assign("Y", true)
	assert.True(t, x.Evaluate())
	assert.True(t, y.Evaluate())
}

func TestConstant(t *testing.T) {
	c := NewConstant(true)
	assert.True(t, c.Evaluate())
}

func TestAndExpr(t *testing.T) {
	c := NewContext()
	a := NewVariable("A", c)
	e := NewAndExp(NewConstant(true), a)

	c.Assign("A", true)

	assert.True(t, e.Evaluate())
	c.Assign("A", false)
	assert.False(t, e.Evaluate())
}

func TestOrExpr(t *testing.T) {
	c := NewContext()
	a := NewVariable("A", c)
	e := NewOrExp(NewConstant(false), a)

	c.Assign("A", true)

	assert.True(t, e.Evaluate())
	c.Assign("A", false)
	assert.False(t, e.Evaluate())
}

func TestNotExp(t *testing.T) {
	c := NewContext()
	a := NewVariable("A", c)

	e := NewNotExp(a)

	c.Assign("A", true)
	assert.False(t, e.Evaluate())

	e1 := NewNotExp(NewConstant(false))
	assert.True(t, e1.Evaluate())
}

func TestEvaluateAll(t *testing.T) {
	// (true and x) or (y and (not x))
	// with x=false and y=true result is true
	c := NewContext()
	x := NewVariable("x", c)
	y := NewVariable("y", c)
	e := NewOrExp(
		NewAndExp(NewConstant(true), x),
		NewAndExp(y, NewNotExp(x)),
	)

	c.Assign("x", false)
	c.Assign("y", true)

	assert.True(t, e.Evaluate())
}
