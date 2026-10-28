// https://leetcode.com/problems/remove-nth-node-from-end-of-list/
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
