// Package p0001 solves LeetCode 1. Two Sum (Easy).
// https://leetcode.com/problems/two-sum/
//
// Given an integer array nums and an integer target, return the indices of
// the two numbers whose sum is target. There is exactly one answer, and the
// same element can't be used twice. The indices can be in any order.
//
// Example 1:
//
//	Input:  nums = [2,7,11,15], target = 9
//	Output: [0,1]
//	Explanation: nums[0] + nums[1] = 2 + 7 = 9
//
// Example 2:
//
//	Input:  nums = [3,2,4], target = 6
//	Output: [1,2]
//
// Example 3:
//
//	Input:  nums = [3,3], target = 6
//	Output: [0,1]
//
// Constraints:
//
//	2 <= nums.length <= 10^4
//	-10^9 <= nums[i] <= 10^9
//	-10^9 <= target <= 10^9
//	Exactly one valid answer exists.
//
// Follow-up: can you do it in less than O(n^2) time?
package p0001

// TwoSum returns the indices of the two numbers that add up to target.
func TwoSum(nums []int, target int) []int {
	seen := make(map[int]int) // value -> index
	for i, v := range nums {
		if j, ok := seen[target-v]; ok {
			return []int{j, i}
		}
		seen[v] = i
	}
	return nil
}
