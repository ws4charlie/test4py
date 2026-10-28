// Package heap is a hand-written binary min-heap, ported from
// python/basics/heap.py. Node i has children 2i+1 and 2i+2 and parent
// (i-1)/2.
package heap

import "cmp"

type MinHeap[T cmp.Ordered] struct {
	data []T
}

// FromSlice heapifies a copy of vals in O(n) by sifting down every
// non-leaf, last one first.
func FromSlice[T cmp.Ordered](vals []T) *MinHeap[T] {
	h := &MinHeap[T]{data: append([]T(nil), vals...)}
	for i := len(h.data)/2 - 1; i >= 0; i-- {
		h.siftDown(i)
	}
	return h
}

func (h *MinHeap[T]) Len() int { return len(h.data) }

// Peek returns the minimum without removing it. Panics if empty.
func (h *MinHeap[T]) Peek() T { return h.data[0] }

// Push appends v and sifts it up. O(log n).
func (h *MinHeap[T]) Push(v T) {
	h.data = append(h.data, v)
	h.siftUp(len(h.data) - 1)
}

// Pop removes and returns the minimum: move the last element to the root
// and sift it down. O(log n). Panics if empty.
func (h *MinHeap[T]) Pop() T {
	top := h.data[0]
	last := len(h.data) - 1
	h.data[0] = h.data[last]
	h.data = h.data[:last]
	h.siftDown(0)
	return top
}

func (h *MinHeap[T]) siftUp(i int) {
	for i > 0 {
		p := (i - 1) / 2
		if h.data[p] <= h.data[i] {
			return
		}
		h.data[i], h.data[p] = h.data[p], h.data[i]
		i = p
	}
}

func (h *MinHeap[T]) siftDown(i int) {
	n := len(h.data)
	for {
		small, l, r := i, 2*i+1, 2*i+2
		if l < n && h.data[l] < h.data[small] {
			small = l
		}
		if r < n && h.data[r] < h.data[small] {
			small = r
		}
		if small == i {
			return
		}
		h.data[i], h.data[small] = h.data[small], h.data[i]
		i = small
	}
}
