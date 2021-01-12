package tree

import (
	"strconv"
	"strings"
)

type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

func New(rootValue int) *BinaryTree {
	return &BinaryTree{rootValue, nil, nil}
}

func (bt *BinaryTree) LeftHeight() int {
	if bt.Left == nil {
		return 0
	}

	return bt.LeftHeight() + 1
}

func (bt *BinaryTree) RightHeight() int {
	if bt.Right == nil {
		return 0
	}

	return bt.RightHeight() + 1
}

func (bt *BinaryTree) Height() int {
	rh := bt.RightHeight()
	if lh := bt.LeftHeight(); lh > rh {
		return lh
	}

	return rh
}

func (bt *BinaryTree) BalanceFactor() int {
	return bt.LeftHeight() - bt.RightHeight()
}

func (bt *BinaryTree) Replace(oldChild *BinaryTree, newChild *BinaryTree) bool {
	if bt.Left == oldChild {
		bt.Left = newChild
		return true
	}

	if bt.Right == oldChild {
		bt.Right = newChild
		return true
	}

	return false
}

func (bt *BinaryTree) Remove(child *BinaryTree) bool {
	return bt.Replace(child, nil)
}

func (bt *BinaryTree) Walk(ch chan int) {
	if bt == nil {
		return
	}

	bt.Left.Walk(ch)
	ch <- bt.Value
	bt.Right.Walk(ch)
}

func (bt *BinaryTree) Walker() <-chan int {
	ch := make(chan int)
	go func() {
		bt.Walk(ch)
		close(ch)
	}()
	return ch
}

func (bt *BinaryTree) String(child *BinaryTree) string {
	ch := bt.Walker()
	var values []string
	for c := range ch {
		values = append(values, strconv.Itoa(c))
	}

	return strings.Join(values, " ")
}
