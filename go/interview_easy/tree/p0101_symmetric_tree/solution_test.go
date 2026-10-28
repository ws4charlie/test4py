package p0101

import (
	"testing"

	"github.com/ws4charlie/test4py/go/internal/bintree"
)

func TestIsSymmetric(t *testing.T) {
	tests := []struct {
		tree []int
		want bool
	}{
		{[]int{1, 2, 2, 3, 4, 4, 3}, true},
		{[]int{1, 2, 2, bintree.Null, 3, bintree.Null, 3}, false},
		{[]int{1}, true},
		{[]int{1, 2, 3}, false},
		{nil, true},
	}
	for _, tt := range tests {
		if got := IsSymmetric(bintree.FromLevelOrder(tt.tree)); got != tt.want {
			t.Errorf("IsSymmetric(%v) = %v, want %v", tt.tree, got, tt.want)
		}
	}
}
