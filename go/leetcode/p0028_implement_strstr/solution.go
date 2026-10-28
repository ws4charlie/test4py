// https://leetcode.com/problems/implement-strstr/
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
