// https://leetcode.com/problems/climbing-stairs/
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
