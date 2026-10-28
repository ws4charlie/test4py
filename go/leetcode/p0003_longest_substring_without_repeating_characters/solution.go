// Package p0003 solves LeetCode 3. Longest Substring Without Repeating Characters (Medium).
// https://leetcode.com/problems/longest-substring-without-repeating-characters/
//
// Given a string s, return the length of the longest substring (a
// contiguous run of characters) in which no character appears twice.
//
// Example 1:
//
//	Input:  s = "abcabcbb"
//	Output: 3
//	Explanation: "abc" (also "bca", "cab")
//
// Example 2:
//
//	Input:  s = "bbbbb"
//	Output: 1
//	Explanation: "b"
//
// Example 3:
//
//	Input:  s = "pwwkew"
//	Output: 3
//	Explanation: "wke". "pwke" doesn't count: it is a subsequence, not a
//	             substring.
//
// Constraints:
//
//	0 <= s.length <= 5 * 10^4
//	s consists of English letters, digits, symbols and spaces.
package p0003

// LengthOfLongestSubstring uses a sliding window [left, i]; when s[i] was
// already seen inside the window, left jumps past its previous position.
func LengthOfLongestSubstring(s string) int {
	last := make(map[byte]int) // char -> last index
	left, res := 0, 0
	for i := 0; i < len(s); i++ {
		if j, ok := last[s[i]]; ok && j >= left {
			left = j + 1
		}
		last[s[i]] = i
		res = max(res, i-left+1)
	}
	return res
}
