/*
 1. Create a Dummy Node: This node acts as the placeholder for the start of the merged list.
 2. Use a Pointer for Merging: Initialize a pointer (let's call it list3) at the dummy node.
    list3 and Link Nodes: Iterate through both lists. At each step, compare the values of the current nodes of list1 and list2:

3. If list1's value is smaller or equal, link current's Next to list1's list3 node and move list1 to the next node.
4. If list2's value is smaller link list3's Next to list2's current node and move list2 to the next node.
5. Move list3 to its next node after each linking.
6. Handle Remaining Nodes: After one of the lists is exhausted, link the remaining part of the non-exhausted list to current.
7. Return the Merged List: Since the dummy node was the placeholder, return dummy.Next as the start of the merged list.
*/
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	dummy := &ListNode{}
	list3 := dummy

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			list3.Next = list1
			list1 = list1.Next
		} else {
			list3.Next = list2
			list2 = list2.Next
		}
		list3 = list3.Next
	}

	if list1 != nil {
		list3.Next = list1
	}
	if list2 != nil {
		list3.Next = list2
	}
	return dummy.Next

}

// //////////////////////
func NewListNode(v int) *ListNode {
	return &ListNode{Val: v}
}

func main() {
	//list1
	head1 := NewListNode(1)
	head1.Next = NewListNode(2)
	head1.Next.Next = NewListNode(3)
	head1.Next.Next.Next = NewListNode(4)
	head1.Next.Next.Next.Next = NewListNode(5)

	//list2
	head2 := NewListNode(3)
	head2.Next = NewListNode(4)
	head2.Next.Next = NewListNode(5)
	head2.Next.Next.Next = NewListNode(6)
	head2.Next.Next.Next.Next = NewListNode(7)

	head3 := mergeTwoLists(head1, head2)

	for head3 != nil {
		fmt.Printf("%d->", head3.Val)
		head3 = head3.Next
	}
	fmt.Println("nil")

}
