// Package p0038 solves LeetCode 38. Count and Say (Medium).
// https://leetcode.com/problems/count-and-say/
//
// The count-and-say sequence starts with countAndSay(1) = "1". Each later
// term describes the previous one: read it left to right and replace every
// run of the same digit with the run length followed by the digit.
//
//	"3322251" -> two 3s, three 2s, one 5, one 1 -> "23321511"
//
// Return the n-th term.
//
// Example 1:
//
//	Input:  n = 4
//	Output: "1211"
//	Explanation:
//	    countAndSay(1) = "1"
//	    countAndSay(2) = "11"     (one 1)
//	    countAndSay(3) = "21"     (two 1s)
//	    countAndSay(4) = "1211"   (one 2, one 1)
//
// Example 2:
//
//	Input:  n = 1
//	Output: "1"
//
// Constraints:
//
//	1 <= n <= 30
//
// Follow-up: can you solve it iteratively?
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
