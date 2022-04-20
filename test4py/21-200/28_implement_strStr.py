# https://leetcode.com/problems/implement-strstr/

class Solution:
    def strStr(self, haystack: str, needle: str) -> int:
        if not needle:
            return 0
        
        for i in range(len(haystack) - len(needle) + 1):
            match = True
            for j in range(len(needle)):
                if haystack[i+j] != needle[j]:
                    match = False
                    break
            if match:
                return i
            
        return -1

if __name__ == "__main__":
    sln = Solution()
    pos = sln.strStr("Hello", "ll")
    print(pos)