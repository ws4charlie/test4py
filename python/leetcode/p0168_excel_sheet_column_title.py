"""
168. Excel Sheet Column Title  (Easy)
https://leetcode.com/problems/excel-sheet-column-title/

Given a column number, return its title as it appears in an Excel sheet:

    1  -> A
    2  -> B
    ...
    26 -> Z
    27 -> AA
    28 -> AB
    ...

Example 1:
    Input:  columnNumber = 1
    Output: "A"

Example 2:
    Input:  columnNumber = 28
    Output: "AB"

Example 3:
    Input:  columnNumber = 701
    Output: "ZY"

Constraints:
    1 <= columnNumber <= 2^31 - 1
"""

class Solution:
    def convertToTitle(self, n: int) -> str:
        scale = 26
        
        res = ''
        while n > 0:
            remd = n % scale
            if remd == 0:
                res = 'Z' + res
            else:
                res = chr(ord('A') + remd - 1) + res
            n = (n-1) // scale
            
        return res

if __name__ == "__main__":
    sln = Solution()
    res = sln.convertToTitle(28)