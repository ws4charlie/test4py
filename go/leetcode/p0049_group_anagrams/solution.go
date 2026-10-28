// Package p0049 solves LeetCode 49. Group Anagrams (Medium).
// https://leetcode.com/problems/group-anagrams/
//
// Group the strings that are anagrams of each other (same letters, any
// order). Return the groups in any order.
//
// Example 1:
//
//	Input:  strs = ["eat","tea","tan","ate","nat","bat"]
//	Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
//
// Example 2:
//
//	Input:  strs = [""]
//	Output: [[""]]
//
// Example 3:
//
//	Input:  strs = ["a"]
//	Output: [["a"]]
//
// Constraints:
//
//	1 <= strs.length <= 10^4
//	0 <= strs[i].length <= 100
//	strs[i] consists of lowercase English letters.
package p0049

// GroupAnagrams keys each word by its letter counts (a comparable array,
// O(k) per word instead of sorting the word). Groups come out in the order
// their first word appears.
func GroupAnagrams(strs []string) [][]string {
	index := make(map[[26]byte]int)
	var res [][]string
	for _, s := range strs {
		var key [26]byte
		for i := 0; i < len(s); i++ {
			key[s[i]-'a']++
		}
		if i, ok := index[key]; ok {
			res[i] = append(res[i], s)
		} else {
			index[key] = len(res)
			res = append(res, []string{s})
		}
	}
	return res
}
