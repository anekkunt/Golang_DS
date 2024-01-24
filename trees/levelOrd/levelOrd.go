package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderTraversal(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:] // Dequeue

		result = append(result, node.Val)

		if node.Left != nil {
			queue = append(queue, node.Left) // Enqueue left child
		}
		if node.Right != nil {
			queue = append(queue, node.Right) // Enqueue right child
		}
	}

	return result
}

func main() {
	// Example tree creation
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}

	fmt.Println("Level Order Traversal:", levelOrderTraversal(root))
}
