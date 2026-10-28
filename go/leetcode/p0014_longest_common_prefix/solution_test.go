package p0014

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		strs []string
		want string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"ab", "a"}, "a"},
		{[]string{"alone"}, "alone"},
		{nil, ""},
	}
	for _, tt := range tests {
		if got := LongestCommonPrefix(tt.strs); got != tt.want {
			t.Errorf("LongestCommonPrefix(%q) = %q, want %q", tt.strs, got, tt.want)
		}
	}
}
