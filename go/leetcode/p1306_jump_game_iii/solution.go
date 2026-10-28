// https://leetcode.com/problems/jump-game-iii/
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
