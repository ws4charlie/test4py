// Package p0226 solves LeetCode 226. Invert Binary Tree (Easy).
// https://leetcode.com/problems/invert-binary-tree/
//
// Mirror a binary tree (swap every node's left and right children) and
// return its root.
//
// Example 1:
//
//	Input:  root = [4,2,7,1,3,6,9]
//	Output: [4,7,2,9,6,3,1]
//
//	          4                     4
//	         / \                   / \
//	        /   \                 /   \
//	       2     7     -->       7     2
//	      / \   / \             / \   / \
//	     1   3 6   9           9   6 3   1
//
// Example 2:
//
//	Input:  root = [2,1,3]
//	Output: [2,3,1]
//
//	      2                2
//	     / \     -->      / \
//	    1   3            3   1
//
// Example 3:
//
//	Input:  root = []
//	Output: []
//
// Constraints:
//
//	0 <= number of nodes <= 100
//	-100 <= Node.val <= 100
package p0226

import "github.com/ws4charlie/test4py/go/internal/bintree"

type TreeNode = bintree.TreeNode

// InvertTree swaps every node's children, bottom-up.
func InvertTree(root *TreeNode) *TreeNode {
	if root != nil {
		root.Left, root.Right = InvertTree(root.Right), InvertTree(root.Left)
	}
	return root
}
