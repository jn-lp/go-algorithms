package linkedlist

type Element struct {
	Value interface{}
	next  *Element
}

func (e *Element) Next() *Element {
	return e.next
}

type LinkedList struct {
	Head *Element
	len  int
}

func New() *LinkedList {
	return new(LinkedList)
}

func (l *LinkedList) Len() int {
	return l.len
}

func (l *LinkedList) Front() *Element {
	if l.len == 0 {
		return nil
	}

	return l.Head
}

func (l *LinkedList) Back() *Element {
	if l.len == 0 {
		return nil
	}

	current := l.Head
	for current.next != nil {
		current = current.next
	}

	return current
}

func (l *LinkedList) Prepend(value interface{}) {
	l.Head = &Element{Value: value, next: l.Head}
	l.len++
}

func (l *LinkedList) Append(value interface{}) {
	current := l.Head
	for current.next != nil {
		current = current.next
	}

	current.next = &Element{Value: value}
	l.len++
}

func (l *LinkedList) Delete(e *Element) (deleted *Element) {
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

func (l *LinkedList) DeleteHead() (deleted *Element) {
	return l.Delete(l.Front())
}

func (l *LinkedList) DeleteTail() (deleted *Element) {
	return l.Delete(l.Back())
}

func (l *LinkedList) Reverse() {
	var prev, next *Element
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
