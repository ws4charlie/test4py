# https://leetcode.com/problems/single-number/

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
