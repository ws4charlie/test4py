// Package p0104 solves LeetCode 104. Maximum Depth of Binary Tree (Easy).
// https://leetcode.com/problems/maximum-depth-of-binary-tree/
//
// Return the maximum depth of a binary tree: the number of nodes on the
// longest path from the root down to a leaf.
//
// Example 1:
//
//	Input:  root = [3,9,20,null,null,15,7]
//	Output: 3
//
//	        3
//	       / \
//	      9   20
//	         /  \
//	        15   7
//
// Example 2:
//
//	Input:  root = [1,null,2]
//	Output: 2
//
// Constraints:
//
//	0 <= number of nodes <= 10^4
//	-100 <= Node.val <= 100
package p0104

import "github.com/ws4charlie/test4py/go/internal/bintree"

type TreeNode = bintree.TreeNode

// MaxDepth is 1 + the deeper subtree.
func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(MaxDepth(root.Left), MaxDepth(root.Right))
}

// MaxDepthBFS counts levels.
func MaxDepthBFS(root *TreeNode) int {
	depth := 0
	var level []*TreeNode
	if root != nil {
		level = append(level, root)
	}
	for len(level) > 0 {
		depth++
		var next []*TreeNode
		for _, n := range level {
			for _, c := range []*TreeNode{n.Left, n.Right} {
				if c != nil {
					next = append(next, c)
				}
			}
		}
		level = next
	}
	return depth
}
