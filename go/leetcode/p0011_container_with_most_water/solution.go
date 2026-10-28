// https://leetcode.com/problems/container-with-most-water/
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
