// Package p0020 solves LeetCode 20. Valid Parentheses (Easy).
// https://leetcode.com/problems/valid-parentheses/
//
// s contains only the characters '(', ')', '[', ']', '{' and '}'. Return
// true if it is valid:
//   - every opening bracket is closed by a bracket of the same type,
//   - brackets close in the right order (innermost first),
//   - every closing bracket has a matching opening bracket.
//
// Example 1:
//
//	Input:  s = "()"
//	Output: true
//
// Example 2:
//
//	Input:  s = "()[]{}"
//	Output: true
//
// Example 3:
//
//	Input:  s = "(]"
//	Output: false
//
// Example 4:
//
//	Input:  s = "([])"
//	Output: true
//
// Example 5:
//
//	Input:  s = "([)]"
//	Output: false
//
// Constraints:
//
//	1 <= s.length <= 10^4
//	s consists of brackets '()[]{}' only.
package p0020

var pairs = map[byte]byte{')': '(', ']': '[', '}': '{'}

// IsValid pushes openers and pops on a matching closer.
func IsValid(s string) bool {
	var stack []byte
	for i := 0; i < len(s); i++ {
		c := s[i]
		open, isCloser := pairs[c]
		if !isCloser {
			stack = append(stack, c)
			continue
		}
		if len(stack) == 0 || stack[len(stack)-1] != open {
			return false
		}
		stack = stack[:len(stack)-1]
	}
	return len(stack) == 0
}
