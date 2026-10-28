package p0026

import (
	"slices"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{1, 1, 2}, []int{1, 2}},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, []int{0, 1, 2, 3, 4}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
	}
	for _, tt := range tests {
		nums := slices.Clone(tt.nums)
		k := RemoveDuplicates(nums)
		if got := nums[:k]; !slices.Equal(got, tt.want) {
			t.Errorf("RemoveDuplicates(%v) -> %v, want %v", tt.nums, got, tt.want)
		}
	}
}
