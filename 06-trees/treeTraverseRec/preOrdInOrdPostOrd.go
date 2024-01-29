package main

import "fmt"

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(value int) *TreeNode {
	return &TreeNode{Value: value}
}

func PreOrderTraversal(n *TreeNode) {
	if n == nil {
		return
	}
	fmt.Printf("%d ", n.Value)
	PreOrderTraversal(n.Left)
	PreOrderTraversal(n.Right)
}

func InOrderTraversal(n *TreeNode) {
	if n == nil {
		return
	}
	InOrderTraversal(n.Left)
	fmt.Printf("%d ", n.Value)
	InOrderTraversal(n.Right)
}

func PostOrderTraversal(n *TreeNode) {
	if n == nil {
		return
	}
	PostOrderTraversal(n.Left)
	PostOrderTraversal(n.Right)
	fmt.Printf("%d ", n.Value)
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

	fmt.Println("Pre-Order Traversal:")
	PreOrderTraversal(root)
	fmt.Println("\nIn-Order Traversal:")
	InOrderTraversal(root)
	fmt.Println("\nPost-Order Traversal:")
	PostOrderTraversal(root)
	fmt.Println()
}

/*out put
Pre-Order Traversal:
50 30 20 40 70 60 80
In-Order Traversal:
20 30 40 50 60 70 80
Post-Order Traversal:
20 40 30 60 80 70 50

*/
