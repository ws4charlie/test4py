// https://leetcode.com/problems/group-anagrams/
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
