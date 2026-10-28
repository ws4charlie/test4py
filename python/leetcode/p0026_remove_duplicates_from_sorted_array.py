"""
26. Remove Duplicates from Sorted Array  (Easy)
https://leetcode.com/problems/remove-duplicates-from-sorted-array/

nums is sorted in non-decreasing order. Remove the duplicates in place so
each value appears once, keeping the original order, and return k, the
number of unique values. The first k elements of nums must hold those
values; whatever comes after them doesn't matter.

Example 1:
    Input:  nums = [1,1,2]
    Output: 2, nums = [1,2,_]

Example 2:
    Input:  nums = [0,0,1,1,1,2,2,3,3,4]
    Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]

Constraints:
    1 <= nums.length <= 3 * 10^4
    -100 <= nums[i] <= 100
    nums is sorted in non-decreasing order.
"""

from typing import List

class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        i, j = -1, 0
        while j < len(nums):
            if j == 0 or nums[i] != nums[j]:
                nums[i+1] = nums[j]
                i, j = i+1, j+1
            else:
                j += 1
        return i+1

if __name__ == "__main__":
    sln = Solution()
    nums = [1,2,2,3,4,4,6]
    res = sln.removeDuplicates(nums)
    for i in range(res):
        print(nums[i])
