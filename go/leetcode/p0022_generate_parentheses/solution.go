// Package p0022 solves LeetCode 22. Generate Parentheses (Medium).
// https://leetcode.com/problems/generate-parentheses/
//
// Given n pairs of parentheses, return every well-formed combination, in
// any order.
//
// Example 1:
//
//	Input:  n = 3
//	Output: ["((()))","(()())","(())()","()(())","()()()"]
//
// Example 2:
//
//	Input:  n = 1
//	Output: ["()"]
//
// Constraints:
//
//	1 <= n <= 8
package p0022

// GenerateParenthesis adds '(' while open < n and ')' while close < open,
// so every prefix stays valid.
func GenerateParenthesis(n int) []string {
	var res []string
	comb := make([]byte, 0, 2*n)
	var backtrack func(open, close int)
	backtrack = func(open, close int) {
		if len(comb) == 2*n {
			res = append(res, string(comb))
			return
		}
		if open < n {
			comb = append(comb, '(')
			backtrack(open+1, close)
			comb = comb[:len(comb)-1]
		}
		if close < open {
			comb = append(comb, ')')
			backtrack(open, close+1)
			comb = comb[:len(comb)-1]
		}
	}
	backtrack(0, 0)
	return res
}
