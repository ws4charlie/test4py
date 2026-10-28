// https://leetcode.com/problems/first-missing-positive/
package p0041

// FirstMissingPositive is O(n) time, O(1) extra space (cyclic sort): swap
// each v in [1, n] into slot v-1, then the first slot i without i+1 gives
// the answer. Modifies nums.
func FirstMissingPositive(nums []int) int {
	n := len(nums)
	for i := range nums {
		for nums[i] >= 1 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			j := nums[i] - 1
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	for i, v := range nums {
		if v != i+1 {
			return i + 1
		}
	}
	return n + 1
}
