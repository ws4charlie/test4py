// Package p0028 solves LeetCode 28. Find the Index of the First Occurrence in a String (Easy).
// https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/
// (formerly "Implement strStr()")
//
// Return the index where needle first appears in haystack, or -1 if it
// doesn't appear.
//
// Example 1:
//
//	Input:  haystack = "sadbutsad", needle = "sad"
//	Output: 0
//	Explanation: "sad" appears at 0 and 6; the first is 0.
//
// Example 2:
//
//	Input:  haystack = "leetcode", needle = "leeto"
//	Output: -1
//
// Constraints:
//
//	1 <= haystack.length, needle.length <= 10^4
//	Both consist of lowercase English letters only.
package p0028

// StrStr tries every start position. O(n*m).
func StrStr(haystack, needle string) int {
	for i := 0; i+len(needle) <= len(haystack); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}

// StrStrKMP is Knuth-Morris-Pratt. O(n+m). lps[i] is the length of the
// longest proper prefix of needle[:i+1] that is also its suffix; on a
// mismatch, j falls back to lps[j-1] instead of restarting at 0.
func StrStrKMP(haystack, needle string) int {
	if needle == "" {
		return 0
	}
	lps := make([]int, len(needle))
	for i, j := 1, 0; i < len(needle); i++ {
		for j > 0 && needle[i] != needle[j] {
			j = lps[j-1]
		}
		if needle[i] == needle[j] {
			j++
		}
		lps[i] = j
	}
	for i, j := 0, 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = lps[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == len(needle) {
			return i - j + 1
		}
	}
	return -1
}
