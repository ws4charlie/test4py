package p0028

import "testing"

func TestStrStr(t *testing.T) {
	tests := []struct {
		haystack, needle string
		want             int
	}{
		{"hello", "ll", 2},
		{"sadbutsad", "sad", 0},
		{"leetcode", "leeto", -1},
		{"aaaaa", "bba", -1},
		{"abc", "", 0},
		{"", "a", -1},
		{"aabaaabaaac", "aabaaac", 4},
		{"mississippi", "issip", 4},
	}
	impls := map[string]func(string, string) int{
		"StrStr":    StrStr,
		"StrStrKMP": StrStrKMP,
	}
	for name, f := range impls {
		for _, tt := range tests {
			if got := f(tt.haystack, tt.needle); got != tt.want {
				t.Errorf("%s(%q, %q) = %d, want %d", name, tt.haystack, tt.needle, got, tt.want)
			}
		}
	}
}
