package p0141

import (
	"testing"

	"github.com/ws4charlie/test4py/go/internal/linkedlist"
)

// build links the tail back to index pos (-1 means no cycle).
func build(vals []int, pos int) *ListNode {
	head := linkedlist.FromSlice(vals)
	if head == nil || pos < 0 {
		return head
	}
	var target, tail *ListNode
	for i, n := 0, head; n != nil; i, n = i+1, n.Next {
		if i == pos {
			target = n
		}
		tail = n
	}
	tail.Next = target
	return head
}

func TestHasCycle(t *testing.T) {
	tests := []struct {
		vals []int
		pos  int
		want bool
	}{
		{[]int{3, 2, 0, -4}, 1, true},
		{[]int{1, 2}, 0, true},
		{[]int{1}, 0, true},
		{[]int{1}, -1, false},
		{[]int{1, 2, 3}, -1, false},
		{nil, -1, false},
	}
	for _, tt := range tests {
		if got := HasCycle(build(tt.vals, tt.pos)); got != tt.want {
			t.Errorf("HasCycle(%v, pos=%d) = %v, want %v", tt.vals, tt.pos, got, tt.want)
		}
	}
}
