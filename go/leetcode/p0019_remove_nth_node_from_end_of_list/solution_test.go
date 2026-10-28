package p0019

import (
	"slices"
	"testing"

	"github.com/ws4charlie/test4py/go/internal/linkedlist"
)

func TestRemoveNthFromEnd(t *testing.T) {
	tests := []struct {
		list []int
		n    int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{1, 2, 3, 5}},
		{[]int{1}, 1, nil},
		{[]int{1, 2}, 1, []int{1}},
		{[]int{1, 2}, 2, []int{2}},
	}
	for _, tt := range tests {
		got := linkedlist.ToSlice(RemoveNthFromEnd(linkedlist.FromSlice(tt.list), tt.n))
		if !slices.Equal(got, tt.want) {
			t.Errorf("RemoveNthFromEnd(%v, %d) = %v, want %v", tt.list, tt.n, got, tt.want)
		}
	}
}
