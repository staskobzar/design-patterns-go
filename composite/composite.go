package equipment

import (
	"fmt"
	"math"
)

type Price float32
type Watt uint16

type Equipment interface {
	Add(e Equipment)
	NetPrice() Price
	Discount() Price
	Price() string
	Power() Watt
}

type Element struct {
	power    Watt
	price    Price
	discount Price
	name     string
}

func (e Element) Add(item Equipment) {}
func (e Element) NetPrice() Price    { return e.price }
func (e Element) Discount() Price    { return e.discount }
func (e Element) Price() string      { return fmt.Sprintf("$%.2f", e.price-e.discount) }
func (e Element) Power() Watt        { return e.power }

type Composite struct {
	Element
	list []Equipment
}

func NewComposite(name string, power Watt, price, discount Price) *Composite {
	return &Composite{
		Element{power, price, discount, name},
		make([]Equipment, 0),
	}
}

func (c *Composite) Add(item Equipment) {
	c.list = append(c.list, item)
}

func (c *Composite) listEach(f func(e Equipment)) {
	for i := 0; i < len(c.list); i++ {
		f(c.list[i])
	}
}

func (c *Composite) NetPrice() Price {
	total := c.price
	c.listEach(func(e Equipment) {
		total += e.NetPrice()
	})
	return Price(math.Round(float64(total*100)) / 100)
}

func (c *Composite) Discount() Price {
	total := c.discount
	c.listEach(func(e Equipment) {
		total += e.Discount()
	})
	return Price(math.Round(float64(total*100)) / 100)
}

func (c *Composite) Power() Watt {
	total := c.power
	c.listEach(func(e Equipment) {
		total += e.Power()
	})
	return total
}

func (c *Composite) Price() string {
	return fmt.Sprintf("$%.2f", c.NetPrice()-c.Discount())
}

func NewCabinet(name string) Equipment {
	return NewComposite(name, 10, 9.75, 1.50)
}

func NewChassis(name string) Equipment {
	return NewComposite(name, 10, 15.25, 0.0)
}

func NewBus(name string) Equipment {
	return NewComposite(name, 12, 32.40, 4.1)
}

func NewCard(name string) Equipment {
	return Element{14, 26.83, 3.11, name}
}

func NewFloppyDisk(name string) Equipment {
	return Element{10, 2.68, 0.43, name}
}
