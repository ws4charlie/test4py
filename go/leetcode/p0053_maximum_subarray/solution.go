// https://leetcode.com/problems/maximum-subarray/
package p0053

// MaxSubArray is Kadane's algorithm: cur is the best sum of a subarray
// ending at i, so either extend the previous one or start fresh at i.
func MaxSubArray(nums []int) int {
	cur, best := nums[0], nums[0]
	for _, v := range nums[1:] {
		cur = max(v, cur+v)
		best = max(best, cur)
	}
	return best
}
