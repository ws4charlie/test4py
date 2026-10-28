// Package p0070 solves LeetCode 70. Climbing Stairs (Easy).
// https://leetcode.com/problems/climbing-stairs/
//
// A staircase has n steps and you climb 1 or 2 steps at a time. Return the
// number of distinct ways to reach the top.
//
// Example 1:
//
//	Input:  n = 2
//	Output: 2
//	Explanation: 1 + 1, or 2
//
// Example 2:
//
//	Input:  n = 3
//	Output: 3
//	Explanation: 1 + 1 + 1, 1 + 2, or 2 + 1
//
// Constraints:
//
//	1 <= n <= 45
package p0070

// ClimbStairs: f(n) = f(n-1) + f(n-2), since the last move is 1 or 2
// stairs. Only the last two values are kept.
func ClimbStairs(n int) int {
	a, b := 1, 1 // f(0), f(1)
	for i := 1; i < n; i++ {
		a, b = b, a+b
	}
	return b
}
