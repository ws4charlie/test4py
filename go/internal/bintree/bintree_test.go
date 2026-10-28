package bintree

import (
	"slices"
	"testing"
)

func TestFromLevelOrder(t *testing.T) {
	root := FromLevelOrder([]int{3, 9, 20, Null, Null, 15, 7})
	if root.Val != 3 || root.Left.Val != 9 || root.Right.Val != 20 ||
		root.Left.Left != nil || root.Right.Left.Val != 15 || root.Right.Right.Val != 7 {
		t.Fatalf("unexpected shape")
	}
	if got := LevelOrder(root); !slices.Equal(got, []int{3, 9, 20, 15, 7}) {
		t.Errorf("LevelOrder = %v", got)
	}
	if FromLevelOrder(nil) != nil || LevelOrder(nil) != nil {
		t.Errorf("empty tree should be nil")
	}
}
