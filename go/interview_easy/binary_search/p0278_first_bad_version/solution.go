// Package p0278 solves LeetCode 278. First Bad Version (Easy).
// https://leetcode.com/problems/first-bad-version/
//
// Versions 1..n of a product were released in order. One of them is bad,
// and every version after it is bad too. Use the API isBadVersion(version)
// to find the first bad version, calling it as few times as possible.
//
//	version:   1     2     3     4     5
//	bad?:      no    no    no    yes   yes
//	                             ^ first bad
//
// Example 1:
//
//	Input:  n = 5, bad = 4
//	Output: 4
//	Explanation: isBadVersion(3) -> false
//	             isBadVersion(5) -> true
//	             isBadVersion(4) -> true, so 4 is the first bad version.
//
// Example 2:
//
//	Input:  n = 1, bad = 1
//	Output: 1
//
// Constraints:
//
//	1 <= bad <= n <= 2^31 - 1
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
