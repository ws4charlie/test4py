// Package p0013 solves LeetCode 13. Roman to Integer (Easy).
// https://leetcode.com/problems/roman-to-integer/
//
// Convert a Roman numeral to an integer.
//
//	Symbol   Value        Subtractive forms
//	I        1            IV   4
//	V        5            IX   9
//	X        10           XL   40
//	L        50           XC   90
//	C        100          CD   400
//	D        500          CM   900
//	M        1000
//
// Symbols are normally written largest to smallest and added up. When a
// smaller symbol comes before a larger one (the subtractive forms above),
// it is subtracted instead.
//
// Example 1:
//
//	Input:  s = "III"
//	Output: 3
//
// Example 2:
//
//	Input:  s = "LVIII"
//	Output: 58
//	Explanation: L = 50, V = 5, III = 3
//
// Example 3:
//
//	Input:  s = "MCMXCIV"
//	Output: 1994
//	Explanation: M = 1000, CM = 900, XC = 90, IV = 4
//
// Constraints:
//
//	1 <= s.length <= 15
//	s contains only I, V, X, L, C, D, M.
//	s is a valid Roman numeral in the range [1, 3999].
package p0013

var value = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

// RomanToInt subtracts a symbol when a larger one follows it (IV, CM, ...).
func RomanToInt(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		v := value[s[i]]
		if i+1 < len(s) && v < value[s[i+1]] {
			res -= v
		} else {
			res += v
		}
	}
	return res
}
