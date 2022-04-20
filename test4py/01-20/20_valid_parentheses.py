# https://leetcode.com/problems/valid-parentheses/

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