package main

import (
	"strconv"
	"strings"
)

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
	var sb strings.Builder
	for n := ll.Head; n != nil; n = n.Next {
		sb.WriteString(strconv.Itoa(n.Val))
		if n.Next != nil {
			sb.WriteString(" -> ")
		}
	}
	sb.WriteString(" -> nil")
	return sb.String()
}

func (ll *List) Reverse() {
	var prev *Node
	curr := ll.Head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	ll.Head = prev
}

func (ll *List) Contains(val int) bool {
	for n := ll.Head; n != nil; n = n.Next {
		if n.Val == val {
			return true
		}
	}
	return false
}
