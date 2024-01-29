package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(value int) *TreeNode {
	return &TreeNode{Val: value}
}

/*
Pre-Order Traversal (Iterative)
For pre-order traversal (Root, Left, Right), a stack can be used. Start by pushing the root node onto the stack. While the stack is not empty, pop a node,
print its value, and push its right and then left children (if they exist) onto the stack.
*/
func PreOrderTraversalIterative(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}

	stack := []*TreeNode{root} // Initialize stack with root node
	for len(stack) > 0 {
		current := stack[len(stack)-1]       // Pop the last element from stack
		stack = stack[:len(stack)-1]         // Slice off the last element
		result = append(result, current.Val) // Process the current node
		//fmt.Printf("%d ", current.Val)

		// Push right and then left child to stack (if they exist),so that left child is processed first
		if current.Right != nil {
			stack = append(stack, current.Right)
		}
		if current.Left != nil {
			stack = append(stack, current.Left)
		}
	}

	return result
}

/*
In-Order Traversal (Iterative)
In-order traversal (Left, Root, Right) iteratively is a bit more complex. You need a stack and a current node pointer. Traverse to the leftmost node,
pushing each node onto the stack. Then, pop from the stack, print the node, and proceed to its right subtree.*/

func InOrderTraversalIterative(root *TreeNode) []int {
	var result []int
	stack := []*TreeNode{}
	current := root
	for current != nil || len(stack) > 0 {
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		//fmt.Printf("%d ", current.Val)
		result = append(result, current.Val)
		current = current.Right
	}

	return result
}

/*
Push the root node to the first stack.
While the first stack is not empty:
Pop a node from the first stack and push it to the second stack.
Push its left and then right child to the first stack (if they exist).
Once the first stack is empty, pop and append the nodes from the second stack to the result list.
*/
func PostOrderTraversalIterative(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}
	stack1 := []*TreeNode{root}
	stack2 := []*TreeNode{}
	for len(stack1) > 0 {
		current := stack1[len(stack1)-1]
		stack1 = stack1[:len(stack1)-1]
		stack2 = append(stack2, current)
		if current.Left != nil {
			stack1 = append(stack1, current.Left)
		}
		if current.Right != nil {
			stack1 = append(stack1, current.Right)
		}
	}
	for len(stack2) > 0 {
		current := stack2[len(stack2)-1]
		stack2 = stack2[:len(stack2)-1]
		//fmt.Printf("%d ", current.Val)
		result = append(result, current.Val)
	}
	return result
}

func main() {
	// Manually constructing the BST
	// The structure of the tree is as follows:
	//           50
	//         /    \
	//       30      70
	//      /  \    /  \
	//    20   40  60   80

	root := NewTreeNode(50)
	root.Left = NewTreeNode(30)
	root.Right = NewTreeNode(70)
	root.Left.Left = NewTreeNode(20)
	root.Left.Right = NewTreeNode(40)
	root.Right.Left = NewTreeNode(60)
	root.Right.Right = NewTreeNode(80)

	fmt.Println("Pre-Order Traversal (Iterative):")
	PreOrderTraversalIterative(root)
	fmt.Println("\nIn-Order Traversal (Iterative):")
	InOrderTraversalIterative(root)
	fmt.Println("\nPost-Order Traversal (Iterative):")
	PostOrderTraversalIterative(root)
}
