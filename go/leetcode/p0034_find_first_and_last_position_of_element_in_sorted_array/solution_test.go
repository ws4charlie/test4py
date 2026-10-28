package p0034

import (
	"slices"
	"testing"
)

func TestSearchRange(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
		{[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
		{[]int{5, 7, 7, 8, 8, 10}, 11, []int{-1, -1}},
		{[]int{2, 2, 2, 2}, 2, []int{0, 3}},
		{[]int{1}, 1, []int{0, 0}},
		{nil, 0, []int{-1, -1}},
	}
	for _, tt := range tests {
		if got := SearchRange(tt.nums, tt.target); !slices.Equal(got, tt.want) {
			t.Errorf("SearchRange(%v, %d) = %v, want %v", tt.nums, tt.target, got, tt.want)
		}
	}
}
