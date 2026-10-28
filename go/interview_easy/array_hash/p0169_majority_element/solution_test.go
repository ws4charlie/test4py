package p0169

import "testing"

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{3, 2, 3}, 3},
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
		{[]int{1}, 1},
		{[]int{-1, -1, 5}, -1},
	}
	for _, tt := range tests {
		if got := MajorityElement(tt.nums); got != tt.want {
			t.Errorf("MajorityElement(%v) = %d, want %d", tt.nums, got, tt.want)
		}
	}
}
