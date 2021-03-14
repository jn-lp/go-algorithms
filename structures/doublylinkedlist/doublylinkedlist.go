package doublylinkedlist

type DoublyLinkedListElement struct {
	Value      interface{}
	prev, next *DoublyLinkedListElement
}

func (e *DoublyLinkedListElement) Next() *DoublyLinkedListElement {
	return e.next
}

func (e *DoublyLinkedListElement) Prev() *DoublyLinkedListElement {
	return e.prev
}

type DoublyLinkedList struct {
	Head *DoublyLinkedListElement
	len  int
}

func New() *DoublyLinkedList {
	return new(DoublyLinkedList)
}

func (l *DoublyLinkedList) Len() int {
	return l.len
}

func (l *DoublyLinkedList) Front() *DoublyLinkedListElement {
	if l.len == 0 {
		return nil
	}

	return l.Head
}

func (l *DoublyLinkedList) Back() *DoublyLinkedListElement {
	if l.len == 0 {
		return nil
	}

	return l.Head.prev
}
