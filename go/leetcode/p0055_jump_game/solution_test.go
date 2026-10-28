package p0055

import "testing"

func TestCanJump(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		{[]int{2, 3, 1, 1, 4}, true},
		{[]int{3, 2, 1, 0, 4}, false},
		{[]int{0}, true},
		{[]int{2, 0, 0}, true},
		{[]int{1, 0, 1, 0}, false},
	}
	impls := map[string]func([]int) bool{
		"CanJump":   CanJump,
		"CanJumpDP": CanJumpDP,
	}
	for name, f := range impls {
		for _, tt := range tests {
			if got := f(tt.nums); got != tt.want {
				t.Errorf("%s(%v) = %v, want %v", name, tt.nums, got, tt.want)
			}
		}
	}
}
