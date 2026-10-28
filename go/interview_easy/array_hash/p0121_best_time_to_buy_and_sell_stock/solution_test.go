package p0121

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		prices []int
		want   int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
		{[]int{2}, 0},
		{[]int{2, 4, 1}, 2},
	}
	for _, tt := range tests {
		if got := MaxProfit(tt.prices); got != tt.want {
			t.Errorf("MaxProfit(%v) = %d, want %d", tt.prices, got, tt.want)
		}
	}
}
