// Package p0033 solves LeetCode 33. Search in Rotated Sorted Array (Medium).
// https://leetcode.com/problems/search-in-rotated-sorted-array/
//
// nums was sorted in ascending order with distinct values, then possibly
// rotated left by an unknown k (1 <= k < n):
//
//	sorted:   0 1 2 4 5 6 7
//	rotated:  4 5 6 7 0 1 2      (k = 3)
//
// Return the index of target in nums, or -1 if it isn't there. The solution
// must run in O(log n) time.
//
// Example 1:
//
//	Input:  nums = [4,5,6,7,0,1,2], target = 0
//	Output: 4
//
// Example 2:
//
//	Input:  nums = [4,5,6,7,0,1,2], target = 3
//	Output: -1
//
// Example 3:
//
//	Input:  nums = [1], target = 0
//	Output: -1
//
// Constraints:
//
//	1 <= nums.length <= 5000
//	-10^4 <= nums[i], target <= 10^4
//	All values of nums are distinct.
//	nums is an ascending array that may have been rotated.
package p0033

// Search is a binary search where at least one half of [lo, hi] is always
// sorted; check whether target lies in that half to pick a side.
func Search(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		switch {
		case nums[mid] == target:
			return mid
		case nums[lo] <= nums[mid]: // left half sorted
			if nums[lo] <= target && target < nums[mid] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		default: // right half sorted
			if nums[mid] < target && target <= nums[hi] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}
	return -1
}
