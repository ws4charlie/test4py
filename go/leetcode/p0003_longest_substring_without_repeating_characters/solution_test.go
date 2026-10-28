package p0003

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
		{"abba", 2},
	}
	for _, tt := range tests {
		if got := LengthOfLongestSubstring(tt.s); got != tt.want {
			t.Errorf("LengthOfLongestSubstring(%q) = %d, want %d", tt.s, got, tt.want)
		}
	}
}
