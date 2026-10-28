// https://leetcode.com/problems/jump-game-ii/
package p0045

// Jump is a greedy BFS over index ranges: [start, end] is what this many
// jumps can reach, and farthest is where one more jump can get to.
func Jump(nums []int) int {
	jumps, end, farthest := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		farthest = max(farthest, i+nums[i])
		if i == end {
			jumps++
			end = farthest
		}
	}
	return jumps
}

// JumpMemo is the top-down DP: f(i) = 1 + min f(j) for i < j <= i+nums[i],
// with f(last) = 0. O(n * max(nums)); -1 means the end can't be reached.
func JumpMemo(nums []int) int {
	memo := make(map[int]int)
	var f func(i int) int
	f = func(i int) int {
		if i >= len(nums)-1 {
			return 0
		}
		if v, ok := memo[i]; ok {
			return v
		}
		best := -1
		for j := i + 1; j <= i+nums[i]; j++ {
			if sub := f(j); sub != -1 && (best == -1 || sub+1 < best) {
				best = sub + 1
			}
		}
		memo[i] = best
		return best
	}
	return f(0)
}
