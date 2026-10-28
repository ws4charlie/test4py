// https://leetcode.com/problems/palindrome-linked-list/
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
