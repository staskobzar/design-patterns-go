package iterator

type ProcFunc func(item interface{}) bool

type ListTraverser struct {
	iter Iterator
	proc ProcFunc
}

func NewTraverser(iter Iterator, proc ProcFunc) *ListTraverser {
	return &ListTraverser{iter, proc}
}

func (t *ListTraverser) Traverse() {
	for t.iter.First(); !t.iter.IsDone(); t.iter.Next() {
		if ok := t.proc(t.iter.CurrentItem()); !ok {
			break
		}
	}
}

type ListNEmployees struct {
	*ListTraverser
	count int
	limit int
}

func NewPrintNEmplotees(iter Iterator, limit int) *ListNEmployees {
	lne := &ListNEmployees{count: 0, limit: limit}
	lne.ListTraverser = &ListTraverser{iter, lne.proc}
	return lne
}

func (lne *ListNEmployees) proc(item interface{}) bool {
	PrintEmplItem(item.(Employee))
	lne.count++
	if lne.count >= lne.limit {
		return false
	}
	return true
}
