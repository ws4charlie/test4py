package p0048

import (
	"reflect"
	"testing"
)

func clone(m [][]int) [][]int {
	out := make([][]int, len(m))
	for i, row := range m {
		out[i] = append([]int(nil), row...)
	}
	return out
}

func TestRotate(t *testing.T) {
	tests := []struct {
		matrix, want [][]int
	}{
		{[][]int{{1}}, [][]int{{1}}},
		{[][]int{{1, 2}, {3, 4}}, [][]int{{3, 1}, {4, 2}}},
		{
			[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			[][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
		},
		{
			[][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}},
			[][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}},
		},
	}
	impls := map[string]func([][]int){
		"Rotate":       Rotate,
		"RotateLayers": RotateLayers,
	}
	for name, f := range impls {
		for _, tt := range tests {
			got := clone(tt.matrix)
			f(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%s(%v) -> %v, want %v", name, tt.matrix, got, tt.want)
			}
		}
	}
}
