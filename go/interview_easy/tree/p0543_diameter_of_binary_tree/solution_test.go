package p0543

import (
	"testing"

	"github.com/ws4charlie/test4py/go/internal/bintree"
)

func TestDiameterOfBinaryTree(t *testing.T) {
	tests := []struct {
		tree []int
		want int
	}{
		{[]int{1, 2, 3, 4, 5}, 3},
		{[]int{1, 2}, 1},
		{[]int{1}, 0},
		{nil, 0},
		// the longest path (5-3-2-4-6) doesn't go through the root
		{[]int{1, 2, bintree.Null, 3, 4, 5, bintree.Null, bintree.Null, 6}, 4},
	}
	for _, tt := range tests {
		if got := DiameterOfBinaryTree(bintree.FromLevelOrder(tt.tree)); got != tt.want {
			t.Errorf("DiameterOfBinaryTree(%v) = %d, want %d", tt.tree, got, tt.want)
		}
	}
}
