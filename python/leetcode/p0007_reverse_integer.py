# https://leetcode.com/problems/reverse-integer/

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