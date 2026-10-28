// https://leetcode.com/problems/move-zeroes/
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
