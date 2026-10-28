package p0013

import "testing"

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"III", 3},
		{"IV", 4},
		{"IX", 9},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
		{"MMMCMXCIX", 3999},
	}
	for _, tt := range tests {
		if got := RomanToInt(tt.s); got != tt.want {
			t.Errorf("RomanToInt(%q) = %d, want %d", tt.s, got, tt.want)
		}
	}
}
