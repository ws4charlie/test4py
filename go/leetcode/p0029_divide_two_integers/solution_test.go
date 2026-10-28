package p0029

import (
	"math"
	"testing"
)

func TestDivide(t *testing.T) {
	tests := []struct {
		dividend, divisor, want int
	}{
		{10, 3, 3},
		{7, -3, -2},
		{-7, 3, -2},
		{-7, -3, 2},
		{0, 5, 0},
		{1, 1, 1},
		{math.MinInt32, -1, math.MaxInt32},
		{math.MinInt32, 1, math.MinInt32},
		{math.MaxInt32, 2, 1073741823},
	}
	for _, tt := range tests {
		if got := Divide(tt.dividend, tt.divisor); got != tt.want {
			t.Errorf("Divide(%d, %d) = %d, want %d", tt.dividend, tt.divisor, got, tt.want)
		}
	}
}
