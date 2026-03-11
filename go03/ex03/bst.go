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
	if val == b.Root.Val {
		return
	}
	if b.Root == nil {
		b.Root = &TreeNode{val, nil, nil}
	}

	for n := b.Root; true; {
		if n.Val == val {
			break
		}
		if n.Val > val {
			if n.Left == nil {
				n.Left = &TreeNode{val, nil, nil}
				break
			}
			n = n.Left
		}
		if n.Val < val {
			if n.Right == nil {
				n.Right = &TreeNode{val, nil, nil}
				break
			}
			n = n.Right
		}
	}
}

func (b BST) Search(val int) bool {
	for n := b.Root; n != nil; {
		if n.Val == val {
			return true
		}
		if n.Val > val {
			n = n.Left
		}
		if n.Val < val {
			n = n.Right
		}
	}
	return false
}
