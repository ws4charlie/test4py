"""
169. Majority Element  (Easy)
https://leetcode.com/problems/majority-element/

Return the element that appears more than n / 2 times in an array of
size n. It always exists.

Example 1:
    Input:  nums = [3,2,3]
    Output: 3

Example 2:
    Input:  nums = [2,2,1,1,1,2,2]
    Output: 2

Constraints:
    1 <= nums.length <= 5 * 10^4
    -10^9 <= nums[i] <= 10^9
    The majority element always exists.

Follow-up: can you do it in O(n) time and O(1) space?
"""

class Solution:
    def majorityElement(self, nums):
        # Boyer-Moore vote: the majority survives pairing off with every
        # other value. O(n) time, O(1) space.
        cand, count = None, 0
        for v in nums:
            if count == 0:
                cand = v
            count += 1 if v == cand else -1
        return cand

if __name__ == "__main__":
    sln = Solution()
    assert sln.majorityElement([3, 2, 3]) == 3
    assert sln.majorityElement([2, 2, 1, 1, 1, 2, 2]) == 2
    assert sln.majorityElement([1]) == 1
    print("ok")
