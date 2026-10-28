// Package p0002 solves LeetCode 2. Add Two Numbers (Medium).
// https://leetcode.com/problems/add-two-numbers/
//
// Two non-empty linked lists each hold a non-negative integer, one digit per
// node, with the digits in reverse order (ones digit first). Return the sum
// of the two numbers as a linked list in the same reverse order. Neither
// number has leading zeros, except the number 0 itself.
//
// Example 1:
//
//	  2 -> 4 -> 3          342
//	+ 5 -> 6 -> 4        + 465
//	-------------        -----
//	  7 -> 0 -> 8          807
//
//	Input:  l1 = [2,4,3], l2 = [5,6,4]
//	Output: [7,0,8]
//	Explanation: 342 + 465 = 807
//
// Example 2:
//
//	Input:  l1 = [0], l2 = [0]
//	Output: [0]
//
// Example 3:
//
//	Input:  l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
//	Output: [8,9,9,9,0,0,0,1]
//
// Constraints:
//
//	Each list has 1 to 100 nodes.
//	0 <= Node.val <= 9
//	The numbers have no leading zeros, except the number 0 itself.
package p0002

import "github.com/ws4charlie/test4py/go/internal/linkedlist"

type ListNode = linkedlist.ListNode

// AddTwoNumbers adds digit by digit (least significant first) with a carry.
func AddTwoNumbers(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	for carry := 0; l1 != nil || l2 != nil || carry > 0; {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10
		tail.Next = &ListNode{Val: sum % 10}
		tail = tail.Next
	}
	return dummy.Next
}
