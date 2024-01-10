//https://leetcode.com/problems/intersection-of-two-linked-lists/
/*
Calculate the Lengths of Both Lists: Traverse both lists to find their lengths.

Align the Start Points: If the lengths are different, advance the start of the longer list by the difference in lengths. This ensures that the two pointers are equally far from the end of the lists.

Traverse Together and Find Intersection: Traverse both lists in tandem, comparing nodes by reference (address). When you find nodes that are the same (identical memory addresses), that node is the intersection.

*/
//we can use hash map but memory consumtion O(n)
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA, lenB := length(headA), length(headB)
	for lenA > lenB {
		headA = headA.Next
		lenA--
	}
	for lenB > lenA {
		headB = headB.Next
		lenB--
	}

	for headA != nil && headB != nil {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}

	return nil
}

func length(head *ListNode) int {
	count := 0
	for head != nil {
		count++
		head = head.Next
	}
	return count
}
