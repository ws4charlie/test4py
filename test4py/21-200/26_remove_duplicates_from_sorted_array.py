# https://leetcode.com/problems/remove-duplicates-from-sorted-array/submissions/

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
