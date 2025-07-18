package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaze(t *testing.T) {
	m := CreateMaze()

	assert.Equal(t, 4, len(m.room))
}
