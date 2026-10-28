package p0033

import "testing"

func TestSearch(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 4, 0},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 2, 6},
		{[]int{1}, 0, -1},
		{[]int{1}, 1, 0},
		{[]int{3, 1}, 1, 1},
		{[]int{1, 2, 3, 4, 5}, 4, 3},
		{nil, 1, -1},
	}
	for _, tt := range tests {
		if got := Search(tt.nums, tt.target); got != tt.want {
			t.Errorf("Search(%v, %d) = %d, want %d", tt.nums, tt.target, got, tt.want)
		}
	}
}
