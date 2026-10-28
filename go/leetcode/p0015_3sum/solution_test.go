package p0015

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		nums []int
		want [][]int
	}{
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{[]int{0, 1, 1}, nil},
		{[]int{0, 0, 0, 0}, [][]int{{0, 0, 0}}},
		{[]int{-2, 0, 0, 2, 2}, [][]int{{-2, 0, 2}}},
	}
	for _, tt := range tests {
		if got := ThreeSum(tt.nums); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("ThreeSum(%v) = %v, want %v", tt.nums, got, tt.want)
		}
	}
}
