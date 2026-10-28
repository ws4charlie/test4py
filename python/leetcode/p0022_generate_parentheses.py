"""
22. Generate Parentheses  (Medium)
https://leetcode.com/problems/generate-parentheses/

Given n pairs of parentheses, return every well-formed combination, in
any order.

Example 1:
    Input:  n = 3
    Output: ["((()))","(()())","(())()","()(())","()()()"]

Example 2:
    Input:  n = 1
    Output: ["()"]

Constraints:
    1 <= n <= 8
"""

from typing import List

class Solution22:
    def back_trace(self, res, comb, _open, close, n):
        if len(comb) == n*2:
            res.append(comb)
            return
        
        if _open < n:
            comb1 = comb + '('
            self.back_trace(res, comb1, _open+1, close, n)
        if _open > close:
            comb2 = comb + ')'
            self.back_trace(res, comb2, _open, close+1, n)
        
    def generateParenthesis(self, n: int) -> List[str]:
        res = []
        self.back_trace(res, "", 0, 0, n)
        return res

if __name__ == "__main__":
    sln = Solution22()
    res = sln.generateParenthesis(3)
    print(res)
