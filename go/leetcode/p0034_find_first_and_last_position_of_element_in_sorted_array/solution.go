// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
package p0034

// SearchRange runs lowerBound twice: the first index >= target, and the
// first index >= target+1 (one past the last target). O(log n) even when
// target repeats many times.
func SearchRange(nums []int, target int) []int {
	first := lowerBound(nums, target)
	if first == len(nums) || nums[first] != target {
		return []int{-1, -1}
	}
	return []int{first, lowerBound(nums, target+1) - 1}
}

// lowerBound returns the first index i with nums[i] >= target, or len(nums).
func lowerBound(nums []int, target int) int {
	lo, hi := 0, len(nums) // search in [lo, hi)
	for lo < hi {
		mid := lo + (hi-lo)/2
		if nums[mid] < target {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}
