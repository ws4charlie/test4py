package p0017

import (
	"slices"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	tests := []struct {
		digits string
		want   []string
	}{
		{"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{"", nil},
		{"2", []string{"a", "b", "c"}},
	}
	impls := map[string]func(string) []string{
		"backtrack": LetterCombinations,
		"iterative": LetterCombinationsIterative,
	}
	for name, f := range impls {
		for _, tt := range tests {
			if got := f(tt.digits); !slices.Equal(got, tt.want) {
				t.Errorf("%s(%q) = %v, want %v", name, tt.digits, got, tt.want)
			}
		}
	}
}
