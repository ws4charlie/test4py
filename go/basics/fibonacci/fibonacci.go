// Package fibonacci shows the same recurrence three ways, ported from
// python/basics/fibonacci.py: F(0)=0, F(1)=1, F(n)=F(n-1)+F(n-2).
package fibonacci

// Recursive is the direct recurrence: O(2^n), recomputes subproblems.
func Recursive(n int) int {
	if n < 2 {
		return n
	}
	return Recursive(n-1) + Recursive(n-2)
}

// Memo is top-down DP: each F(i) is computed once. O(n) time and space.
func Memo(n int) int {
	memo := map[int]int{0: 0, 1: 1}
	var f func(i int) int
	f = func(i int) int {
		if v, ok := memo[i]; ok {
			return v
		}
		memo[i] = f(i-1) + f(i-2)
		return memo[i]
	}
	return f(n)
}

// Iterative is bottom-up DP keeping only the last two values. O(n) time,
// O(1) space.
func Iterative(n int) int {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return a
}
