package shape

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuilder(t *testing.T) {
	circle := NewBuilder(4, 5, &Circle{})
	assert.Equal(t, "Draw Circle with center (4,5)", circle.draw())

	triangle := NewBuilder(8, -2, &Triangle{})
	assert.Equal(t, "Draw Triangle with top (8,-2)", triangle.draw())
}
