# https://leetcode.com/problems/binary-search/

class Solution:
    def search(self, nums, target):
        # closed interval [lo, hi]
        lo, hi = 0, len(nums) - 1
        while lo <= hi:
            mid = (lo + hi) // 2
            if nums[mid] == target:
                return mid
            if nums[mid] < target:
                lo = mid + 1
            else:
                hi = mid - 1
        return -1

if __name__ == "__main__":
    sln = Solution()
    assert sln.search([-1, 0, 3, 5, 9, 12], 9) == 4
    assert sln.search([-1, 0, 3, 5, 9, 12], 2) == -1
    assert sln.search([5], 5) == 0
    assert sln.search([], 1) == -1
    print("ok")
