// Package binarytree has the classic traversals, recursive and iterative,
// ported from python/basics/binary_tree.py. Each returns the values in
// visit order instead of printing them.
package binarytree

type Node[T any] struct {
	Val         T
	Left, Right *Node[T]
}

// PreOrder: root, left, right.
func PreOrder[T any](root *Node[T]) []T {
	var res []T
	var walk func(n *Node[T])
	walk = func(n *Node[T]) {
		if n == nil {
			return
		}
		res = append(res, n.Val)
		walk(n.Left)
		walk(n.Right)
	}
	walk(root)
	return res
}

// PreOrderIter pushes right before left so left is popped first.
func PreOrderIter[T any](root *Node[T]) []T {
	if root == nil {
		return nil
	}
	var res []T
	stack := []*Node[T]{root}
	for len(stack) > 0 {
		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, n.Val)
		if n.Right != nil {
			stack = append(stack, n.Right)
		}
		if n.Left != nil {
			stack = append(stack, n.Left)
		}
	}
	return res
}

// InOrder: left, root, right. On a BST this gives sorted order.
func InOrder[T any](root *Node[T]) []T {
	var res []T
	var walk func(n *Node[T])
	walk = func(n *Node[T]) {
		if n == nil {
			return
		}
		walk(n.Left)
		res = append(res, n.Val)
		walk(n.Right)
	}
	walk(root)
	return res
}

// InOrderIter goes left as far as possible, visits, then moves to the
// right subtree.
func InOrderIter[T any](root *Node[T]) []T {
	var res []T
	var stack []*Node[T]
	for cur := root; cur != nil || len(stack) > 0; {
		for ; cur != nil; cur = cur.Left {
			stack = append(stack, cur)
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, cur.Val)
		cur = cur.Right
	}
	return res
}

// PostOrder: left, right, root.
func PostOrder[T any](root *Node[T]) []T {
	var res []T
	var walk func(n *Node[T])
	walk = func(n *Node[T]) {
		if n == nil {
			return
		}
		walk(n.Left)
		walk(n.Right)
		res = append(res, n.Val)
	}
	walk(root)
	return res
}

// PostOrderIter uses one stack plus the last visited node: a node is
// visited only once its right subtree is empty or was just finished.
func PostOrderIter[T any](root *Node[T]) []T {
	var res []T
	var stack []*Node[T]
	var last *Node[T]
	for cur := root; cur != nil || len(stack) > 0; {
		for ; cur != nil; cur = cur.Left {
			stack = append(stack, cur)
		}
		top := stack[len(stack)-1]
		if top.Right != nil && top.Right != last {
			cur = top.Right
			continue
		}
		stack = stack[:len(stack)-1]
		res = append(res, top.Val)
		last = top
	}
	return res
}

// PostOrderReverse does root, right, left with a stack, then reverses.
// Simpler, but needs the whole result before it can emit anything.
func PostOrderReverse[T any](root *Node[T]) []T {
	if root == nil {
		return nil
	}
	var res []T
	stack := []*Node[T]{root}
	for len(stack) > 0 {
		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, n.Val)
		if n.Left != nil {
			stack = append(stack, n.Left)
		}
		if n.Right != nil {
			stack = append(stack, n.Right)
		}
	}
	for l, r := 0, len(res)-1; l < r; l, r = l+1, r-1 {
		res[l], res[r] = res[r], res[l]
	}
	return res
}

// LevelOrder is BFS with a queue, top to bottom, left to right.
func LevelOrder[T any](root *Node[T]) []T {
	if root == nil {
		return nil
	}
	var res []T
	queue := []*Node[T]{root}
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		res = append(res, n.Val)
		if n.Left != nil {
			queue = append(queue, n.Left)
		}
		if n.Right != nil {
			queue = append(queue, n.Right)
		}
	}
	return res
}
