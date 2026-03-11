package main

import "fmt"

func main() {
	bST := BST{nil}
	bST.Insert(5)
	bST.Insert(3)
	bST.Insert(1)
	bST.Insert(2)
	bST.Insert(4)
	bST.Insert(6)
	bST.Insert(8)
	bST.Insert(7)

	fmt.Println("In Order: ", bST.InOrder())
	fmt.Println("Pre Order: ", bST.PreOrder())
}
