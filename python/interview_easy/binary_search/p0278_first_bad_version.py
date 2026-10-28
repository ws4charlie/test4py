"""
278. First Bad Version  (Easy)
https://leetcode.com/problems/first-bad-version/

Versions 1..n of a product were released in order. One of them is bad,
and every version after it is bad too. Use the API isBadVersion(version)
to find the first bad version, calling it as few times as possible.

    version:   1     2     3     4     5
    bad?:      no    no    no    yes   yes
                                 ^ first bad

Example 1:
    Input:  n = 5, bad = 4
    Output: 4
    Explanation: isBadVersion(3) -> False
                 isBadVersion(5) -> True
                 isBadVersion(4) -> True, so 4 is the first bad version.

Example 2:
    Input:  n = 1, bad = 1
    Output: 1

Constraints:
    1 <= bad <= n <= 2^31 - 1
"""

# LeetCode provides isBadVersion(version); the tests below set it.
isBadVersion = None

class Solution:
    def firstBadVersion(self, n):
        # lower bound: find the first version where isBadVersion is True
        lo, hi = 1, n
        while lo < hi:
            mid = lo + (hi - lo) // 2
            if isBadVersion(mid):
                hi = mid
            else:
                lo = mid + 1
        return lo

if __name__ == "__main__":
    sln = Solution()
    for n, bad in [(5, 4), (1, 1), (10, 1), (10, 10), (2126753390, 1702766719)]:
        isBadVersion = lambda v, bad=bad: v >= bad
        assert sln.firstBadVersion(n) == bad, (n, bad)
    print("ok")
