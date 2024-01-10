package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

var gH *ListNode
var gResult bool

func isPalindrome(h *ListNode) bool {
	gResult = true
	gH = h
	recursiveCheck(h)
	return gResult
}

func recursiveCheck(p *ListNode) {
	if p == nil {
		return
	}
	recursiveCheck(p.Next)
	if p.Val != gH.Val {
		gResult = false
	}
	gH = gH.Next
}

// /////////
func NewListNode(v int) *ListNode {
	return &ListNode{Val: v}
}

func main() {
	head := NewListNode(1)
	head.Next = NewListNode(2)
	head.Next.Next = NewListNode(3)
	head.Next.Next.Next = NewListNode(2)
	head.Next.Next.Next.Next = NewListNode(1)

	result := isPalindrome(head)

	if result {
		fmt.Println("list is palindrome")
	} else {
		fmt.Println("list is not palindrome")
	}

}
