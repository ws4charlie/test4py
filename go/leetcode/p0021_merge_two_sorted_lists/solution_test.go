package p0021

import (
	"slices"
	"testing"

	"github.com/ws4charlie/test4py/go/internal/linkedlist"
)

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		l1, l2, want []int
	}{
		{[]int{1, 2, 4}, []int{1, 3, 4}, []int{1, 1, 2, 3, 4, 4}},
		{nil, nil, nil},
		{nil, []int{0}, []int{0}},
		{[]int{5}, []int{1, 2, 3}, []int{1, 2, 3, 5}},
	}
	for _, tt := range tests {
		got := linkedlist.ToSlice(MergeTwoLists(linkedlist.FromSlice(tt.l1), linkedlist.FromSlice(tt.l2)))
		if !slices.Equal(got, tt.want) {
			t.Errorf("MergeTwoLists(%v, %v) = %v, want %v", tt.l1, tt.l2, got, tt.want)
		}
	}
}
