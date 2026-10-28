"""
217. Contains Duplicate  (Easy)
https://leetcode.com/problems/contains-duplicate/

Return True if any value appears at least twice in nums, False if every
value is distinct.

Example 1:
    Input:  nums = [1,2,3,1]
    Output: True
    Explanation: 1 appears at index 0 and 3.

Example 2:
    Input:  nums = [1,2,3,4]
    Output: False

Example 3:
    Input:  nums = [1,1,1,3,3,4,3,2,4,2]
    Output: True

Constraints:
    1 <= nums.length <= 10^5
    -10^9 <= nums[i] <= 10^9
"""

class Solution:
    def containsDuplicate(self, nums):
        seen = set()
        for v in nums:
            if v in seen:
                return True
            seen.add(v)
        return False

if __name__ == "__main__":
    sln = Solution()
    assert sln.containsDuplicate([1, 2, 3, 1])
    assert not sln.containsDuplicate([1, 2, 3, 4])
    assert sln.containsDuplicate([1, 1, 1, 3, 3, 4, 3, 2, 4, 2])
    assert not sln.containsDuplicate([])
    print("ok")
