# https://leetcode.com/problems/contains-duplicate/

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
