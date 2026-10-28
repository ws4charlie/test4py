package p0088

import (
	"slices"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		nums1 []int
		m     int
		nums2 []int
		want  []int
	}{
		{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, []int{1, 2, 2, 3, 5, 6}},
		{[]int{1}, 1, nil, []int{1}},
		{[]int{0}, 0, []int{1}, []int{1}},
		{[]int{4, 5, 6, 0, 0, 0}, 3, []int{1, 2, 3}, []int{1, 2, 3, 4, 5, 6}},
	}
	for _, tt := range tests {
		got := slices.Clone(tt.nums1)
		Merge(got, tt.m, tt.nums2, len(tt.nums2))
		if !slices.Equal(got, tt.want) {
			t.Errorf("Merge(%v, %d, %v) -> %v, want %v", tt.nums1, tt.m, tt.nums2, got, tt.want)
		}
	}
}
