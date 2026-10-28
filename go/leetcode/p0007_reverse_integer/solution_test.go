package p0007

import (
	"math"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		x, want int
	}{
		{123, 321},
		{-123, -321},
		{120, 21},
		{0, 0},
		{1534236469, 0}, // 9646324351 overflows
		{-2147483648, 0},
		{1463847412, 2147483641},
		{-1463847412, -2147483641},
		{math.MaxInt32, 0},
	}
	for _, tt := range tests {
		if got := Reverse(tt.x); got != tt.want {
			t.Errorf("Reverse(%d) = %d, want %d", tt.x, got, tt.want)
		}
	}
}
