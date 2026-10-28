package graph

import (
	"reflect"
	"slices"
	"testing"
)

// sample is the graph from the Python file:
//
//	1---2---4---8
//	|   |      /
//	|   |    /
//	3   5  /
//	| \
//	|  \
//	6---7
//
// (5-8 is also an edge.)
func sample() *Graph {
	g := New()
	for v := 1; v <= 8; v++ {
		g.AddNode(v)
	}
	for _, e := range [][2]int{{1, 2}, {1, 3}, {2, 4}, {2, 5}, {4, 8}, {5, 8}, {3, 6}, {3, 7}, {6, 7}} {
		g.AddEdge(e[0], e[1])
	}
	return g
}

func TestDFS(t *testing.T) {
	if got, want := sample().DFS(1), []int{1, 2, 4, 8, 5, 3, 6, 7}; !slices.Equal(got, want) {
		t.Errorf("DFS(1) = %v, want %v", got, want)
	}
}

func TestBFS(t *testing.T) {
	if got, want := sample().BFS(1), []int{1, 2, 3, 4, 5, 6, 7, 8}; !slices.Equal(got, want) {
		t.Errorf("BFS(1) = %v, want %v", got, want)
	}
	if got, want := sample().BFS(8), []int{8, 4, 5, 2, 1, 3, 6, 7}; !slices.Equal(got, want) {
		t.Errorf("BFS(8) = %v, want %v", got, want)
	}
}

func TestComponents(t *testing.T) {
	g := sample()
	g.AddNode(9)
	g.AddEdge(10, 11)
	g.AddEdge(10, 11) // duplicate edge is ignored
	want := [][]int{{1, 2, 4, 8, 5, 3, 6, 7}, {9}, {10, 11}}
	if got := g.Components(); !reflect.DeepEqual(got, want) {
		t.Errorf("Components() = %v, want %v", got, want)
	}
	if got := g.BFS(10); !slices.Equal(got, []int{10, 11}) {
		t.Errorf("BFS(10) = %v, want [10 11]", got)
	}
}
