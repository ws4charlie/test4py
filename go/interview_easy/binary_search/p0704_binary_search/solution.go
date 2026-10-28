// Package p0704 solves LeetCode 704. Binary Search (Easy).
// https://leetcode.com/problems/binary-search/
//
// nums is sorted in ascending order. Return the index of target, or -1 if
// it isn't there. The solution must run in O(log n) time.
//
// Example 1:
//
//	Input:  nums = [-1,0,3,5,9,12], target = 9
//	Output: 4
//
// Example 2:
//
//	Input:  nums = [-1,0,3,5,9,12], target = 2
//	Output: -1
//
// Constraints:
//
//	1 <= nums.length <= 10^4
//	-10^4 < nums[i], target < 10^4
//	All values of nums are unique.
//	nums is sorted in ascending order.
package p0704

// Search is binary search on the closed interval [lo, hi].
func Search(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		switch {
		case nums[mid] == target:
			return mid
		case nums[mid] < target:
			lo = mid + 1
		default:
			hi = mid - 1
		}
	}
	return -1
}
