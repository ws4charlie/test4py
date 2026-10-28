// https://leetcode.com/problems/invert-binary-tree/
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
