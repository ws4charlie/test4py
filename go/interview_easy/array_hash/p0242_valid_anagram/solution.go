// https://leetcode.com/problems/valid-anagram/
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
