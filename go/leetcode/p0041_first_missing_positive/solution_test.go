package p0041

import (
	"slices"
	"testing"
)

func TestFirstMissingPositive(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 0}, 3},
		{[]int{3, 4, -1, 1}, 2},
		{[]int{7, 8, 9, 11, 12}, 1},
		{[]int{0, 2, 2, 1, 1}, 3},
		{[]int{1}, 2},
		{[]int{}, 1},
	}
	for _, tt := range tests {
		if got := FirstMissingPositive(slices.Clone(tt.nums)); got != tt.want {
			t.Errorf("FirstMissingPositive(%v) = %d, want %d", tt.nums, got, tt.want)
		}
	}
}
