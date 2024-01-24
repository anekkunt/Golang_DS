package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	// Both trees are empty
	if p == nil && q == nil {
		return true
	}

	// One of the trees is empty, or the values at the current nodes are different
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}

	// Check recursively if the left subtrees and right subtrees are the same
	resultL := isSameTree(p.Left, q.Left)
	resultR := isSameTree(p.Right, q.Right)

	if resultL && resultR {
		return true
	}
	return false
}
