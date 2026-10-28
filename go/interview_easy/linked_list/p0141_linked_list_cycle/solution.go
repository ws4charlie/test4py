// Package p0141 solves LeetCode 141. Linked List Cycle (Easy).
// https://leetcode.com/problems/linked-list-cycle/
//
// Return true if the linked list has a cycle: some node can be reached
// again by following next pointers. In the examples, pos is the index the
// tail links back to (-1 means no cycle); it is not passed to the function.
//
// Example 1:
//
//	Input:  head = [3,2,0,-4], pos = 1
//	Output: true
//
//	    3 -> 2 -> 0 -> -4
//	         ^          |
//	         |          |
//	         +----------+
//
// Example 2:
//
//	Input:  head = [1,2], pos = 0
//	Output: true
//
//	    1 -> 2
//	    ^    |
//	    +----+
//
// Example 3:
//
//	Input:  head = [1], pos = -1
//	Output: false
//
// Constraints:
//
//	0 <= number of nodes <= 10^4
//	-10^5 <= Node.val <= 10^5
//	pos is -1 or a valid index in the list.
//
// Follow-up: can you do it with O(1) memory?
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
