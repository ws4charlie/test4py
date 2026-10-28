// Package p0088 solves LeetCode 88. Merge Sorted Array (Easy).
// https://leetcode.com/problems/merge-sorted-array/
//
// nums1 and nums2 are sorted in non-decreasing order and hold m and n
// elements. nums1 has length m + n: its last n slots are 0 placeholders.
// Merge nums2 into nums1 in place so nums1 ends up sorted. Return nothing.
//
// Example 1:
//
//	Input:  nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
//	Output: [1,2,2,3,5,6]
//
//	    nums1: 1 2 3 _ _ _
//	    nums2: 2 5 6
//	    ->     1 2 2 3 5 6
//
// Example 2:
//
//	Input:  nums1 = [1], m = 1, nums2 = [], n = 0
//	Output: [1]
//
// Example 3:
//
//	Input:  nums1 = [0], m = 0, nums2 = [1], n = 1
//	Output: [1]
//	Explanation: nums1 has no real elements; the 0 is only a placeholder.
//
// Constraints:
//
//	nums1.length == m + n, nums2.length == n
//	0 <= m, n <= 200, 1 <= m + n <= 200
//	-10^9 <= nums1[i], nums2[j] <= 10^9
//
// Follow-up: can you do it in O(m + n) time?
package p0088

// Merge fills nums1 from the back, so nothing in nums1 is overwritten
// before it has been used.
func Merge(nums1 []int, m int, nums2 []int, n int) {
	i, j := m-1, n-1
	for k := m + n - 1; j >= 0; k-- {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
	}
}
