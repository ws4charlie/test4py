package p0020

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"{[]}", true},
		{"(]", false},
		{"([)]", false},
		{")(", false},
		{"((", false},
		{"", true},
	}
	for _, tt := range tests {
		if got := IsValid(tt.s); got != tt.want {
			t.Errorf("IsValid(%q) = %v, want %v", tt.s, got, tt.want)
		}
	}
}
