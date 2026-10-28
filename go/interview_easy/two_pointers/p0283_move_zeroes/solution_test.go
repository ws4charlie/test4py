package p0283

import (
	"slices"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		nums, want []int
	}{
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{[]int{0}, []int{0}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{0, 0, 1}, []int{1, 0, 0}},
	}
	for _, tt := range tests {
		got := slices.Clone(tt.nums)
		MoveZeroes(got)
		if !slices.Equal(got, tt.want) {
			t.Errorf("MoveZeroes(%v) -> %v, want %v", tt.nums, got, tt.want)
		}
	}
}
