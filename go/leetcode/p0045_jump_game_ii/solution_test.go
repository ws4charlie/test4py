package p0045

import "testing"

func TestJump(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{2, 3, 1, 1, 4}, 2},
		{[]int{2, 3, 0, 1, 4}, 2},
		{[]int{0}, 0},
		{[]int{1, 2}, 1},
		{[]int{1, 1, 1, 1}, 3},
		{[]int{5, 9, 3, 2, 1, 0, 2, 3, 3, 1, 0, 0}, 3},
	}
	impls := map[string]func([]int) int{
		"Jump":     Jump,
		"JumpMemo": JumpMemo,
	}
	for name, f := range impls {
		for _, tt := range tests {
			if got := f(tt.nums); got != tt.want {
				t.Errorf("%s(%v) = %d, want %d", name, tt.nums, got, tt.want)
			}
		}
	}
}
