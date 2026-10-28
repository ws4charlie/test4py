package p0005

import "testing"

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"babad", "bab"},
		{"cbbd", "bb"},
		{"a", "a"},
		{"", ""},
		{"forgeeksskeegfor", "geeksskeeg"},
	}
	for _, tt := range tests {
		if got := LongestPalindrome(tt.s); got != tt.want {
			t.Errorf("LongestPalindrome(%q) = %q, want %q", tt.s, got, tt.want)
		}
	}
}
