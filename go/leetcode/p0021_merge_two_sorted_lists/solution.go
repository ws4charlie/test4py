// https://leetcode.com/problems/merge-two-sorted-lists/
package p0021

import "github.com/ws4charlie/test4py/go/internal/linkedlist"

type ListNode = linkedlist.ListNode

// MergeTwoLists splices the smaller head onto the tail until one list
// runs out, then attaches the rest of the other.
func MergeTwoLists(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			tail.Next, l1 = l1, l1.Next
		} else {
			tail.Next, l2 = l2, l2.Next
		}
		tail = tail.Next
	}
	if l1 != nil {
		tail.Next = l1
	} else {
		tail.Next = l2
	}
	return dummy.Next
}
