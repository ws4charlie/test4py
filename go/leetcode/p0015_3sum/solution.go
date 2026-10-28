// https://leetcode.com/problems/3sum/
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
