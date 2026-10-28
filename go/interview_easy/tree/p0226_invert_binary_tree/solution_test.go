package p0226

import (
	"slices"
	"testing"

	"github.com/ws4charlie/test4py/go/internal/bintree"
)

func TestInvertTree(t *testing.T) {
	tests := []struct {
		tree, want []int
	}{
		{[]int{4, 2, 7, 1, 3, 6, 9}, []int{4, 7, 2, 9, 6, 3, 1}},
		{[]int{2, 1, 3}, []int{2, 3, 1}},
		{[]int{1, 2}, []int{1, 2}}, // shape flips; level order is the same
		{nil, nil},
	}
	for _, tt := range tests {
		got := InvertTree(bintree.FromLevelOrder(tt.tree))
		if vals := bintree.LevelOrder(got); !slices.Equal(vals, tt.want) {
			t.Errorf("InvertTree(%v) = %v, want %v", tt.tree, vals, tt.want)
		}
	}
	if root := InvertTree(bintree.FromLevelOrder([]int{1, 2})); root.Left != nil || root.Right.Val != 2 {
		t.Errorf("InvertTree([1,2]): left child should move to the right")
	}
}
