package doublylinkedlist

type Element struct {
	Value      interface{}
	prev, next *Element
}

func (e *Element) Next() *Element {
	return e.next
}

func (e *Element) Prev() *Element {
	return e.prev
}

type DoublyLinkedList struct {
	Head *Element
	len  int
}

func New() *DoublyLinkedList {
	return new(DoublyLinkedList)
}

func (l *DoublyLinkedList) Len() int {
	return l.len
}

func (l *DoublyLinkedList) Front() *Element {
	if l.len == 0 {
		return nil
	}

	return l.Head
}

func (l *DoublyLinkedList) Back() *Element {
	if l.len == 0 {
		return nil
	}

	return l.Head.prev
}
