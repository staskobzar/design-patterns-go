package visitor

import (
	"fmt"
	"math"
)

type Price int64

func NewPrice(price float64) Price {
	return Price(math.Round(price * 100))
}

func (p Price) Dollars() int {
	return int(p) / 100
}

func (p Price) Cents() int {
	return int(p) % 100
}

func (p Price) String() string {
	return fmt.Sprintf("%d.%d$", p.Dollars(), p.Cents())
}

func (p Price) Add(add Price) Price {
	return p + add
}

func (p Price) Sub(sub Price) Price {
	return p - sub
}
