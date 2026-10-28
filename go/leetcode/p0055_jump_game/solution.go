// https://leetcode.com/problems/jump-game/
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
