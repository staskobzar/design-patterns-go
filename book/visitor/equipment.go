package visitor

import (
	"github.com/staskobzar/design-patterns-go/iterator"
)

type Watt uint16

type Equipment interface {
	PrintName() string
	GPower() Watt
	NetPrice() Price
	DiscountPrice() Price
	Add(Equipment)
	Accept(Visitor)
}

type Element struct {
	Name     string
	Power    Watt
	Price    Price
	Discount Price
}

func NewElement(name string, power int, price float64, discount float64) Element {
	return Element{name, Watt(power), NewPrice(price), NewPrice(discount)}
}

func (e Element) PrintName() string    { return e.Name }
func (e Element) GPower() Watt         { return e.Power }
func (e Element) NetPrice() Price      { return e.Price }
func (e Element) DiscountPrice() Price { return e.Discount }
func (e Element) Add(Equipment)        {}
func (e Element) Accept(Visitor)       {}

type Composite struct {
	Element
	list *iterator.List
	iter iterator.Iterator
}

func NewComposite(name string, power int, price float64, discount float64) *Composite {
	el := NewElement(name, power, price, discount)
	list := iterator.NewList()
	iter := iterator.NewForwardIterator(list)
	return &Composite{el, list, iter}
}

func (c *Composite) Add(e Equipment) {
	c.list.Append(e)
}

type Cabinet struct{ *Composite }

func NewCabinet(name string, price float64) Equipment {
	return &Cabinet{NewComposite(name, 12, price, 0.5)}
}
func (c *Cabinet) Accept(v Visitor) {
	for c.iter.First(); !c.iter.IsDone(); c.iter.Next() {
		e := c.iter.CurrentItem()
		e.(Equipment).Accept(v)
	}
	v.VisitCabinet(c)
}

type MotherBoard struct{ *Composite }

func NewMotherBoard(name string, price float64) Equipment {
	return &MotherBoard{NewComposite(name, 18, price, 3.50)}
}
func (m *MotherBoard) Accept(v Visitor) {
	for m.iter.First(); !m.iter.IsDone(); m.iter.Next() {
		e := m.iter.CurrentItem()
		e.(Equipment).Accept(v)
	}
	v.VisitMotherBoard(m)
}

type CPU struct{ Element }

func NewCPU(name string, price float64) Equipment {
	return &CPU{NewElement(name, 22, price, 30.50)}
}
func (cpu *CPU) Accept(v Visitor) {
	v.VisitCPU(cpu)
}

type RAM struct{ Element }

func NewRAM(name string, price float64) Equipment {
	return &RAM{NewElement(name, 14, price, 2.35)}
}
func (ram *RAM) Accept(v Visitor) {
	v.VisitRAM(ram)
}

type HDD struct{ Element }

func NewHDD(name string, price float64) Equipment {
	return &HDD{NewElement(name, 24, price, 14.0)}
}
func (hdd *HDD) Accept(v Visitor) {
	v.VisitHDD(hdd)
}

type VideoCard struct{ Element }

func NewVideoCard(name string, price float64) Equipment {
	return &VideoCard{NewElement(name, 18, price, 4.85)}
}
func (card *VideoCard) Accept(v Visitor) {
	v.VisitVideoCard(card)
}
