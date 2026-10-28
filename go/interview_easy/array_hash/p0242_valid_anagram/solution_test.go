package p0242

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		s, t string
		want bool
	}{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
		{"a", "ab", false},
		{"", "", true},
	}
	for _, tt := range tests {
		if got := IsAnagram(tt.s, tt.t); got != tt.want {
			t.Errorf("IsAnagram(%q, %q) = %v, want %v", tt.s, tt.t, got, tt.want)
		}
	}
}
