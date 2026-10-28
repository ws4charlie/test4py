// Package p1306 solves LeetCode 1306. Jump Game III (Medium).
// https://leetcode.com/problems/jump-game-iii/
//
// arr is an array of non-negative integers and you start at index start.
// From index i you can jump to i + arr[i] or i - arr[i], but never outside
// the array. Return true if you can reach any index whose value is 0.
//
// Example 1:
//
//	Input:  arr = [4,2,3,0,3,1,2], start = 5
//	Output: true
//	Explanation: index 5 -> 4 -> 1 -> 3 (value 0)
//
// Example 2:
//
//	Input:  arr = [4,2,3,0,3,1,2], start = 0
//	Output: true
//	Explanation: index 0 -> 4 -> 1 -> 3
//
// Example 3:
//
//	Input:  arr = [3,0,2,1,2], start = 2
//	Output: false
//	Explanation: from index 2 you can only ever reach indices 0, 3 and
//	             4; the only 0 is at index 1.
//
// Constraints:
//
//	1 <= arr.length <= 5 * 10^4
//	0 <= arr[i] < arr.length
//	0 <= start < arr.length
package p1306

// CanReach runs BFS from start; each index i has edges to i±arr[i].
func CanReach(arr []int, start int) bool {
	seen := make([]bool, len(arr))
	seen[start] = true
	queue := []int{start}
	for len(queue) > 0 {
		i := queue[0]
		queue = queue[1:]
		if arr[i] == 0 {
			return true
		}
		for _, j := range []int{i - arr[i], i + arr[i]} {
			if j >= 0 && j < len(arr) && !seen[j] {
				seen[j] = true
				queue = append(queue, j)
			}
		}
	}
	return false
}

// CanReachDFS is the recursive version; seen also breaks cycles.
func CanReachDFS(arr []int, start int) bool {
	seen := make([]bool, len(arr))
	var dfs func(i int) bool
	dfs = func(i int) bool {
		if i < 0 || i >= len(arr) || seen[i] {
			return false
		}
		if arr[i] == 0 {
			return true
		}
		seen[i] = true
		return dfs(i-arr[i]) || dfs(i+arr[i])
	}
	return dfs(start)
}
