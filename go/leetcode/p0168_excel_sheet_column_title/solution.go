// Package p0168 solves LeetCode 168. Excel Sheet Column Title (Easy).
// https://leetcode.com/problems/excel-sheet-column-title/
//
// Given a column number, return its title as it appears in an Excel sheet:
//
//	1  -> A
//	2  -> B
//	...
//	26 -> Z
//	27 -> AA
//	28 -> AB
//	...
//
// Example 1:
//
//	Input:  columnNumber = 1
//	Output: "A"
//
// Example 2:
//
//	Input:  columnNumber = 28
//	Output: "AB"
//
// Example 3:
//
//	Input:  columnNumber = 701
//	Output: "ZY"
//
// Constraints:
//
//	1 <= columnNumber <= 2^31 - 1
package p0168

// ConvertToTitle is base 26 with digits 1..26 instead of 0..25, so shift
// down by one before each % and /.
func ConvertToTitle(n int) string {
	var buf []byte
	for n > 0 {
		n--
		buf = append(buf, byte('A'+n%26))
		n /= 26
	}
	for l, r := 0, len(buf)-1; l < r; l, r = l+1, r-1 {
		buf[l], buf[r] = buf[r], buf[l]
	}
	return string(buf)
}
