"""
7. Reverse Integer  (Medium)
https://leetcode.com/problems/reverse-integer/

Given a signed 32-bit integer x, return x with its digits reversed. If the
result falls outside the signed 32-bit range [-2^31, 2^31 - 1], return 0.
Assume you can't store 64-bit integers.

Example 1:
    Input:  x = 123
    Output: 321

Example 2:
    Input:  x = -123
    Output: -321

Example 3:
    Input:  x = 120
    Output: 21

Constraints:
    -2^31 <= x <= 2^31 - 1
"""

class Solution:
    def reverse(self, x):
        negative = x < 0
        x = abs(x)
        
        res = 0
        while x > 0:
            res = res * 10 + x % 10
            x = x // 10
        
        if res < -2**31 or res > 2**31 - 1:
            return 0
        
        if negative:
            return -res
        else:
            return res

if __name__ == '__main__':
    sln = Solution()
    print("")