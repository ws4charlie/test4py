"""
268. Missing Number  (Easy)
https://leetcode.com/problems/missing-number/

nums holds n distinct numbers from the range [0, n], so exactly one
number in the range is missing. Return it.

Example 1:
    Input:  nums = [3,0,1]
    Output: 2
    Explanation: n = 3, range [0,3], 2 is missing.

Example 2:
    Input:  nums = [0,1]
    Output: 2
    Explanation: n = 2, range [0,2], 2 is missing.

Example 3:
    Input:  nums = [9,6,4,2,3,5,7,0,1]
    Output: 8

Constraints:
    n == nums.length
    1 <= n <= 10^4
    0 <= nums[i] <= n
    All values of nums are unique.

Follow-up: can you do it in O(1) extra space and O(n) time?
"""

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
