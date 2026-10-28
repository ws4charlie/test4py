# https://leetcode.com/problems/excel-sheet-column-title/

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