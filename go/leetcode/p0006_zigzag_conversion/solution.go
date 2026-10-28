// Package p0006 solves LeetCode 6. Zigzag Conversion (Medium).
// https://leetcode.com/problems/zigzag-conversion/
//
// Write the string s in a zigzag on numRows rows: go down the rows, then
// diagonally back up, and repeat. Then read the rows from top to bottom and
// return the result.
//
// Example 1:
//
//	Input:  s = "PAYPALISHIRING", numRows = 3
//	Output: "PAHNAPLSIIGYIR"
//
//	    P   A   H   N
//	    A P L S I I G
//	    Y   I   R
//
// Example 2:
//
//	Input:  s = "PAYPALISHIRING", numRows = 4
//	Output: "PINALSIGYAHRPI"
//
//	    P     I     N
//	    A   L S   I G
//	    Y A   H R
//	    P     I
//
// Example 3:
//
//	Input:  s = "A", numRows = 1
//	Output: "A"
//
// Constraints:
//
//	1 <= s.length <= 1000
//	s consists of English letters, ',' and '.'.
//	1 <= numRows <= 1000
package p0006

import "strings"

// Convert writes each char to its row, bouncing between the first and
// last row, then joins the rows.
func Convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}
	rows := make([]strings.Builder, numRows)
	row, step := 0, 1
	for i := 0; i < len(s); i++ {
		rows[row].WriteByte(s[i])
		if row == 0 {
			step = 1
		} else if row == numRows-1 {
			step = -1
		}
		row += step
	}
	var sb strings.Builder
	sb.Grow(len(s))
	for i := range rows {
		sb.WriteString(rows[i].String())
	}
	return sb.String()
}
