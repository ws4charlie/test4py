// Package p0021 solves LeetCode 21. Merge Two Sorted Lists (Easy).
// https://leetcode.com/problems/merge-two-sorted-lists/
//
// Merge two sorted linked lists into one sorted list by splicing their
// nodes together, and return its head.
//
// Example 1:
//
//	Input:  list1 = [1,2,4], list2 = [1,3,4]
//	Output: [1,1,2,3,4,4]
//
//	    list1:  1 -> 2 -> 4
//	    list2:  1 -> 3 -> 4
//	    merged: 1 -> 1 -> 2 -> 3 -> 4 -> 4
//
// Example 2:
//
//	Input:  list1 = [], list2 = []
//	Output: []
//
// Example 3:
//
//	Input:  list1 = [], list2 = [0]
//	Output: [0]
//
// Constraints:
//
//	Each list has 0 to 50 nodes.
//	-100 <= Node.val <= 100
//	Both lists are sorted in non-decreasing order.
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
