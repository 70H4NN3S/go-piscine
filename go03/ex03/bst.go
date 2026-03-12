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

func (b BST) PostOrder() []int {
	result := make([]int, 0)
	postOrder := make([]*TreeNode, 0)
	var prev *TreeNode
	n := b.Root

	for n != nil || len(postOrder) != 0 {
		if n.Left != nil && prev != n.Left && prev != n.Right {
			prev = n
			postOrder = append(postOrder, n)
			n = n.Left
		} else if n.Right != nil && n.Right != prev {
			prev = n
			postOrder = append(postOrder, n)
			n = n.Right
		} else {
			prev = n
			result = append(result, n.Val)
			if len(postOrder) == 0 {
				break
			}
			n = postOrder[len(postOrder)-1]
			postOrder = postOrder[:len(postOrder)-1]
		}
	}
	return result
}

func (b BST) IsBST() bool {
	return b.IsBalanced() && checkValuesOrder(b.Root)
}

func (b BST) IsBalanced() bool {
	result := recursiveIsBalancedCheck(b.Root)

	return result != -1
}

func recursiveIsBalancedCheck(node *TreeNode) int {
	if node == nil {
		return 0
	}
	leftSide := recursiveIsBalancedCheck(node.Left)
	if leftSide == -1 {
		return -1
	}
	rightSide := recursiveIsBalancedCheck(node.Right)
	if rightSide == -1 {
		return -1
	}

	if rightSide > leftSide {
		if rightSide-leftSide > 1 {
			return -1
		}
		return 1 + rightSide
	} else {
		if leftSide-rightSide > 1 {
			return -1
		}
		return 1 + leftSide
	}
}

func (b BST) Height() int {
	return recursiveHeight(b.Root)
}

func recursiveHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}
	leftSide := recursiveHeight(node.Left)
	rightSide := recursiveHeight(node.Right)
	if leftSide > rightSide {
		return 1 + leftSide
	}
	return 1 + rightSide
}

func checkValuesOrder(node *TreeNode) bool {
	if node == nil {
		return true
	}
	if node.Left == nil && node.Right == nil {
		return true
	}
	if node.Left == nil {
		if node.Val >= node.Right.Val {
			return false
		}
		return checkValuesOrder(node.Right)
	}
	if node.Right == nil {
		if node.Left.Val >= node.Val {
			return false
		}
		return checkValuesOrder(node.Left)
	}
	if node.Left.Val >= node.Val || node.Val >= node.Right.Val {
		return false
	}
	return checkValuesOrder(node.Left) && checkValuesOrder(node.Right)
}
