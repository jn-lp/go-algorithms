package linkedlist

type LinkedListElement struct {
	Value interface{}
	next  *LinkedListElement
}

func (e *LinkedListElement) Next() *LinkedListElement {
	return e.next
}

type LinkedList struct {
	Head *LinkedListElement
	len  int
}

func New() *LinkedList {
	return new(LinkedList)
}

func (l *LinkedList) Len() int {
	return l.len
}

func (l *LinkedList) Front() *LinkedListElement {
	if l.len == 0 {
		return nil
	}

	return l.Head
}

func (l *LinkedList) Back() *LinkedListElement {
	if l.len == 0 {
		return nil
	}

	current := l.Head
	for current.next != nil {
		current = current.next
	}
	return current
}

func (l *LinkedList) Prepend(Value interface{}) {
	l.Head = &LinkedListElement{Value, l.Head}
	l.len++
}

func (l *LinkedList) Append(Value interface{}) {
	current := l.Head
	for current.next != nil {
		current = current.next
	}
	current.next = &LinkedListElement{Value, nil}
	l.len++
}

//TODO: simplify Delete
func (l *LinkedList) Delete(e *LinkedListElement) (deleted *LinkedListElement) {
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

func (l *LinkedList) DeleteHead() (deleted *LinkedListElement) {
	return l.Delete(l.Front())
}

func (l *LinkedList) DeleteTail() (deleted *LinkedListElement) {
	return l.Delete(l.Back())
}

func (l *LinkedList) Reverse() {
	var prev, next *LinkedListElement
	for current := l.Head; current != nil; prev, current = current, next {
		next, current.next = current.next, prev
	}
	l.Head = prev
}

func (l *LinkedList) Slice() (s []interface{}) {
	for current := l.Head; current != nil; current = current.next {
		s = append(s, current.Value)
	}
	return
}
