// Package bintree holds the LeetCode binary tree node and helpers to
// convert to and from LeetCode's level-order array form in tests.
package bintree

import "math"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// Null marks a missing node in a level-order slice (LeetCode's "null").
const Null = math.MinInt

// FromLevelOrder builds a tree from LeetCode level order, e.g.
// [3, 9, 20, Null, Null, 15, 7].
func FromLevelOrder(vals []int) *TreeNode {
	if len(vals) == 0 || vals[0] == Null {
		return nil
	}
	root := &TreeNode{Val: vals[0]}
	queue := []*TreeNode{root}
	for i := 1; len(queue) > 0 && i < len(vals); {
		n := queue[0]
		queue = queue[1:]
		for _, child := range []**TreeNode{&n.Left, &n.Right} {
			if i < len(vals) && vals[i] != Null {
				*child = &TreeNode{Val: vals[i]}
				queue = append(queue, *child)
			}
			i++
		}
	}
	return root
}

// LevelOrder returns the node values top to bottom, left to right,
// skipping missing nodes.
func LevelOrder(root *TreeNode) []int {
	var vals []int
	queue := []*TreeNode{}
	if root != nil {
		queue = append(queue, root)
	}
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		vals = append(vals, n.Val)
		for _, c := range []*TreeNode{n.Left, n.Right} {
			if c != nil {
				queue = append(queue, c)
			}
		}
	}
	return vals
}
