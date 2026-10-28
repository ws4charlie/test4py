"""
704. Binary Search  (Easy)
https://leetcode.com/problems/binary-search/

nums is sorted in ascending order. Return the index of target, or -1 if
it isn't there. The solution must run in O(log n) time.

Example 1:
    Input:  nums = [-1,0,3,5,9,12], target = 9
    Output: 4

Example 2:
    Input:  nums = [-1,0,3,5,9,12], target = 2
    Output: -1

Constraints:
    1 <= nums.length <= 10^4
    -10^4 < nums[i], target < 10^4
    All values of nums are unique.
    nums is sorted in ascending order.
"""

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
