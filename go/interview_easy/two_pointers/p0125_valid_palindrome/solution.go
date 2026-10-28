// Package p0125 solves LeetCode 125. Valid Palindrome (Easy).
// https://leetcode.com/problems/valid-palindrome/
//
// Lowercase s and drop every character that isn't a letter or digit. Return
// true if what's left reads the same forwards and backwards.
//
// Example 1:
//
//	Input:  s = "A man, a plan, a canal: Panama"
//	Output: true
//	Explanation: "amanaplanacanalpanama" is a palindrome.
//
// Example 2:
//
//	Input:  s = "race a car"
//	Output: false
//	Explanation: "raceacar" is not a palindrome.
//
// Example 3:
//
//	Input:  s = " "
//	Output: true
//	Explanation: nothing is left, and the empty string is a palindrome.
//
// Constraints:
//
//	1 <= s.length <= 2 * 10^5
//	s consists of printable ASCII characters.
package p0125

// IsPalindrome skips non-alphanumeric bytes from both ends and compares
// the rest case-insensitively (the input is ASCII).
func IsPalindrome(s string) bool {
	for l, r := 0, len(s)-1; l < r; {
		switch {
		case !isAlnum(s[l]):
			l++
		case !isAlnum(s[r]):
			r--
		case lower(s[l]) != lower(s[r]):
			return false
		default:
			l, r = l+1, r-1
		}
	}
	return true
}

func isAlnum(c byte) bool {
	return 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || '0' <= c && c <= '9'
}

func lower(c byte) byte {
	if 'A' <= c && c <= 'Z' {
		return c + 'a' - 'A'
	}
	return c
}
