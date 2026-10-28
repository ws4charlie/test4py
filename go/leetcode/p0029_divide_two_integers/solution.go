// Package p0029 solves LeetCode 29. Divide Two Integers (Medium).
// https://leetcode.com/problems/divide-two-integers/
//
// Divide dividend by divisor without using *, / or %, and return the
// quotient truncated toward zero (8.345 -> 8, -2.7335 -> -2). Results are
// limited to the signed 32-bit range [-2^31, 2^31 - 1]: anything larger is
// returned as 2^31 - 1, anything smaller as -2^31.
//
// Example 1:
//
//	Input:  dividend = 10, divisor = 3
//	Output: 3
//	Explanation: 10 / 3 = 3.33333..., truncated to 3
//
// Example 2:
//
//	Input:  dividend = 7, divisor = -3
//	Output: -2
//	Explanation: 7 / -3 = -2.33333..., truncated to -2
//
// Constraints:
//
//	-2^31 <= dividend, divisor <= 2^31 - 1
//	divisor != 0
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
