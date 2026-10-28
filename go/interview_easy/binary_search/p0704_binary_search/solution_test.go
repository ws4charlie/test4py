package p0704

import "testing"

func TestSearch(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{-1, 0, 3, 5, 9, 12}, 9, 4},
		{[]int{-1, 0, 3, 5, 9, 12}, 2, -1},
		{[]int{-1, 0, 3, 5, 9, 12}, -1, 0},
		{[]int{-1, 0, 3, 5, 9, 12}, 12, 5},
		{[]int{5}, 5, 0},
		{nil, 1, -1},
	}
	for _, tt := range tests {
		if got := Search(tt.nums, tt.target); got != tt.want {
			t.Errorf("Search(%v, %d) = %d, want %d", tt.nums, tt.target, got, tt.want)
		}
	}
}
