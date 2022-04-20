
# https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/

class Solution:
    def bs(self, nums, i, j, target):
        if i > j:
            return -1
        
        mid = (i+j) // 2
        if target == nums[mid]:
            return mid
        elif target < nums[mid]:
            return self.bs(nums, i, mid-1, target)
        else:
            return self.bs(nums, mid+1, j, target)
        
    def searchRange(self, nums: List[int], target: int) -> List[int]:
        if not nums:
            return [-1, -1]
        
        pos = self.bs(nums, 0, len(nums)-1,  target)
        
        if pos == -1:
            return [-1, -1]
        
        i, j = pos, pos
        while i >= 0 and nums[i] == target:
            i -= 1
        while j < len(nums) and nums[j] == target:
            j += 1
            
        return [i+1, j-1]

if __name__ == "__main__":
    nums = [1,3,4,7,7,8,8,9]
    sln = Solution()
    res = sln.searchRange(nums, 7)