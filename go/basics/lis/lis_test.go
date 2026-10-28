package lis

import (
	"math/rand/v2"
	"testing"
)

func TestLIS(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{[]int{0, 1, 0, 3, 2, 3}, 4},
		{[]int{7, 7, 7, 7}, 1},
		{[]int{1, 5, 3, 4, 6, 9, 7, 8}, 6},
		{[]int{5, 4, 3, 2, 1}, 1},
		{[]int{3, 1, 2}, 2}, // the best run doesn't start at index 0
		{nil, 0},
	}
	impls := map[string]func([]int) int{"DP": DP, "Patience": Patience}
	for name, f := range impls {
		for _, tt := range tests {
			if got := f(tt.nums); got != tt.want {
				t.Errorf("%s(%v) = %d, want %d", name, tt.nums, got, tt.want)
			}
		}
	}
}

func TestDPMatchesPatience(t *testing.T) {
	r := rand.New(rand.NewPCG(5, 6))
	for n := 0; n < 200; n++ {
		nums := make([]int, r.IntN(40))
		for i := range nums {
			nums[i] = r.IntN(20)
		}
		if a, b := DP(nums), Patience(nums); a != b {
			t.Fatalf("DP=%d Patience=%d for %v", a, b, nums)
		}
	}
}
