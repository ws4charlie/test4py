"""
125. Valid Palindrome  (Easy)
https://leetcode.com/problems/valid-palindrome/

Lowercase s and drop every character that isn't a letter or digit. Return
True if what's left reads the same forwards and backwards.

Example 1:
    Input:  s = "A man, a plan, a canal: Panama"
    Output: True
    Explanation: "amanaplanacanalpanama" is a palindrome.

Example 2:
    Input:  s = "race a car"
    Output: False
    Explanation: "raceacar" is not a palindrome.

Example 3:
    Input:  s = " "
    Output: True
    Explanation: nothing is left, and the empty string is a palindrome.

Constraints:
    1 <= s.length <= 2 * 10^5
    s consists of printable ASCII characters.
"""

class Solution:
    def isPalindrome(self, s):
        # skip non-alphanumeric chars from both ends, compare lowercased
        l, r = 0, len(s) - 1
        while l < r:
            if not s[l].isalnum():
                l += 1
            elif not s[r].isalnum():
                r -= 1
            elif s[l].lower() != s[r].lower():
                return False
            else:
                l, r = l + 1, r - 1
        return True

if __name__ == "__main__":
    sln = Solution()
    assert sln.isPalindrome("A man, a plan, a canal: Panama")
    assert not sln.isPalindrome("race a car")
    assert sln.isPalindrome(" ")
    assert not sln.isPalindrome("0P")
    print("ok")
