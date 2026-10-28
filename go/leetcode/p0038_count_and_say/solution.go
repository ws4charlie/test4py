// https://leetcode.com/problems/count-and-say/
package p0038

import (
	"strconv"
	"strings"
)

// CountAndSay builds each term from the previous one by run-length
// encoding it ("1211" -> one 1, one 2, two 1s -> "111221").
func CountAndSay(n int) string {
	s := "1"
	for k := 1; k < n; k++ {
		var sb strings.Builder
		for i := 0; i < len(s); {
			j := i
			for j < len(s) && s[j] == s[i] {
				j++
			}
			sb.WriteString(strconv.Itoa(j - i))
			sb.WriteByte(s[i])
			i = j
		}
		s = sb.String()
	}
	return s
}
