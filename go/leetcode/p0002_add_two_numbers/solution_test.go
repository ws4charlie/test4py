package p0002

import (
	"slices"
	"testing"

	"github.com/ws4charlie/test4py/go/internal/linkedlist"
)

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		l1, l2, want []int
	}{
		{[]int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8}},
		{[]int{0}, []int{0}, []int{0}},
		{[]int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}, []int{8, 9, 9, 9, 0, 0, 0, 1}},
	}
	for _, tt := range tests {
		got := linkedlist.ToSlice(AddTwoNumbers(linkedlist.FromSlice(tt.l1), linkedlist.FromSlice(tt.l2)))
		if !slices.Equal(got, tt.want) {
			t.Errorf("AddTwoNumbers(%v, %v) = %v, want %v", tt.l1, tt.l2, got, tt.want)
		}
	}
}
