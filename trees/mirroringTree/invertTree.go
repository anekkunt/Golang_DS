package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// Swap the left and right children
	root.Left, root.Right = root.Right, root.Left

	// Recursively invert the left and right subtree
	invertTree(root.Left)
	invertTree(root.Right)

	return root
}
