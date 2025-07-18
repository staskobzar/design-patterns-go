package visitor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPrice(t *testing.T) {
	p := NewPrice(13.56)
	assert.Equal(t, 1356, int(p))
}

func TestPriceDollarsCents(t *testing.T) {
	p := NewPrice(3.56)

	assert.Equal(t, 3, p.Dollars())
	assert.Equal(t, 56, p.Cents())
}

func TestPriceString(t *testing.T) {
	p := NewPrice(32.65)

	assert.Equal(t, "32.65$", p.String())
}

func TestPriceAdd(t *testing.T) {
	p := NewPrice(2.5)

	p = p.Add(NewPrice(4.01))

	assert.Equal(t, "6.51$", p.String())
}

func TestPriceSub(t *testing.T) {
	p := NewPrice(100.23)
	sub := NewPrice(4.1)

	p = p.Sub(sub)

	assert.Equal(t, "96.13$", p.String())
}
