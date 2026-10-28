// https://leetcode.com/problems/longest-substring-without-repeating-characters/
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
