//https://leetcode.com/problems/linked-list-cycle/

package main

type ListNode struct {
	val  int
	Next *ListNode
}

func hasCycleAndRemoveIt(head *ListNode) {

	fp, sp := head, head

	loop := false
	for fp != nil && fp.Next != nil {
		fp = fp.Next.Next // Move fast pointer by two steps
		sp = sp.Next      // Move slow pointer by one step

		if fp == sp { // Check after moving the pointers
			loop = true // Cycle detected
			break
		}
	}
	if !loop {
		return
	}

	//to remove loop
	start := head
	for sp.Next != start.Next {
		sp = sp.Next
		start = start.Next
	}
	// Break the loop
	sp.Next = nil

}
