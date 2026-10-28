// https://leetcode.com/problems/valid-parentheses/
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
