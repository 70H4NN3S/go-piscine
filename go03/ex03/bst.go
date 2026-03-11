package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BST struct {
	Root *TreeNode
}

func (b *BST) Insert(val int) {
	if b.Root == nil {
		b.Root = &TreeNode{val, nil, nil}
		return
	}

	for n := b.Root; ; {
		if n.Val == val {
			return
		} else if n.Val > val {
			if n.Left == nil {
				n.Left = &TreeNode{val, nil, nil}
				return
			}
			n = n.Left
		} else {
			if n.Right == nil {
				n.Right = &TreeNode{val, nil, nil}
				return
			}
			n = n.Right
		}
	}
}

func (b BST) Search(val int) bool {
	for n := b.Root; n != nil; {
		if n.Val == val {
			return true
		} else if n.Val > val {
			n = n.Left
		} else {
			n = n.Right
		}
	}
	return false
}

func (b BST) InOrder() []int {
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	n := b.Root

	for n != nil || len(stack) != 0 {
		// walk all the way left, pushing nodes as we go
		for n != nil {
			stack = append(stack, n)
			n = n.Left
		}
		// visit the node on top of the stack
		n = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, n.Val)
		// now explore it's right subtree
		n = n.Right
	}
	return result
}

func (b BST) PreOrder() []int {
	result := make([]int, 0)
	preOrder := make([]*TreeNode, 0)
	n := b.Root
	preOrder = append(preOrder, nil)

	for n != nil || len(preOrder) != 0 {
		for n != nil {
			result = append(result, n.Val)
			if n.Right != nil {
				preOrder = append(preOrder, n.Right)
			}
			n = n.Left
		}
		n = preOrder[len(preOrder)-1]
		preOrder = preOrder[:len(preOrder)-1]
	}
	return result
}
