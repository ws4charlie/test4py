// Package p0007 solves LeetCode 7. Reverse Integer (Medium).
// https://leetcode.com/problems/reverse-integer/
//
// Given a signed 32-bit integer x, return x with its digits reversed. If the
// result falls outside the signed 32-bit range [-2^31, 2^31 - 1], return 0.
// Assume you can't store 64-bit integers.
//
// Example 1:
//
//	Input:  x = 123
//	Output: 321
//
// Example 2:
//
//	Input:  x = -123
//	Output: -321
//
// Example 3:
//
//	Input:  x = 120
//	Output: 21
//
// Constraints:
//
//	-2^31 <= x <= 2^31 - 1
package p0007

import "math"

// Reverse pops digits off x and pushes them onto res. The problem forbids
// 64-bit storage, so overflow is checked before res*10 instead of after.
func Reverse(x int) int {
	res := 0
	for x != 0 {
		d := x % 10 // keeps x's sign in Go
		x /= 10
		if res > math.MaxInt32/10 || (res == math.MaxInt32/10 && d > 7) {
			return 0
		}
		if res < math.MinInt32/10 || (res == math.MinInt32/10 && d < -8) {
			return 0
		}
		res = res*10 + d
	}
	return res
}
