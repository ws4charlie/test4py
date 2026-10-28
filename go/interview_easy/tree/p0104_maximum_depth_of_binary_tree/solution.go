// https://leetcode.com/problems/maximum-depth-of-binary-tree/
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
