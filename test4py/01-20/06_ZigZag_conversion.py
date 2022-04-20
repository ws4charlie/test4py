# https://leetcode.com/problems/zigzag-conversion/

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

if __file__ == "__main__":
    sln = Solution()
    res = sln.convert("PAYPALISHIRING", 3)
    print(res)