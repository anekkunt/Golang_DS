package main

import "fmt"

func findMiddleNode(h *ListNode) *ListNode {
	fp, sp := h, h

	for fp != nil && fp.Next != nil {
		fp = fp.Next.Next
		sp = sp.Next
	}

	return sp

}

// /////////////////////////////////////////////
type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(val int) *ListNode {
	return &ListNode{Val: val}
}

func main() {

	head := NewListNode(1)
	head.Next = NewListNode(2)
	head.Next.Next = NewListNode(3)
	head.Next.Next.Next = NewListNode(4)
	head.Next.Next.Next = NewListNode(5)

	mid := findMiddleNode(head)
	fmt.Println("middle node:", mid)

}
