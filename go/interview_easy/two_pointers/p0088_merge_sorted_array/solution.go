// https://leetcode.com/problems/merge-sorted-array/
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
