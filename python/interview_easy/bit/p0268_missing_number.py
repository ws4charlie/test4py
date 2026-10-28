# https://leetcode.com/problems/missing-number/

class Solution:
    def missingNumber(self, nums):
        # XOR every index 0..n with every value; all but the missing one pair up
        res = len(nums)
        for i, v in enumerate(nums):
            res ^= i ^ v
        return res

    def missingNumberSum(self, nums):
        n = len(nums)
        return n * (n + 1) // 2 - sum(nums)

if __name__ == "__main__":
    sln = Solution()
    for f in (sln.missingNumber, sln.missingNumberSum):
        assert f([3, 0, 1]) == 2
        assert f([0, 1]) == 2
        assert f([9, 6, 4, 2, 3, 5, 7, 0, 1]) == 8
        assert f([1]) == 0
    print("ok")
