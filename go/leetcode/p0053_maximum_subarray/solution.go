// Package p0053 solves LeetCode 53. Maximum Subarray (Medium).
// https://leetcode.com/problems/maximum-subarray/
//
// Return the largest sum of any non-empty contiguous subarray of nums.
//
// Example 1:
//
//	Input:  nums = [-2,1,-3,4,-1,2,1,-5,4]
//	Output: 6
//	Explanation: the subarray [4,-1,2,1] sums to 6.
//
// Example 2:
//
//	Input:  nums = [1]
//	Output: 1
//
// Example 3:
//
//	Input:  nums = [5,4,-1,7,8]
//	Output: 23
//	Explanation: the whole array.
//
// Constraints:
//
//	1 <= nums.length <= 10^5
//	-10^4 <= nums[i] <= 10^4
//
// Follow-up: after the O(n) solution, try divide and conquer.
package p0053

// MaxSubArray is Kadane's algorithm: cur is the best sum of a subarray
// ending at i, so either extend the previous one or start fresh at i.
func MaxSubArray(nums []int) int {
	cur, best := nums[0], nums[0]
	for _, v := range nums[1:] {
		cur = max(v, cur+v)
		best = max(best, cur)
	}
	return best
}
