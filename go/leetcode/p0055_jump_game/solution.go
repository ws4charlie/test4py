// Package p0055 solves LeetCode 55. Jump Game (Medium).
// https://leetcode.com/problems/jump-game/
//
// You start at index 0. nums[i] is the longest jump you can make from index
// i (any length from 0 to nums[i]). Return true if you can reach the last
// index.
//
// Example 1:
//
//	Input:  nums = [2,3,1,1,4]
//	Output: true
//	Explanation: jump 1 step to index 1, then 3 steps to the last index.
//
// Example 2:
//
//	Input:  nums = [3,2,1,0,4]
//	Output: false
//	Explanation: every path lands on index 3, whose jump length is 0.
//
// Constraints:
//
//	1 <= nums.length <= 10^4
//	0 <= nums[i] <= 10^5
package p0055

// CanJump is greedy: track the farthest reachable index; if i is ever
// beyond it, i (and the end) can't be reached.
func CanJump(nums []int) bool {
	farthest := 0
	for i, v := range nums {
		if i > farthest {
			return false
		}
		farthest = max(farthest, i+v)
	}
	return true
}

// CanJumpDP is bottom-up: ok[i] is true if some j in (i, i+nums[i]] is ok.
func CanJumpDP(nums []int) bool {
	n := len(nums)
	ok := make([]bool, n)
	ok[n-1] = true
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j <= min(n-1, i+nums[i]); j++ {
			if ok[j] {
				ok[i] = true
				break
			}
		}
	}
	return ok[0]
}
