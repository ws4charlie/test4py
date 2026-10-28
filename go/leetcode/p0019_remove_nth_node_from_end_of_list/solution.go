// Package p0019 solves LeetCode 19. Remove Nth Node From End of List (Medium).
// https://leetcode.com/problems/remove-nth-node-from-end-of-list/
//
// Remove the n-th node from the end of a linked list and return its head.
//
// Example 1:
//
//	Input:  head = [1,2,3,4,5], n = 2
//	Output: [1,2,3,5]
//
//	    1 -> 2 -> 3 -> 4 -> 5
//	                   x            (2nd from the end)
//	    1 -> 2 -> 3 ------> 5
//
// Example 2:
//
//	Input:  head = [1], n = 1
//	Output: []
//
// Example 3:
//
//	Input:  head = [1,2], n = 1
//	Output: [1]
//
// Constraints:
//
//	The list has sz nodes, 1 <= sz <= 30
//	0 <= Node.val <= 100
//	1 <= n <= sz
//
// Follow-up: can you do it in one pass?
package p0019

import "github.com/ws4charlie/test4py/go/internal/linkedlist"

type ListNode = linkedlist.ListNode

// RemoveNthFromEnd moves fast n+1 steps ahead of slow (both from a dummy
// head), so when fast hits the end slow sits right before the target.
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		slow, fast = slow.Next, fast.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}
