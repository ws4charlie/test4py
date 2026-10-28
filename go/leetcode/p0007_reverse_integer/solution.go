// https://leetcode.com/problems/reverse-integer/
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
