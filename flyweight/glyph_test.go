package flyweight

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlyweight(t *testing.T) {
	g := CreateFactory()

	want := "[-d-100-10]h[-d-100-10]e[-d-100-10]" +
		"l[-d-100-10]l[-d-100-10]o"
	assert.Equal(t, want, g.Print("hello"))

	g.SetSize(12)
	g.SetWeight(300)
	g.SetFont("tt")
	want = "[-tt-300-12]h[-tt-300-12]e[-tt-300-12]" +
		"l[-tt-300-12]l[-tt-300-12]o"
	assert.Equal(t, want, g.Print("hello"))
}
