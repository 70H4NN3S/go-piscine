package main

import "fmt"

func main() {
	linkedList := List{nil, 0}

	fmt.Println(linkedList.String())
	if peeked, ok := linkedList.Peek(); ok {
		fmt.Println("Peeked on an empty list: ", peeked)
	}
	for i := range 20 {
		linkedList.Push(i)
	}
	fmt.Println(linkedList.String())
	if popped, ok := linkedList.Pop(); ok {
		fmt.Println("popped: ", popped)
	}
	if peeked, ok := linkedList.Peek(); ok {
		fmt.Println("Peeked: ", peeked)
	}
	linkedList.Reverse()
	fmt.Println(linkedList.String())
}
