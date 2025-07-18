package iterator

// Design Patterns Appendix C.1
// List template
type List struct {
	item []interface{}
}

func NewList() *List {
	return &List{
		item: make([]interface{}, 0),
	}
}

func (l *List) indexOk(index int) bool {
	if index < 0 || index >= l.Count() {
		return false
	}
	return true
}

func (l *List) Count() int {
	return len(l.item)
}

func (l *List) Get(index int) interface{} {
	if ok := l.indexOk(index); !ok {
		return nil
	}
	return l.item[index]
}

func (l *List) First() interface{} {
	return l.item[0]
}

func (l *List) Last() interface{} {
	return l.item[l.Count()-1]
}

func (l *List) Append(item interface{}) {
	l.item = append(l.item, item)
}

func (l *List) Prepend(item interface{}) {
	newList := make([]interface{}, l.Count()+1)
	newList[0] = item
	copy(newList[1:], l.item)
	l.item = newList
}

func (l *List) Remove(index int) bool {
	if ok := l.indexOk(index); !ok {
		return false
	}
	copy(l.item[index:], l.item[index+1:])
	l.item = l.item[:l.Count()-1]
	return true
}

func (l *List) RemoveLast() bool {
	return l.Remove(l.Count() - 1)
}

func (l *List) RemoveFirst() bool {
	return l.Remove(0)
}

func (l *List) RemoveAll() {
	l.item = make([]interface{}, 0)
}
