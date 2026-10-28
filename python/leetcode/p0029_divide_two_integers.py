"""
29. Divide Two Integers  (Medium)
https://leetcode.com/problems/divide-two-integers/

Divide dividend by divisor without using *, / or %, and return the
quotient truncated toward zero (8.345 -> 8, -2.7335 -> -2). Results are
limited to the signed 32-bit range [-2^31, 2^31 - 1]: anything larger is
returned as 2^31 - 1, anything smaller as -2^31.

Example 1:
    Input:  dividend = 10, divisor = 3
    Output: 3
    Explanation: 10 / 3 = 3.33333..., truncated to 3

Example 2:
    Input:  dividend = 7, divisor = -3
    Output: -2
    Explanation: 7 / -3 = -2.33333..., truncated to -2

Constraints:
    -2^31 <= dividend, divisor <= 2^31 - 1
    divisor != 0
"""

class Solution:
    def divide(self, dividend: int, divisor: int) -> int:
        neg = (dividend > 0 and divisor < 0) or (dividend < 0 and divisor > 0);
        dividend = abs(dividend)
        divisor = abs(divisor)
        
        res = 0
        times = 1
        while dividend >= divisor:
            if dividend >= divisor * times:
                dividend -= divisor * times
                res += times
                times *= 2
            else:
                times = 1
        
        res = -res if neg else res
        if res < -2**31 or res > 2**31 - 1:
            return 2**31 - 1
        else:
            return res

if __name__ == "__main__":
    sln = Solution()
    res = sln.divide(-2147483648, 1)