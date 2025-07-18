package memento

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoveCommand(t *testing.T) {
	graph := NewGraphic(100, 100)

	cmd := NewMoveCommand(graph)

	assert.Equal(t, 100, graph.pos.top)
	assert.Equal(t, 100, graph.pos.left)

	cmd.Move(200, 130)
	cmd.Execute()
	assert.Equal(t, 200, graph.pos.top)
	assert.Equal(t, 130, graph.pos.left)

	cmd.Unexecute()
	assert.Equal(t, 100, graph.pos.top)
	assert.Equal(t, 100, graph.pos.left)
}
