// Package p0543 solves LeetCode 543. Diameter of Binary Tree (Easy).
// https://leetcode.com/problems/diameter-of-binary-tree/
//
// Return the diameter of a binary tree: the number of edges on the longest
// path between any two nodes. The path may or may not pass through the
// root.
//
// Example 1:
//
//	Input:  root = [1,2,3,4,5]
//	Output: 3
//	Explanation: 4 -> 2 -> 1 -> 3 (or 5 -> 2 -> 1 -> 3) has 3 edges.
//
//	          1
//	         / \
//	        /   \
//	       2     3
//	      / \
//	     4   5
//
// Example 2:
//
//	Input:  root = [1,2]
//	Output: 1
//
// Constraints:
//
//	1 <= number of nodes <= 10^4
//	-100 <= Node.val <= 100
package p0543

import "github.com/ws4charlie/test4py/go/internal/bintree"

type TreeNode = bintree.TreeNode

// DiameterOfBinaryTree: the longest path through a node is
// depth(left) + depth(right). Compute depths bottom-up and keep the best
// sum seen; the answer need not pass through the root.
func DiameterOfBinaryTree(root *TreeNode) int {
	best := 0
	var depth func(n *TreeNode) int
	depth = func(n *TreeNode) int {
		if n == nil {
			return 0
		}
		l, r := depth(n.Left), depth(n.Right)
		best = max(best, l+r)
		return 1 + max(l, r)
	}
	depth(root)
	return best
}
