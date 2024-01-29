package main

// TreeNode struct definition
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}

	// The last element in postorder is the root
	rootVal := postorder[len(postorder)-1]
	root := &TreeNode{Val: rootVal}

	// Find the root in inorder to divide into left and right subtrees
	index := find(inorder, rootVal)

	root.Left = buildTree(inorder[:index], postorder[:index])
	root.Right = buildTree(inorder[index+1:], postorder[index:len(postorder)-1])

	return root
}

// Helper function to find the index of a value in a slice
func find(slice []int, value int) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1 // Not found
}

func main() {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}

	tree := buildTree(inorder, postorder)
	_ = tree
	// Add a function to print the tree if needed for testing
	// printTree(tree)
}
