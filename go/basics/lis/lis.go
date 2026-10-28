// Package lis computes the length of the longest strictly increasing
// subsequence (LeetCode 300), ported from
// python/basics/longest_increasing_subsequence.py.
package lis

import "sort"

// DP: f[i] is the LIS length ending at i, f[i] = 1 + max f[j] for j < i
// with nums[j] < nums[i]. The answer is the max over all i, not just the
// last one. O(n^2).
func DP(nums []int) int {
	f := make([]int, len(nums))
	best := 0
	for i := range nums {
		f[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				f[i] = max(f[i], f[j]+1)
			}
		}
		best = max(best, f[i])
	}
	return best
}

// Patience is O(n log n). tails[k] is the smallest possible tail of an
// increasing subsequence of length k+1; each number replaces the first
// tail >= it, or extends tails if it is bigger than all of them.
func Patience(nums []int) int {
	var tails []int
	for _, v := range nums {
		i := sort.SearchInts(tails, v)
		if i == len(tails) {
			tails = append(tails, v)
		} else {
			tails[i] = v
		}
	}
	return len(tails)
}
