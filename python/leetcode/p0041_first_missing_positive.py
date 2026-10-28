"""
41. First Missing Positive  (Hard)
https://leetcode.com/problems/first-missing-positive/

Given an unsorted integer array nums, return the smallest positive integer
that is not in nums. The solution must run in O(n) time and use O(1)
extra space.

Example 1:
    Input:  nums = [1,2,0]
    Output: 3
    Explanation: 1 and 2 are present.

Example 2:
    Input:  nums = [3,4,-1,1]
    Output: 2
    Explanation: 1 is present, 2 is missing.

Example 3:
    Input:  nums = [7,8,9,11,12]
    Output: 1

Constraints:
    1 <= nums.length <= 10^5
    -2^31 <= nums[i] <= 2^31 - 1
"""

import heapq
from typing import List

class Solution_1:
    def firstMissingPositive(self, nums: List[int]) -> int:
        if not nums:
            return 1
        
        nums.sort()
        
        res = 1
        for v in nums:
            if v < res:
                continue
            elif v == res:
                res += 1
            else:
                break
        return res

class Solution_2:
    def firstMissingPositive(self, nums: List[int]) -> int:
        if not nums:
            return 1
        
        heapq.heapify(nums)
        
        res = 1
        while nums:
            v = heapq.heappop(nums)
            if v < res:
                continue
            elif v == res:
                res += 1
            else:
                break
        return res

if __name__ == "__main__":
    sln = Solution_1()
    nums = [0,2,2,1,1]
    res = sln.firstMissingPositive(nums)
    print(res)
