# https://leetcode.com/problems/move-zeroes/

class Solution:
    def moveZeroes(self, nums):
        """
        Do not return anything, modify nums in-place instead.
        """
        # k is where the next non-zero goes; swapping keeps relative order
        k = 0
        for i in range(len(nums)):
            if nums[i] != 0:
                nums[k], nums[i] = nums[i], nums[k]
                k += 1

if __name__ == "__main__":
    sln = Solution()
    nums = [0, 1, 0, 3, 12]
    sln.moveZeroes(nums)
    assert nums == [1, 3, 12, 0, 0]
    nums = [0]
    sln.moveZeroes(nums)
    assert nums == [0]
    nums = [1, 2]
    sln.moveZeroes(nums)
    assert nums == [1, 2]
    print("ok")
