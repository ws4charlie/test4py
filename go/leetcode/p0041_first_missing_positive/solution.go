// Package p0041 solves LeetCode 41. First Missing Positive (Hard).
// https://leetcode.com/problems/first-missing-positive/
//
// Given an unsorted integer array nums, return the smallest positive integer
// that is not in nums. The solution must run in O(n) time and use O(1)
// extra space.
//
// Example 1:
//
//	Input:  nums = [1,2,0]
//	Output: 3
//	Explanation: 1 and 2 are present.
//
// Example 2:
//
//	Input:  nums = [3,4,-1,1]
//	Output: 2
//	Explanation: 1 is present, 2 is missing.
//
// Example 3:
//
//	Input:  nums = [7,8,9,11,12]
//	Output: 1
//
// Constraints:
//
//	1 <= nums.length <= 10^5
//	-2^31 <= nums[i] <= 2^31 - 1
package p0041

// FirstMissingPositive is O(n) time, O(1) extra space (cyclic sort): swap
// each v in [1, n] into slot v-1, then the first slot i without i+1 gives
// the answer. Modifies nums.
func FirstMissingPositive(nums []int) int {
	n := len(nums)
	for i := range nums {
		for nums[i] >= 1 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			j := nums[i] - 1
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	for i, v := range nums {
		if v != i+1 {
			return i + 1
		}
	}
	return n + 1
}
