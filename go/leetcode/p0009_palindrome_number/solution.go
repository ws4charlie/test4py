// Package p0009 solves LeetCode 9. Palindrome Number (Easy).
// https://leetcode.com/problems/palindrome-number/
//
// Given an integer x, return true if it reads the same forwards and
// backwards.
//
// Example 1:
//
//	Input:  x = 121
//	Output: true
//
// Example 2:
//
//	Input:  x = -121
//	Output: false
//	Explanation: backwards it reads "121-".
//
// Example 3:
//
//	Input:  x = 10
//	Output: false
//	Explanation: backwards it reads "01".
//
// Constraints:
//
//	-2^31 <= x <= 2^31 - 1
//
// Follow-up: can you do it without converting x to a string?
package p0009

// IsPalindrome reverses only the second half of the digits, without
// converting to a string. Negatives and numbers ending in 0 (except 0)
// can't be palindromes.
func IsPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	half := 0
	for x > half {
		half = half*10 + x%10
		x /= 10
	}
	return x == half || x == half/10 // half/10 drops the middle digit
}
