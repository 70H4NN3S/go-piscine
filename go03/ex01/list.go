package main

type Node struct {
	Val  int
	Next *Node
}

type List struct {
	Head *Node
	size int
}

func (ll *List) Push(val int) {
	node := Node{val, nil}
	if ll.size != 0 {
		node.Next = ll.Head
	}
	ll.Head = &node
	ll.size++
}

func (ll *List) Pop() (int, bool) {
	if ll.size == 0 {
		return 0, false
	}
	nextHead := ll.Head.Next
	val := ll.Head.Val
	ll.Head = nextHead
	ll.size--
	return val, true
}

func (ll List) Peek() (int, bool) {
	if ll.size == 0 {
		return 0, false
	}
	return ll.Head.Val, true
}

func (ll List) Len() int {
	return ll.size
}
