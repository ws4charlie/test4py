"""
242. Valid Anagram  (Easy)
https://leetcode.com/problems/valid-anagram/

Return True if t is an anagram of s: the same letters with the same
counts, in any order.

Example 1:
    Input:  s = "anagram", t = "nagaram"
    Output: True

Example 2:
    Input:  s = "rat", t = "car"
    Output: False

Constraints:
    1 <= s.length, t.length <= 5 * 10^4
    s and t consist of lowercase English letters.

Follow-up: what if the inputs contain Unicode characters?
"""

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
