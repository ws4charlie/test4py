package p0206

import (
	"slices"
	"testing"

	"github.com/ws4charlie/test4py/go/internal/linkedlist"
)

func TestReverseList(t *testing.T) {
	tests := []struct {
		vals, want []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{1}, []int{1}},
		{nil, nil},
	}
	impls := map[string]func(*ListNode) *ListNode{
		"ReverseList":          ReverseList,
		"ReverseListRecursive": ReverseListRecursive,
	}
	for name, f := range impls {
		for _, tt := range tests {
			got := linkedlist.ToSlice(f(linkedlist.FromSlice(tt.vals)))
			if !slices.Equal(got, tt.want) {
				t.Errorf("%s(%v) = %v, want %v", name, tt.vals, got, tt.want)
			}
		}
	}
}
