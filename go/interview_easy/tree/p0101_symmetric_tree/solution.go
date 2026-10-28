// Package p0101 solves LeetCode 101. Symmetric Tree (Easy).
// https://leetcode.com/problems/symmetric-tree/
//
// Return true if a binary tree is a mirror image of itself around its
// center.
//
// Example 1:
//
//	Input:  root = [1,2,2,3,4,4,3]
//	Output: true
//
//	          1
//	         / \
//	        /   \
//	       2     2
//	      / \   / \
//	     3   4 4   3
//
// Example 2:
//
//	Input:  root = [1,2,2,null,3,null,3]
//	Output: false
//
//	        1
//	       / \
//	      2   2
//	       \   \
//	        3   3
//
// Constraints:
//
//	1 <= number of nodes <= 1000
//	-100 <= Node.val <= 100
//
// Follow-up: can you solve it both recursively and iteratively?
package p0101

import "github.com/ws4charlie/test4py/go/internal/bintree"

type TreeNode = bintree.TreeNode

// IsSymmetric checks that the left and right subtrees mirror each other.
func IsSymmetric(root *TreeNode) bool {
	return root == nil || mirror(root.Left, root.Right)
}

// mirror: roots match, outer children mirror, inner children mirror.
func mirror(a, b *TreeNode) bool {
	if a == nil || b == nil {
		return a == b
	}
	return a.Val == b.Val && mirror(a.Left, b.Right) && mirror(a.Right, b.Left)
}
