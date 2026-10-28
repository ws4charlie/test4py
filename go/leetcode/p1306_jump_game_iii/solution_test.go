package p1306

import "testing"

func TestCanReach(t *testing.T) {
	tests := []struct {
		arr   []int
		start int
		want  bool
	}{
		{[]int{4, 2, 3, 0, 3, 1, 2}, 5, true},
		{[]int{4, 2, 3, 0, 3, 1, 2}, 0, true},
		{[]int{3, 0, 2, 1, 2}, 2, false},
		{[]int{0}, 0, true},
		{[]int{1, 1, 1}, 1, false},
	}
	impls := map[string]func([]int, int) bool{
		"CanReach":    CanReach,
		"CanReachDFS": CanReachDFS,
	}
	for name, f := range impls {
		for _, tt := range tests {
			if got := f(tt.arr, tt.start); got != tt.want {
				t.Errorf("%s(%v, %d) = %v, want %v", name, tt.arr, tt.start, got, tt.want)
			}
		}
	}
}
