// Package p0026 solves LeetCode 26. Remove Duplicates from Sorted Array (Easy).
// https://leetcode.com/problems/remove-duplicates-from-sorted-array/
//
// nums is sorted in non-decreasing order. Remove the duplicates in place so
// each value appears once, keeping the original order, and return k, the
// number of unique values. The first k elements of nums must hold those
// values; whatever comes after them doesn't matter.
//
// Example 1:
//
//	Input:  nums = [1,1,2]
//	Output: 2, nums = [1,2,_]
//
// Example 2:
//
//	Input:  nums = [0,0,1,1,1,2,2,3,3,4]
//	Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]
//
// Constraints:
//
//	1 <= nums.length <= 3 * 10^4
//	-100 <= nums[i] <= 100
//	nums is sorted in non-decreasing order.
package p0026

// RemoveDuplicates keeps nums[:k] unique in place and returns k.
func RemoveDuplicates(nums []int) int {
	k := 0
	for _, v := range nums {
		if k == 0 || nums[k-1] != v {
			nums[k] = v
			k++
		}
	}
	return k
}
