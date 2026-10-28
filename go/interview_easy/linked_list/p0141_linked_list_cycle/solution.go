// https://leetcode.com/problems/linked-list-cycle/
package p0141

import "github.com/ws4charlie/test4py/go/internal/linkedlist"

type ListNode = linkedlist.ListNode

// HasCycle is Floyd's algorithm: fast moves 2 steps, slow 1; they can
// only meet inside a cycle.
func HasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
