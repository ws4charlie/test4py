"""
53. Maximum Subarray  (Medium)
https://leetcode.com/problems/maximum-subarray/

Return the largest sum of any non-empty contiguous subarray of nums.

Example 1:
    Input:  nums = [-2,1,-3,4,-1,2,1,-5,4]
    Output: 6
    Explanation: the subarray [4,-1,2,1] sums to 6.

Example 2:
    Input:  nums = [1]
    Output: 1

Example 3:
    Input:  nums = [5,4,-1,7,8]
    Output: 23
    Explanation: the whole array.

Constraints:
    1 <= nums.length <= 10^5
    -10^4 <= nums[i] <= 10^4

Follow-up: after the O(n) solution, try divide and conquer.
"""

from typing import List

class Solution_N:
    def maxSubArray(self, nums: List[int]) -> int:
        sum = 0;
        tot = nums[0];
        for i in range(len(nums)):
            sum += nums[i]
            tot = max(tot, sum)
            sum = max(sum, 0)
        return tot

class Solution_N2:
    def maxSubArray(self, nums: List[int]) -> int:
        n = len(nums)
        max_sum = nums[0]
        for i in range(n):
            sum = 0
            for j in range(i, n):
                sum += nums[j]
                max_sum = max(max_sum, sum)
                
        return max_sum

if __name__ == "__main__":
    nums = [-2,1,-3,4,-1,2,1,-5,4]
    sln = Solution_N()
    res = sln.maxSubArray(nums)
