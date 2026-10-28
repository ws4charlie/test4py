// Package p0014 solves LeetCode 14. Longest Common Prefix (Easy).
// https://leetcode.com/problems/longest-common-prefix/
//
// Return the longest string that every string in strs starts with, or ""
// if there is none.
//
// Example 1:
//
//	Input:  strs = ["flower","flow","flight"]
//	Output: "fl"
//
// Example 2:
//
//	Input:  strs = ["dog","racecar","car"]
//	Output: ""
//
// Constraints:
//
//	1 <= strs.length <= 200
//	0 <= strs[i].length <= 200
//	strs[i] consists of lowercase English letters (if non-empty).
package p0014

// LongestCommonPrefix compares column by column against the first string.
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		for _, s := range strs[1:] {
			if i == len(s) || s[i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
