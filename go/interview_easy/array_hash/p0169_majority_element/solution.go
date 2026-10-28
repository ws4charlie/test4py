// Package p0169 solves LeetCode 169. Majority Element (Easy).
// https://leetcode.com/problems/majority-element/
//
// Return the element that appears more than n / 2 times in an array of
// size n. It always exists.
//
// Example 1:
//
//	Input:  nums = [3,2,3]
//	Output: 3
//
// Example 2:
//
//	Input:  nums = [2,2,1,1,1,2,2]
//	Output: 2
//
// Constraints:
//
//	1 <= nums.length <= 5 * 10^4
//	-10^9 <= nums[i] <= 10^9
//	The majority element always exists.
//
// Follow-up: can you do it in O(n) time and O(1) space?
package p0169

// MajorityElement is the Boyer-Moore vote: the majority survives pairing
// off against every other value. O(n) time, O(1) space.
func MajorityElement(nums []int) int {
	cand, count := 0, 0
	for _, v := range nums {
		if count == 0 {
			cand = v
		}
		if v == cand {
			count++
		} else {
			count--
		}
	}
	return cand
}
