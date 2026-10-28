
class Solution:        
    def countAndSay(self, n: int) -> str:
        if n == 0:
            return ""
        if n == 1:
            return "1"
                
        s = self.countAndSay(n-1)
        
        cnt = 1
        res = ""
        for i in range(1, len(s)):
            if s[i] == s[i-1]:
                cnt += 1
            else:
                res += str(cnt) + s[i-1]
                cnt = 1
        
        res += str(cnt) + s[-1]
        return res

if __name__ == "__main__":
    sln = Solution()
    res = sln.countAndSay(5)
    print(res)
