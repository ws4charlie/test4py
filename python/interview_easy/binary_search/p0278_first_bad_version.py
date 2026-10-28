# https://leetcode.com/problems/first-bad-version/

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
