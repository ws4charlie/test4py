package p0125

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{" ", true},
		{"0P", false},
		{"a.", true},
	}
	for _, tt := range tests {
		if got := IsPalindrome(tt.s); got != tt.want {
			t.Errorf("IsPalindrome(%q) = %v, want %v", tt.s, got, tt.want)
		}
	}
}
