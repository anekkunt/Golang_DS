package main

import "fmt"

func reverseList(head *ListNode) *ListNode {
	cur := head
	var prev *ListNode
	for cur != nil {
		next := cur.Next
		cur.Next = prev

		prev = cur
		cur = next
	}
	return prev
}

// //////////////////////////////////
type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(v int) *ListNode {
	return &ListNode{Val: v}
}

func createLinkedList() *ListNode {
	head := NewListNode(1)

	node := head
	for i := 2; i <= 10; i++ {
		node.Next = NewListNode(i)
		node = node.Next
	}

	return head
}

func (h *ListNode) printList() {
	fmt.Println()
	for h != nil {
		fmt.Printf("%d->", h.Val)
		h = h.Next
	}
	fmt.Println("nil")
}

func main() {
	h := createLinkedList()
	h.printList()
	h = reverseList(h)
	h.printList()

}

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	// If the head needs to be deleted
	if head.Val == val {
		return head.Next
	}

	current := head
	for current.Next != nil {
		if current.Next.Val == val {
			current.Next = current.Next.Next
			break
		}
		current = current.Next
	}

	return head // Return the original head as it didn't change
}
