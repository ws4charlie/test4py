// Package sorting holds hand-written sorting algorithms and a shuffle,
// ported from python/basics/sort.py. All sorts are in place unless noted.
package sorting

import (
	"cmp"
	"math/rand/v2"
)

// Bubble sort: stable, O(n^2). Stops early once a pass makes no swaps.
func Bubble[T cmp.Ordered](a []T) {
	for i := 0; i < len(a)-1; i++ {
		swapped := false
		for j := len(a) - 1; j > i; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
				swapped = true
			}
		}
		if !swapped {
			return
		}
	}
}

// Insertion sort: stable, O(n^2), O(n) on nearly sorted input.
func Insertion[T cmp.Ordered](a []T) {
	for i := 1; i < len(a); i++ {
		v, j := a[i], i
		for ; j > 0 && a[j-1] > v; j-- {
			a[j] = a[j-1]
		}
		a[j] = v
	}
}

// Selection sort: unstable (the long-distance swap), O(n^2).
func Selection[T cmp.Ordered](a []T) {
	for i := 0; i < len(a)-1; i++ {
		m := i
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[m] {
				m = j
			}
		}
		a[i], a[m] = a[m], a[i]
	}
}

// Merge sort, top-down: stable, O(n log n), O(n) extra space.
func Merge[T cmp.Ordered](a []T) {
	buf := make([]T, len(a))
	var sort func(lo, hi int) // sorts a[lo:hi]
	sort = func(lo, hi int) {
		if hi-lo <= 1 {
			return
		}
		mid := lo + (hi-lo)/2
		sort(lo, mid)
		sort(mid, hi)
		merge(a, buf, lo, mid, hi)
	}
	sort(0, len(a))
}

// MergeBottomUp is iterative: merge runs of width 1, 2, 4, ...
// Stable, O(n log n), O(n) extra space.
func MergeBottomUp[T cmp.Ordered](a []T) {
	n := len(a)
	buf := make([]T, n)
	for width := 1; width < n; width *= 2 {
		for lo := 0; lo < n-width; lo += 2 * width {
			merge(a, buf, lo, lo+width, min(lo+2*width, n))
		}
	}
}

// merge combines sorted a[lo:mid] and a[mid:hi]. Taking from the left on
// ties (<=) keeps it stable.
func merge[T cmp.Ordered](a, buf []T, lo, mid, hi int) {
	copy(buf[lo:hi], a[lo:hi])
	i, j := lo, mid
	for k := lo; k < hi; k++ {
		if j >= hi || (i < mid && buf[i] <= buf[j]) {
			a[k] = buf[i]
			i++
		} else {
			a[k] = buf[j]
			j++
		}
	}
}

// Heap sort: unstable, O(n log n), O(1) extra space. Build a max-heap,
// then repeatedly swap the max to the end and sift the new root down.
func Heap[T cmp.Ordered](a []T) {
	n := len(a)
	for i := n/2 - 1; i >= 0; i-- {
		siftDown(a, i, n)
	}
	for end := n - 1; end > 0; end-- {
		a[0], a[end] = a[end], a[0]
		siftDown(a, 0, end)
	}
}

// siftDown restores the max-heap property for a[:end] starting at i.
func siftDown[T cmp.Ordered](a []T, i, end int) {
	for {
		big, l, r := i, 2*i+1, 2*i+2
		if l < end && a[l] > a[big] {
			big = l
		}
		if r < end && a[r] > a[big] {
			big = r
		}
		if big == i {
			return
		}
		a[i], a[big] = a[big], a[i]
		i = big
	}
}

// Quick sort with Lomuto partition: unstable, O(n log n) average,
// O(n^2) worst case (e.g. already sorted input with last-element pivot).
func Quick[T cmp.Ordered](a []T) {
	if len(a) <= 1 {
		return
	}
	p := partition(a)
	Quick(a[:p])
	Quick(a[p+1:])
}

// QuickIterative replaces recursion with an explicit stack of ranges.
func QuickIterative[T cmp.Ordered](a []T) {
	type span struct{ lo, hi int } // a[lo:hi]
	stack := []span{{0, len(a)}}
	for len(stack) > 0 {
		s := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if s.hi-s.lo <= 1 {
			continue
		}
		p := s.lo + partition(a[s.lo:s.hi])
		stack = append(stack, span{s.lo, p}, span{p + 1, s.hi})
	}
}

// partition uses the last element as pivot; everything < pivot ends up
// left of the returned index, everything >= pivot right of it.
func partition[T cmp.Ordered](a []T) int {
	last := len(a) - 1
	pivot, i := a[last], 0
	for j := 0; j < last; j++ {
		if a[j] < pivot {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	a[i], a[last] = a[last], a[i]
	return i
}

// Shuffle is Fisher-Yates: every permutation is equally likely.
func Shuffle[T any](a []T) {
	for i := len(a) - 1; i > 0; i-- {
		j := rand.IntN(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}
