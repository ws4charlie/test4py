package binarytree

import (
	"slices"
	"testing"
)

// sample builds the tree from the Python file:
//
//	     F
//	   /   \
//	  B     G
//	 / \     \
//	A   D     I
//	   / \   /
//	  C   E H
func sample() *Node[string] {
	n := func(v string, l, r *Node[string]) *Node[string] { return &Node[string]{v, l, r} }
	return n("F",
		n("B", n("A", nil, nil), n("D", n("C", nil, nil), n("E", nil, nil))),
		n("G", nil, n("I", n("H", nil, nil), nil)))
}

func split(s string) []string {
	var out []string
	for _, c := range s {
		out = append(out, string(c))
	}
	return out
}

func TestTraversals(t *testing.T) {
	tests := []struct {
		name string
		f    func(*Node[string]) []string
		want string
	}{
		{"PreOrder", PreOrder[string], "FBADCEGIH"},
		{"PreOrderIter", PreOrderIter[string], "FBADCEGIH"},
		{"InOrder", InOrder[string], "ABCDEFGHI"},
		{"InOrderIter", InOrderIter[string], "ABCDEFGHI"},
		{"PostOrder", PostOrder[string], "ACEDBHIGF"},
		{"PostOrderIter", PostOrderIter[string], "ACEDBHIGF"},
		{"PostOrderReverse", PostOrderReverse[string], "ACEDBHIGF"},
		{"LevelOrder", LevelOrder[string], "FBGADICEH"},
	}
	for _, tt := range tests {
		if got := tt.f(sample()); !slices.Equal(got, split(tt.want)) {
			t.Errorf("%s = %v, want %v", tt.name, got, split(tt.want))
		}
		if got := tt.f(nil); len(got) != 0 {
			t.Errorf("%s(nil) = %v, want empty", tt.name, got)
		}
	}
}
