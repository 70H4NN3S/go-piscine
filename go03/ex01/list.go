package main

import "strconv"

type Node struct {
	Val  int
	Next *Node
}

type List struct {
	Head *Node
	size int
}

func (ll *List) Push(val int) {
	ll.Head = &Node{val, ll.Head}
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

func (ll List) String() string {
	if ll.Head == nil {
		return "nil"
	}
	if ll.Head.Next == nil {
		return strconv.Itoa(ll.Head.Val) + " -> nil"
	}
	return strconv.Itoa(ll.Head.Val) + stringRecursive(*ll.Head.Next)
}

func stringRecursive(n Node) string {
	if n.Next == nil {
		return " -> " + strconv.Itoa(n.Val) + " -> nil"
	}
	return " -> " + strconv.Itoa(n.Val) + stringRecursive(*n.Next)
}

func (ll *List) Reverse() {
	if ll.Head == nil {
		return
	}
	if ll.Head.Next == nil {
		return
	}
	ll.Head = ll.Head.reverseRecursive(nil)
}

func (n *Node) reverseRecursive(prev *Node) *Node {
	if n.Next == nil {
		n.Next = prev
		return n
	}
	nextNode := n.Next
	n.Next = prev
	return nextNode.reverseRecursive(n)
}
