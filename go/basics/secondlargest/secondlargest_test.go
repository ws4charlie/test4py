package secondlargest

import "testing"

func TestFind(t *testing.T) {
	tests := []struct {
		nums   []int
		want   int
		wantOK bool
	}{
		{[]int{4, 7, 3, 6, 8, 0, 2, 6}, 7, true},
		{[]int{1, 2}, 1, true},
		{[]int{2, 1}, 1, true},
		{[]int{5, 5}, 5, true},
		{[]int{8, 8, 3}, 8, true},
		{[]int{-3, -1, -2}, -2, true},
		{[]int{1}, 0, false},
		{nil, 0, false},
	}
	for _, tt := range tests {
		got, ok := Find(tt.nums)
		if got != tt.want || ok != tt.wantOK {
			t.Errorf("Find(%v) = %d, %v, want %d, %v", tt.nums, got, ok, tt.want, tt.wantOK)
		}
	}
}
