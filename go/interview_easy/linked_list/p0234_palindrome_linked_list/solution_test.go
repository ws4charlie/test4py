package p0234

import (
	"testing"

	"github.com/ws4charlie/test4py/go/internal/linkedlist"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		vals []int
		want bool
	}{
		{[]int{1, 2, 2, 1}, true},
		{[]int{1, 2, 3, 2, 1}, true},
		{[]int{1, 2}, false},
		{[]int{1, 2, 3, 1}, false},
		{[]int{1}, true},
		{nil, true},
	}
	for _, tt := range tests {
		if got := IsPalindrome(linkedlist.FromSlice(tt.vals)); got != tt.want {
			t.Errorf("IsPalindrome(%v) = %v, want %v", tt.vals, got, tt.want)
		}
	}
}
