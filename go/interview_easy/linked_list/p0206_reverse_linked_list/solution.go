// Package p0206 solves LeetCode 206. Reverse Linked List (Easy).
// https://leetcode.com/problems/reverse-linked-list/
//
// Reverse a singly linked list and return the new head.
//
// Example 1:
//
//	Input:  head = [1,2,3,4,5]
//	Output: [5,4,3,2,1]
//
//	    before: 1 -> 2 -> 3 -> 4 -> 5
//	    after:  5 -> 4 -> 3 -> 2 -> 1
//
// Example 2:
//
//	Input:  head = [1,2]
//	Output: [2,1]
//
// Example 3:
//
//	Input:  head = []
//	Output: []
//
// Constraints:
//
//	0 <= number of nodes <= 5000
//	-5000 <= Node.val <= 5000
//
// Follow-up: can you do it both iteratively and recursively?
package p0206

import "github.com/ws4charlie/test4py/go/internal/linkedlist"

type ListNode = linkedlist.ListNode

// ReverseList points each node back at the previous one.
func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		head.Next, prev, head = prev, head, head.Next
	}
	return prev
}

// ReverseListRecursive reverses the rest, then hooks head on at the end.
func ReverseListRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := ReverseListRecursive(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
