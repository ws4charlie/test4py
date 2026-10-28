"""
136. Single Number  (Easy)
https://leetcode.com/problems/single-number/

Every element of nums appears exactly twice, except one that appears
once. Return that one. Use O(n) time and O(1) extra space.

Example 1:
    Input:  nums = [2,2,1]
    Output: 1

Example 2:
    Input:  nums = [4,1,2,1,2]
    Output: 4

Example 3:
    Input:  nums = [1]
    Output: 1

Constraints:
    1 <= nums.length <= 3 * 10^4
    -3 * 10^4 <= nums[i] <= 3 * 10^4
    Every element appears twice except for exactly one.
"""

class Solution:
    def singleNumber(self, nums):
        # x ^ x == 0 and x ^ 0 == x, so every pair cancels out
        res = 0
        for v in nums:
            res ^= v
        return res

if __name__ == "__main__":
    sln = Solution()
    assert sln.singleNumber([2, 2, 1]) == 1
    assert sln.singleNumber([4, 1, 2, 1, 2]) == 4
    assert sln.singleNumber([1]) == 1
    assert sln.singleNumber([-3, 7, 7]) == -3
    print("ok")
