# https://leetcode.com/problems/valid-palindrome/

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
