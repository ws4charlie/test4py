// https://leetcode.com/problems/symmetric-tree/
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
