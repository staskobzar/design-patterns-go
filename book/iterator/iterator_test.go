package iterator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func ListStub() *List {
	l := NewList()
	l.Append(Employee{"Curtis Kiehn", "Joellenton", 25000})
	l.Append(Employee{"Burl Morar", "Anastasiaberg", 40000})
	l.Append(Employee{"Marge Marks", "Geraldfort", 55000})
	return l
}

func TestForwardIterator(t *testing.T) {
	list := ListStub()
	iter := NewForwardIterator(list)
	iter.First()
	assert.EqualValues(t, list.First(), iter.CurrentItem())
	iter.Next()
	assert.EqualValues(t, list.Get(1), iter.CurrentItem())
	assert.False(t, iter.IsDone())

	iter.Next()
	assert.EqualValues(t, list.Last(), iter.CurrentItem())
	assert.False(t, iter.IsDone())
	iter.Next()
	assert.True(t, iter.IsDone())
}

func TestBackwardIterator(t *testing.T) {
	list := ListStub()
	iter := NewBackwardIterator(list)
	iter.First()
	assert.EqualValues(t, list.Last(), iter.CurrentItem())
	iter.Next()
	assert.EqualValues(t, list.Get(1), iter.CurrentItem())
	assert.False(t, iter.IsDone())

	iter.Next()
	assert.EqualValues(t, list.First(), iter.CurrentItem())
	assert.False(t, iter.IsDone())
	iter.Next()
	assert.True(t, iter.IsDone())
}

func Example_PrintEmployees_Forward_iterator() {
	list := ListStub()
	iter := NewForwardIterator(list)

	PrintEmployees(iter)
	// Output:
	// Curtis Kiehn from Joellenton earns 25000$
	// Burl Morar from Anastasiaberg earns 40000$
	// Marge Marks from Geraldfort earns 55000$
}

func Example_PrintEmployees_Backward_iterator() {
	list := ListStub()
	iter := NewBackwardIterator(list)

	PrintEmployees(iter)
	// Output:
	// Marge Marks from Geraldfort earns 55000$
	// Burl Morar from Anastasiaberg earns 40000$
	// Curtis Kiehn from Joellenton earns 25000$
}
