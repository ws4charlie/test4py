package linkedlist

import (
	"slices"
	"testing"
)

func TestRoundTrip(t *testing.T) {
	for _, vals := range [][]int{nil, {1}, {1, 2, 3}} {
		if got := ToSlice(FromSlice(vals)); !slices.Equal(got, vals) {
			t.Errorf("ToSlice(FromSlice(%v)) = %v", vals, got)
		}
	}
}
