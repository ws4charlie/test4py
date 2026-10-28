"""
6. Zigzag Conversion  (Medium)
https://leetcode.com/problems/zigzag-conversion/

Write the string s in a zigzag on numRows rows: go down the rows, then
diagonally back up, and repeat. Then read the rows from top to bottom and
return the result.

Example 1:
    Input:  s = "PAYPALISHIRING", numRows = 3
    Output: "PAHNAPLSIIGYIR"

        P   A   H   N
        A P L S I I G
        Y   I   R

Example 2:
    Input:  s = "PAYPALISHIRING", numRows = 4
    Output: "PINALSIGYAHRPI"

        P     I     N
        A   L S   I G
        Y A   H R
        P     I

Example 3:
    Input:  s = "A", numRows = 1
    Output: "A"

Constraints:
    1 <= s.length <= 1000
    s consists of English letters, ',' and '.'.
    1 <= numRows <= 1000
"""

class Solution:
    def convert(self, s: str, numRows: int) -> str:
        if numRows == 1:
            return s
        
        # list of rows
        rows = [''] * numRows
        
        cur_row = 0;
        go_down = True;
        for c in s:
            rows[cur_row] += c
            
            if go_down:
                cur_row += 1
            else:
                cur_row -= 1
                
            if cur_row == numRows-1 or cur_row == 0:
                go_down = not go_down
                
        res = ''
        for row in rows:
            res += row
        return res

if __name__ == "__main__":
    sln = Solution()
    res = sln.convert("PAYPALISHIRING", 3)
    print(res)