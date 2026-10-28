"""
33. Search in Rotated Sorted Array  (Medium)
https://leetcode.com/problems/search-in-rotated-sorted-array/

nums was sorted in ascending order with distinct values, then possibly
rotated left by an unknown k (1 <= k < n):

    sorted:   0 1 2 4 5 6 7
    rotated:  4 5 6 7 0 1 2      (k = 3)

Return the index of target in nums, or -1 if it isn't there. The solution
must run in O(log n) time.

Example 1:
    Input:  nums = [4,5,6,7,0,1,2], target = 0
    Output: 4

Example 2:
    Input:  nums = [4,5,6,7,0,1,2], target = 3
    Output: -1

Example 3:
    Input:  nums = [1], target = 0
    Output: -1

Constraints:
    1 <= nums.length <= 5000
    -10^4 <= nums[i], target <= 10^4
    All values of nums are distinct.
    nums is an ascending array that may have been rotated.
"""

class Solution:
    def bs(self, nums, i, j, target) -> int:
        mid = (i+j)//2
        if mid == i:
            if nums[i] == target:
                return i
            if nums[j] == target:
                return j
            return -1
        
        pv = nums[0]
        mv = nums[mid]
        if target >= pv:
            if mv >= pv:
                if target <= mv:
                    return self.bs(nums, i, mid, target)
                else:
                    return self.bs(nums, mid, j, target)
            else:
                return self.bs(nums, i, mid, target)
        else:
            if mv < pv:
                if target <= mv:
                    return self.bs(nums, i, mid, target)
                else:
                    return self.bs(nums, mid, j, target)
            else:
                return self.bs(nums, mid, j, target);
        
    def search(self, nums, target: int) -> int:
        if len(nums) == 0:
            return -1
        if len(nums) == 1:
            if nums[0] == target:
                return 0
            else:
                return -1
            
        index = self.bs(nums, 0, len(nums)-1, target)
        return index

if __name__ == "__main__":
    nums = [4,5,6,7,0,1,2]
    target = 0
    sln = Solution()
    res = sln.search(nums, target)
