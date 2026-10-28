// https://leetcode.com/problems/zigzag-conversion/
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
