package proxy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImageGraph(t *testing.T) {
	g := ImageGraph("dog.png")
	assert.Equal(t, "Draw dog.png. Type: png", g.Draw())
	assert.Equal(t, "dog.png", g.Name())
}
