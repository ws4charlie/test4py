"""
12. Integer to Roman  (Medium)
https://leetcode.com/problems/integer-to-roman/

Convert an integer to a Roman numeral.

    Symbol   Value        Subtractive forms
    I        1            IV   4
    V        5            IX   9
    X        10           XL   40
    L        50           XC   90
    C        100          CD   400
    D        500          CM   900
    M        1000

Write the numeral from the largest place value to the smallest. A digit 4
or 9 uses a subtractive form. Only I, X, C and M can repeat, at most 3
times in a row; V, L and D never repeat.

Example 1:
    Input:  num = 3749
    Output: "MMMDCCXLIX"
    Explanation: 3000 = MMM, 700 = DCC, 40 = XL, 9 = IX

Example 2:
    Input:  num = 58
    Output: "LVIII"
    Explanation: 50 = L, 8 = VIII

Example 3:
    Input:  num = 1994
    Output: "MCMXCIV"
    Explanation: 1000 = M, 900 = CM, 90 = XC, 4 = IV

Constraints:
    1 <= num <= 3999
"""

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

if __name__ == "__main__":
    sln = Solution()
    res = sln.intToRoman(1994)#MCMXCIV
    print(res)