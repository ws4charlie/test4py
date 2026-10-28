// Package p0242 solves LeetCode 242. Valid Anagram (Easy).
// https://leetcode.com/problems/valid-anagram/
//
// Return true if t is an anagram of s: the same letters with the same
// counts, in any order.
//
// Example 1:
//
//	Input:  s = "anagram", t = "nagaram"
//	Output: true
//
// Example 2:
//
//	Input:  s = "rat", t = "car"
//	Output: false
//
// Constraints:
//
//	1 <= s.length, t.length <= 5 * 10^4
//	s and t consist of lowercase English letters.
//
// Follow-up: what if the inputs contain Unicode characters?
package p0242

// IsAnagram counts letters up for s and down for t (lowercase a-z).
func IsAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var count [26]int
	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
		count[t[i]-'a']--
	}
	return count == [26]int{}
}
