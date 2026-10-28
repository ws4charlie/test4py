package p0268

import "testing"

func TestMissingNumber(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{3, 0, 1}, 2},
		{[]int{0, 1}, 2},
		{[]int{9, 6, 4, 2, 3, 5, 7, 0, 1}, 8},
		{[]int{1}, 0},
	}
	impls := map[string]func([]int) int{"MissingNumber": MissingNumber, "MissingNumberSum": MissingNumberSum}
	for name, f := range impls {
		for _, tt := range tests {
			if got := f(tt.nums); got != tt.want {
				t.Errorf("%s(%v) = %d, want %d", name, tt.nums, got, tt.want)
			}
		}
	}
}
