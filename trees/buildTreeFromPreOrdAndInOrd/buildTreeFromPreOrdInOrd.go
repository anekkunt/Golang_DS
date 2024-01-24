package main

// TreeNode struct definition
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*Base Case: If the preorder or inorder array is empty, return nil.
Identify the Root: The first element in preorder is the root.
Find Root in inorder: Find the position of the root element in the inorder array. This divides the array into left and right subtrees.
Recursive Construction:
 Construct the left subtree using the left part of the inorder array and the corresponding elements in the preorder array.
 Construct the right subtree using the right part of the inorder array and the corresponding elements in the preorder array.
Return the Root: Create a new TreeNode with the root value, set the left and right subtrees from the recursive calls, and return this node*/

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	// The first element in preorder is the root
	rootVal := preorder[0]
	root := &TreeNode{Val: rootVal}

	// Find the root in inorder to divide into left and right subtrees
	rootIndexInInorder := find(inorder, rootVal)

	root.Left = buildTree(preorder[1:1+rootIndexInInorder], inorder[:rootIndexInInorder])
	root.Right = buildTree(preorder[1+rootIndexInInorder:], inorder[rootIndexInInorder+1:])

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
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}

	tree := buildTree(preorder, inorder)
	_ = tree
	// Add a function to print the tree if needed for testing
	// printTree(tree)
}
