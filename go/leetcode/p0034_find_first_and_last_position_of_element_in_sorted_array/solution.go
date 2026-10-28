// Package p0034 solves LeetCode 34. Find First and Last Position of Element in Sorted Array (Medium).
// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
//
// nums is sorted in non-decreasing order. Return [first, last], the first
// and last index of target, or [-1, -1] if target isn't there. The solution
// must run in O(log n) time.
//
// Example 1:
//
//	Input:  nums = [5,7,7,8,8,10], target = 8
//	Output: [3,4]
//
// Example 2:
//
//	Input:  nums = [5,7,7,8,8,10], target = 6
//	Output: [-1,-1]
//
// Example 3:
//
//	Input:  nums = [], target = 0
//	Output: [-1,-1]
//
// Constraints:
//
//	0 <= nums.length <= 10^5
//	-10^9 <= nums[i], target <= 10^9
//	nums is sorted in non-decreasing order.
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
