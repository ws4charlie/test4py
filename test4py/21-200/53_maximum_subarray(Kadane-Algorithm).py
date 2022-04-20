
# https://leetcode.com/problems/maximum-subarray/

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
