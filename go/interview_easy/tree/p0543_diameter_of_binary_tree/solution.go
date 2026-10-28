// https://leetcode.com/problems/diameter-of-binary-tree/
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
