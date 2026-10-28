// Package p0136 solves LeetCode 136. Single Number (Easy).
// https://leetcode.com/problems/single-number/
//
// Every element of nums appears exactly twice, except one that appears
// once. Return that one. Use O(n) time and O(1) extra space.
//
// Example 1:
//
//	Input:  nums = [2,2,1]
//	Output: 1
//
// Example 2:
//
//	Input:  nums = [4,1,2,1,2]
//	Output: 4
//
// Example 3:
//
//	Input:  nums = [1]
//	Output: 1
//
// Constraints:
//
//	1 <= nums.length <= 3 * 10^4
//	-3 * 10^4 <= nums[i] <= 3 * 10^4
//	Every element appears twice except for exactly one.
package p0136

// SingleNumber XORs everything: x^x == 0 and x^0 == x, so pairs cancel.
func SingleNumber(nums []int) int {
	res := 0
	for _, v := range nums {
		res ^= v
	}
	return res
}
