// https://leetcode.com/problems/first-bad-version/
package p0278

// isBadVersion is provided by LeetCode; tests set it.
var isBadVersion func(version int) bool

// FirstBadVersion is a lower-bound search for the first bad version.
// mid = lo + (hi-lo)/2 avoids overflow for n near MaxInt32.
func FirstBadVersion(n int) int {
	lo, hi := 1, n
	for lo < hi {
		mid := lo + (hi-lo)/2
		if isBadVersion(mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}
