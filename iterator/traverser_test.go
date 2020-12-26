package iterator

func Example_Traverser_Print_all() {
	list := ListStub()
	iter := NewForwardIterator(list)

	procItem := func(item interface{}) bool {
		PrintEmplItem(item.(Employee))
		return true
	}
	t := NewTraverser(iter, procItem)
	t.Traverse()
	// Output:
	// Curtis Kiehn from Joellenton earns 25000$
	// Burl Morar from Anastasiaberg earns 40000$
	// Marge Marks from Geraldfort earns 55000$
}

func Example_Traverser_PrintNEmployees() {
	list := ListStub()
	list.Append(Employee{"Alice", "Montreal", 40000})
	list.Append(Employee{"Bob Marks", "New York", 55000})
	iter := NewForwardIterator(list)

	t := NewPrintNEmplotees(iter, 2)
	t.Traverse()
	// Output:
	// Curtis Kiehn from Joellenton earns 25000$
	// Burl Morar from Anastasiaberg earns 40000$
}
