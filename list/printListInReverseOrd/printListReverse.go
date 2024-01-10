package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func printListInReverseOrd(h *ListNode) {
	if h == nil {
		return
	}

	printListInReverseOrd(h.Next)

	fmt.Printf("%d->", h.Val)

}

// ///////////////
func NewListNode(v int) *ListNode {
	return &ListNode{Val: v}
}

func main() {

	head := NewListNode(1)
	head.Next = NewListNode(2)
	head.Next.Next = NewListNode(3)
	head.Next.Next.Next = NewListNode(4)
	head.Next.Next.Next.Next = NewListNode(5)
	head.Next.Next.Next.Next.Next = NewListNode(6)

	printListInReverseOrd(head)
	fmt.Println("nil")

}
