// Package p0234 solves LeetCode 234. Palindrome Linked List (Easy).
// https://leetcode.com/problems/palindrome-linked-list/
//
// Return true if the values of a singly linked list read the same forwards
// and backwards.
//
// Example 1:
//
//	Input:  head = [1,2,2,1]
//	Output: true
//
//	    1 -> 2 -> 2 -> 1
//
// Example 2:
//
//	Input:  head = [1,2]
//	Output: false
//
//	    1 -> 2
//
// Constraints:
//
//	1 <= number of nodes <= 10^5
//	0 <= Node.val <= 9
//
// Follow-up: can you do it in O(n) time and O(1) space?
package p0234

import "github.com/ws4charlie/test4py/go/internal/linkedlist"

type ListNode = linkedlist.ListNode

// IsPalindrome uses O(1) space: find the middle with slow/fast pointers,
// reverse the second half, then compare it with the first half.
func IsPalindrome(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	var prev *ListNode
	for slow != nil {
		slow.Next, prev, slow = prev, slow, slow.Next
	}
	for ; prev != nil; prev, head = prev.Next, head.Next {
		if prev.Val != head.Val {
			return false
		}
	}
	return true
}
