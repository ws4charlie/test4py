package p0104

import (
	"testing"

	"github.com/ws4charlie/test4py/go/internal/bintree"
)

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		tree []int
		want int
	}{
		{[]int{3, 9, 20, bintree.Null, bintree.Null, 15, 7}, 3},
		{[]int{1, bintree.Null, 2}, 2},
		{[]int{1}, 1},
		{nil, 0},
	}
	impls := map[string]func(*TreeNode) int{"MaxDepth": MaxDepth, "MaxDepthBFS": MaxDepthBFS}
	for name, f := range impls {
		for _, tt := range tests {
			if got := f(bintree.FromLevelOrder(tt.tree)); got != tt.want {
				t.Errorf("%s(%v) = %d, want %d", name, tt.tree, got, tt.want)
			}
		}
	}
}
