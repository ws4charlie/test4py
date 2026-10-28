package heap

import (
	"math/rand/v2"
	"slices"
	"testing"
)

func drain(h *MinHeap[int]) []int {
	var out []int
	for h.Len() > 0 {
		out = append(out, h.Pop())
	}
	return out
}

func TestPushPop(t *testing.T) {
	r := rand.New(rand.NewPCG(3, 4))
	for n := 0; n < 50; n++ {
		vals := make([]int, r.IntN(100))
		h := &MinHeap[int]{}
		for i := range vals {
			vals[i] = r.IntN(30)
			h.Push(vals[i])
		}
		slices.Sort(vals)
		if got := drain(h); !slices.Equal(got, vals) {
			t.Fatalf("pop order = %v, want %v", got, vals)
		}
	}
}

func TestFromSlice(t *testing.T) {
	vals := []int{9, 4, 7, 1, 8, 2, 2, 6}
	h := FromSlice(vals)
	if h.Peek() != 1 {
		t.Errorf("Peek() = %d, want 1", h.Peek())
	}
	want := slices.Clone(vals)
	slices.Sort(want)
	if got := drain(h); !slices.Equal(got, want) {
		t.Errorf("pop order = %v, want %v", got, want)
	}
	if vals[0] != 9 {
		t.Errorf("FromSlice modified its input")
	}
}

func TestInterleaved(t *testing.T) {
	h := &MinHeap[int]{}
	h.Push(5)
	h.Push(3)
	if v := h.Pop(); v != 3 {
		t.Fatalf("Pop() = %d, want 3", v)
	}
	h.Push(1)
	h.Push(4)
	if got, want := drain(h), []int{1, 4, 5}; !slices.Equal(got, want) {
		t.Errorf("drain = %v, want %v", got, want)
	}
}
