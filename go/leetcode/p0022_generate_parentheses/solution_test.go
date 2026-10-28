package p0022

import (
	"slices"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	tests := []struct {
		n    int
		want []string
	}{
		{1, []string{"()"}},
		{2, []string{"(())", "()()"}},
		{3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
	}
	for _, tt := range tests {
		if got := GenerateParenthesis(tt.n); !slices.Equal(got, tt.want) {
			t.Errorf("GenerateParenthesis(%d) = %v, want %v", tt.n, got, tt.want)
		}
	}
}
