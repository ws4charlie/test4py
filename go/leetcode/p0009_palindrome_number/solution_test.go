package p0009

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		x    int
		want bool
	}{
		{121, true},
		{-121, false},
		{10, false},
		{0, true},
		{1221, true},
		{12321, true},
		{123, false},
	}
	for _, tt := range tests {
		if got := IsPalindrome(tt.x); got != tt.want {
			t.Errorf("IsPalindrome(%d) = %v, want %v", tt.x, got, tt.want)
		}
	}
}
