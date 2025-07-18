package iterator

import "fmt"

type Employee struct {
	Name   string
	City   string
	Salary int
}

type Iterator interface {
	First()
	Next()
	IsDone() bool
	CurrentItem() interface{}
}

type ListIterator struct {
	list    *List
	current int
}

func (li *ListIterator) CurrentItem() interface{} {
	return li.list.Get(li.current)
}

func (li *ListIterator) IsDone() bool {
	return li.current < 0 || li.current >= li.list.Count()
}

type ForwardIterator struct {
	ListIterator
}

func NewForwardIterator(l *List) Iterator {
	return &ForwardIterator{
		ListIterator{l, 0},
	}
}

func (fi *ForwardIterator) First() {
	fi.current = 0
}

func (fi *ForwardIterator) Next() {
	fi.current++
}

type BackwardIterator struct {
	ListIterator
}

func NewBackwardIterator(l *List) Iterator {
	return &BackwardIterator{
		ListIterator{l, l.Count() - 1},
	}
}

func (bi *BackwardIterator) First() {
	bi.current = bi.list.Count() - 1
}

func (bi *BackwardIterator) Next() {
	bi.current--
}

// using iterators
func PrintEmployees(i Iterator) {
	for i.First(); !i.IsDone(); i.Next() {
		PrintEmplItem(i.CurrentItem().(Employee))
	}
}

func PrintEmplItem(empl Employee) {
	fmt.Printf("%s from %s earns %d$\n", empl.Name, empl.City, empl.Salary)
}
