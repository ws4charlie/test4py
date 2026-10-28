// Package p0283 solves LeetCode 283. Move Zeroes (Easy).
// https://leetcode.com/problems/move-zeroes/
//
// Move every 0 in nums to the end, in place, keeping the other elements in
// their original order. Don't copy the array.
//
// Example 1:
//
//	Input:  nums = [0,1,0,3,12]
//	Output: [1,3,12,0,0]
//
// Example 2:
//
//	Input:  nums = [0]
//	Output: [0]
//
// Constraints:
//
//	1 <= nums.length <= 10^4
//	-2^31 <= nums[i] <= 2^31 - 1
//
// Follow-up: can you minimize the total number of operations?
package p0283

// MoveZeroes swaps each non-zero into slot k, keeping their order.
func MoveZeroes(nums []int) {
	k := 0
	for i, v := range nums {
		if v != 0 {
			nums[k], nums[i] = nums[i], nums[k]
			k++
		}
	}
}
