// Package p0011 solves LeetCode 11. Container With Most Water (Medium).
// https://leetcode.com/problems/container-with-most-water/
//
// height[i] is a vertical line at x = i. Pick two lines that, together with
// the x-axis, hold the most water, and return that amount. The container
// can't be tilted, so the shorter line sets the water level.
//
// Example 1:
//
//	Input:  height = [1,8,6,2,5,4,8,3,7]
//	Output: 49
//	Explanation: lines 1 and 8 hold min(8, 7) * (8 - 1) = 49 units (~).
//
//	8 |   #         #
//	7 |   # ~ ~ ~ ~ # ~ #
//	6 |   # # ~ ~ ~ # ~ #
//	5 |   # # ~ # ~ # ~ #
//	4 |   # # ~ # # # ~ #
//	3 |   # # ~ # # # # #
//	2 |   # # # # # # # #
//	1 | # # # # # # # # #
//	  +------------------
//	    0 1 2 3 4 5 6 7 8
//
// Example 2:
//
//	Input:  height = [1,1]
//	Output: 1
//
// Constraints:
//
//	2 <= n <= 10^5
//	0 <= height[i] <= 10^4
package p0011

// MaxArea moves the shorter side inward: the area limited by it can't grow.
func MaxArea(height []int) int {
	l, r, res := 0, len(height)-1, 0
	for l < r {
		res = max(res, (r-l)*min(height[l], height[r]))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return res
}
