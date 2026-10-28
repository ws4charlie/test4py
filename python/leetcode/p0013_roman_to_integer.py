"""
13. Roman to Integer  (Easy)
https://leetcode.com/problems/roman-to-integer/

Convert a Roman numeral to an integer.

    Symbol   Value        Subtractive forms
    I        1            IV   4
    V        5            IX   9
    X        10           XL   40
    L        50           XC   90
    C        100          CD   400
    D        500          CM   900
    M        1000

Symbols are normally written largest to smallest and added up. When a
smaller symbol comes before a larger one (the subtractive forms above),
it is subtracted instead.

Example 1:
    Input:  s = "III"
    Output: 3

Example 2:
    Input:  s = "LVIII"
    Output: 58
    Explanation: L = 50, V = 5, III = 3

Example 3:
    Input:  s = "MCMXCIV"
    Output: 1994
    Explanation: M = 1000, CM = 900, XC = 90, IV = 4

Constraints:
    1 <= s.length <= 15
    s contains only I, V, X, L, C, D, M.
    s is a valid Roman numeral in the range [1, 3999].
"""

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