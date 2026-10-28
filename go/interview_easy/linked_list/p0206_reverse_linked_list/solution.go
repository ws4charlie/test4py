// https://leetcode.com/problems/reverse-linked-list/
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
