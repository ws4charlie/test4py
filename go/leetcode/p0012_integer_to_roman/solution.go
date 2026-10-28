// https://leetcode.com/problems/integer-to-roman/
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
