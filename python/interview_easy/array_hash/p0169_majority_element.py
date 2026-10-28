# https://leetcode.com/problems/majority-element/

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
