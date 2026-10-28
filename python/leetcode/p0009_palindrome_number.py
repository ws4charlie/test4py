"""
9. Palindrome Number  (Easy)
https://leetcode.com/problems/palindrome-number/

Given an integer x, return True if it reads the same forwards and
backwards.

Example 1:
    Input:  x = 121
    Output: True

Example 2:
    Input:  x = -121
    Output: False
    Explanation: backwards it reads "121-".

Example 3:
    Input:  x = 10
    Output: False
    Explanation: backwards it reads "01".

Constraints:
    -2^31 <= x <= 2^31 - 1

Follow-up: can you do it without converting x to a string?
"""

class Solution:
    def isPalindrome(self, x: 'int') -> 'bool':
        if x < 0:
            return False
        
        x_str = str(abs(x))
        l, r = 0, len(x_str) - 1
        while l < r:
            if x_str[l] != x_str[r]:
                return False
            l += 1
            r -= 1
            
        return True

if __name__ == "__main__":
    sln = Solution()
    res = sln.isPalindrome(121)
