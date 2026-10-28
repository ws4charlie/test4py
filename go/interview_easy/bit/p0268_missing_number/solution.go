// Package p0268 solves LeetCode 268. Missing Number (Easy).
// https://leetcode.com/problems/missing-number/
//
// nums holds n distinct numbers from the range [0, n], so exactly one
// number in the range is missing. Return it.
//
// Example 1:
//
//	Input:  nums = [3,0,1]
//	Output: 2
//	Explanation: n = 3, range [0,3], 2 is missing.
//
// Example 2:
//
//	Input:  nums = [0,1]
//	Output: 2
//	Explanation: n = 2, range [0,2], 2 is missing.
//
// Example 3:
//
//	Input:  nums = [9,6,4,2,3,5,7,0,1]
//	Output: 8
//
// Constraints:
//
//	n == nums.length
//	1 <= n <= 10^4
//	0 <= nums[i] <= n
//	All values of nums are unique.
//
// Follow-up: can you do it in O(1) extra space and O(n) time?
package p0268

// MissingNumber XORs every index 0..n with every value; all but the
// missing number pair up and cancel.
func MissingNumber(nums []int) int {
	res := len(nums)
	for i, v := range nums {
		res ^= i ^ v
	}
	return res
}

// MissingNumberSum subtracts the actual sum from 0+1+...+n.
func MissingNumberSum(nums []int) int {
	n := len(nums)
	sum := n * (n + 1) / 2
	for _, v := range nums {
		sum -= v
	}
	return sum
}
