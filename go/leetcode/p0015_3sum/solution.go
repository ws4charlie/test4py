// Package p0015 solves LeetCode 15. 3Sum (Medium).
// https://leetcode.com/problems/3sum/
//
// Return every distinct triplet [nums[i], nums[j], nums[k]] with i, j, k
// all different and nums[i] + nums[j] + nums[k] == 0. The result must not
// contain the same triplet twice; order doesn't matter.
//
// Example 1:
//
//	Input:  nums = [-1,0,1,2,-1,-4]
//	Output: [[-1,-1,2],[-1,0,1]]
//	Explanation: -1 + 0 + 1 = 0 and -1 + -1 + 2 = 0. The triplet
//	             [-1,0,1] can be formed twice but appears once.
//
// Example 2:
//
//	Input:  nums = [0,1,1]
//	Output: []
//
// Example 3:
//
//	Input:  nums = [0,0,0]
//	Output: [[0,0,0]]
//
// Constraints:
//
//	3 <= nums.length <= 3000
//	-10^5 <= nums[i] <= 10^5
package p0015

import "slices"

// ThreeSum sorts, fixes nums[i], then closes in with two pointers.
// Duplicates are skipped in place instead of collected in a set.
func ThreeSum(nums []int) [][]int {
	nums = slices.Clone(nums)
	slices.Sort(nums)

	var res [][]int
	for i := 0; i < len(nums)-2 && nums[i] <= 0; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, len(nums)-1
		for j < k {
			switch sum := nums[i] + nums[j] + nums[k]; {
			case sum < 0:
				j++
			case sum > 0:
				k--
			default:
				res = append(res, []int{nums[i], nums[j], nums[k]})
				for j++; j < k && nums[j] == nums[j-1]; j++ {
				}
				k--
			}
		}
	}
	return res
}
