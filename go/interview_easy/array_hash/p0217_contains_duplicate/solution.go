// Package p0217 solves LeetCode 217. Contains Duplicate (Easy).
// https://leetcode.com/problems/contains-duplicate/
//
// Return true if any value appears at least twice in nums, false if every
// value is distinct.
//
// Example 1:
//
//	Input:  nums = [1,2,3,1]
//	Output: true
//	Explanation: 1 appears at index 0 and 3.
//
// Example 2:
//
//	Input:  nums = [1,2,3,4]
//	Output: false
//
// Example 3:
//
//	Input:  nums = [1,1,1,3,3,4,3,2,4,2]
//	Output: true
//
// Constraints:
//
//	1 <= nums.length <= 10^5
//	-10^9 <= nums[i] <= 10^9
package p0217

// ContainsDuplicate uses a set (map to empty struct).
func ContainsDuplicate(nums []int) bool {
	seen := make(map[int]struct{}, len(nums))
	for _, v := range nums {
		if _, ok := seen[v]; ok {
			return true
		}
		seen[v] = struct{}{}
	}
	return false
}
