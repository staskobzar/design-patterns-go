package iterator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewList(t *testing.T) {
	l := NewList()
	assert.Equal(t, 0, l.Count())
}

func TestAppend(t *testing.T) {
	l := NewList()
	l.Append("Foo")
	assert.Equal(t, 1, l.Count())
	assert.EqualValues(t, "Foo", l.Get(0))
	assert.Nil(t, l.Get(100))
}

func TestPrepend(t *testing.T) {
	l := NewList()
	l.Append("alpha")
	l.Append("betta")
	l.Prepend("gamma")
	assert.Equal(t, 3, l.Count())
	assert.EqualValues(t, "gamma", l.First())
	assert.EqualValues(t, "alpha", l.Get(1))
	assert.EqualValues(t, "betta", l.Last())
}

func TestRemove(t *testing.T) {
	l := NewList()
	l.Append("alpha")
	l.Append(100)
	l.Append(200)
	l.Append("gamma")
	l.Append(5.101)
	assert.Equal(t, 5, l.Count())

	assert.False(t, l.Remove(100))
	assert.True(t, l.Remove(2))
	assert.Equal(t, 4, l.Count())

	assert.EqualValues(t, 5.101, l.Last())
	assert.True(t, l.RemoveLast())
	assert.Equal(t, 3, l.Count())
	assert.EqualValues(t, "gamma", l.Last())

	assert.EqualValues(t, "alpha", l.First())
	assert.True(t, l.RemoveFirst())
	assert.EqualValues(t, 100, l.First())

	assert.Equal(t, 2, l.Count())
	l.RemoveAll()
	assert.Equal(t, 0, l.Count())
}
