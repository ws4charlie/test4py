// Package linkedlist holds the LeetCode singly-linked list node and
// helpers to build and read lists in tests.
package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

// FromSlice builds a list from vals; an empty slice gives nil.
func FromSlice(vals []int) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	for _, v := range vals {
		tail.Next = &ListNode{Val: v}
		tail = tail.Next
	}
	return dummy.Next
}

// ToSlice returns the list values in order; nil gives nil.
func ToSlice(head *ListNode) []int {
	var vals []int
	for ; head != nil; head = head.Next {
		vals = append(vals, head.Val)
	}
	return vals
}
