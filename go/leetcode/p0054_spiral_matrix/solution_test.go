package p0054

import (
	"slices"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	tests := []struct {
		matrix [][]int
		want   []int
	}{
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
		{[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}, []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}},
		{[][]int{{1, 2, 3}}, []int{1, 2, 3}},
		{[][]int{{1}, {2}, {3}}, []int{1, 2, 3}},
		{[][]int{{1, 2}, {3, 4}, {5, 6}}, []int{1, 2, 4, 6, 5, 3}},
		{nil, nil},
	}
	for _, tt := range tests {
		if got := SpiralOrder(tt.matrix); !slices.Equal(got, tt.want) {
			t.Errorf("SpiralOrder(%v) = %v, want %v", tt.matrix, got, tt.want)
		}
	}
}
