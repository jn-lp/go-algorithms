package linkedlist

type element struct {
	Value interface{}
	next  *element
}

func (e *element) Next() *element {
	return e.next
}

type List struct {
	Head *element
	len  int
}

func New() *List {
	return new(List)
}

func (l *List) Len() int {
	return l.len
}

func (l *List) Front() *element {
	if l.len == 0 {
		return nil
	}

	return l.Head
}

func (l *List) Back() *element {
	if l.len == 0 {
		return nil
	}

	current := l.Head
	for current.next != nil {
		current = current.next
	}
	return current
}

func (l *List) Prepend(Value interface{}) {
	l.Head = &element{Value, l.Head}
	l.len++
}

func (l *List) Append(Value interface{}) {
	current := l.Head
	for current.next != nil {
		current = current.next
	}
	current.next = &element{Value, nil}
	l.len++
}

//TODO: simplify Delete
func (l *List) Delete(e *element) (deleted *element) {
	if l.Head == nil {
		return
	}

	for l.Head == e {
		deleted = l.Head
		l.Head = l.Head.next
	}

	current := l.Head

	if current != nil {
		for current.next != nil {
			if current.next == e {
				deleted = current.next
				current.next = current.next.next
			} else {
				current = current.next
			}
		}
	}

	l.len--
	return
}

func (l *List) DeleteHead() (deleted *element) {
	return l.Delete(l.Front())
}

func (l *List) DeleteTail() (deleted *element) {
	return l.Delete(l.Back())
}

func (l *List) Reverse() {
	var prev, next *element
	for current := l.Head; current != nil; prev, current = current, next {
		next, current.next = current.next, prev
	}
	l.Head = prev
}

func (l *List) Slice() (s []interface{}) {
	for current := l.Head; current != nil; current = current.next {
		s = append(s, current.Value)
	}
	return
}
