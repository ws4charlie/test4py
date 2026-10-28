// https://leetcode.com/problems/divide-two-integers/
package p0029

import "math"

// Divide repeatedly subtracts the largest divisor<<k that still fits,
// like long division in base 2. No *, / or %.
func Divide(dividend, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32 // the only case that overflows
	}
	neg := (dividend < 0) != (divisor < 0)
	a, b := abs(dividend), abs(divisor)
	res := 0
	for a >= b {
		shift := 0
		for a >= b<<(shift+1) {
			shift++
		}
		a -= b << shift
		res += 1 << shift
	}
	if neg {
		return -res
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
