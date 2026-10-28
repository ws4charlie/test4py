// https://leetcode.com/problems/remove-duplicates-from-sorted-array/
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
