"""
20. Valid Parentheses  (Easy)
https://leetcode.com/problems/valid-parentheses/

s contains only the characters '(', ')', '[', ']', '{' and '}'. Return
True if it is valid:
    - every opening bracket is closed by a bracket of the same type,
    - brackets close in the right order (innermost first),
    - every closing bracket has a matching opening bracket.

Example 1:
    Input:  s = "()"
    Output: True

Example 2:
    Input:  s = "()[]{}"
    Output: True

Example 3:
    Input:  s = "(]"
    Output: False

Example 4:
    Input:  s = "([])"
    Output: True

Example 5:
    Input:  s = "([)]"
    Output: False

Constraints:
    1 <= s.length <= 10^4
    s consists of brackets '()[]{}' only.
"""

class Solution:
    def match(self, l, r):
        if l == '(' and r == ')':
            return True
        if l == '[' and r == ']':
            return True
        if l == '{' and r == '}':
            return True
        return False
    
    def isValid(self, s: str) -> bool:
        if not s:
            return True
        if len(s) % 2 != 0:
            return False
        
        l = []
        for i in range(len(s)):
            if not l:
                l.append(s[i])
            else:
                if self.match(l[-1], s[i]):
                    l.pop()
                else:
                    l.append(s[i])
                
        if not l:
            return True
        return False

if __name__ == "__main__":
    sln = Solution()
    res = sln.isValid("{[]}")
    print(res)