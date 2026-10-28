// Package p0005 solves LeetCode 5. Longest Palindromic Substring (Medium).
// https://leetcode.com/problems/longest-palindromic-substring/
//
// Given a string s, return its longest substring that reads the same
// forwards and backwards.
//
// Example 1:
//
//	Input:  s = "babad"
//	Output: "bab"
//	Explanation: "aba" is also a valid answer.
//
// Example 2:
//
//	Input:  s = "cbbd"
//	Output: "bb"
//
// Constraints:
//
//	1 <= s.length <= 1000
//	s consists of digits and English letters only.
package p0005

// LongestPalindrome expands around every center: one char (odd length)
// and between two chars (even length).
func LongestPalindrome(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		for _, sub := range []string{expand(s, i, i), expand(s, i, i+1)} {
			if len(sub) > len(res) {
				res = sub
			}
		}
	}
	return res
}

func expand(s string, i, j int) string {
	for i >= 0 && j < len(s) && s[i] == s[j] {
		i, j = i-1, j+1
	}
	return s[i+1 : j]
}
