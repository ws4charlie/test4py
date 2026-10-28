
# https://leetcode.com/problems/divide-two-integers/

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