// Package p0012 solves LeetCode 12. Integer to Roman (Medium).
// https://leetcode.com/problems/integer-to-roman/
//
// Convert an integer to a Roman numeral.
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
// Write the numeral from the largest place value to the smallest. A digit 4
// or 9 uses a subtractive form. Only I, X, C and M can repeat, at most 3
// times in a row; V, L and D never repeat.
//
// Example 1:
//
//	Input:  num = 3749
//	Output: "MMMDCCXLIX"
//	Explanation: 3000 = MMM, 700 = DCC, 40 = XL, 9 = IX
//
// Example 2:
//
//	Input:  num = 58
//	Output: "LVIII"
//	Explanation: 50 = L, 8 = VIII
//
// Example 3:
//
//	Input:  num = 1994
//	Output: "MCMXCIV"
//	Explanation: 1000 = M, 900 = CM, 90 = XC, 4 = IV
//
// Constraints:
//
//	1 <= num <= 3999
package p0012

import "strings"

var (
	values  = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
)

// IntToRoman greedily takes the largest value that still fits.
func IntToRoman(num int) string {
	var sb strings.Builder
	for i, v := range values {
		for num >= v {
			sb.WriteString(symbols[i])
			num -= v
		}
	}
	return sb.String()
}
