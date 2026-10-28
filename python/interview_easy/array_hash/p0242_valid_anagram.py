# https://leetcode.com/problems/valid-anagram/

class Solution:
    def isAnagram(self, s, t):
        if len(s) != len(t):
            return False
        count = {}
        for c in s:
            count[c] = count.get(c, 0) + 1
        for c in t:
            if count.get(c, 0) == 0:
                return False
            count[c] -= 1
        return True

if __name__ == "__main__":
    sln = Solution()
    assert sln.isAnagram("anagram", "nagaram")
    assert not sln.isAnagram("rat", "car")
    assert not sln.isAnagram("a", "ab")
    assert sln.isAnagram("", "")
    print("ok")
