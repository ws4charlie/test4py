# https://leetcode.com/problems/integer-to-roman/

class Solution:
    def intToRoman(self, num: int) -> str:
        rmap = {1:'I',
                4:'IV',
                5:'V',
                9:'IX',
                10:'X',
                40:'XL',
                50:'L',
                90:'XC',
                100:'C',
                400:'CD',
                500:'D',
                900:'CM',
                1000:'M'}
        
        res = ""
        bases = [1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1]
        
        for base in bases:
            times = num // base
            num = num % base
            res += times*rmap[base]
        return res

if __file__ == "__main__":
    sln = Solution()
    res = sln.intToRoman(1994)#MCMXCIV
    print(res)