# https://leetcode.com/problems/roman-to-integer/

class Solution:
    def romanToInt(self, s: str) -> int:
        rmap = {'I':1,
               'IV':4,
                'V':5,
               'IX':9,
                'X':10,
               'XL':40,
                'L':50,
               'XC':90,
                'C':100,
               'CD':400,
                'D':500,
               'CM':900,
                'M':1000}
        
        i, res = 0, 0
        while i < len(s):
            if i+1 < len(s):# 2-char symbol
                sb = s[i:i+2]
                if sb in rmap:
                    res += rmap[sb]
                    i += 2
                    continue
            res += rmap[s[i]]
            i += 1
        return res

if __name__ == "__main__":
    sln = Solution()
    res = sln.romanToInt("MCMXCIV")