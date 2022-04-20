# https://leetcode.com/problems/longest-substring-without-repeating-characters/

class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        seen = {}
        i, lmax = -1, 0
        
        for j in range(len(s)):
            c = s[j]
            if c not in seen:
                lmax = max(lmax, j - i)
            else:
                pos = seen[c]
                for k in range(i+1, pos):
                    del seen[s[k]]
                i = pos
            seen[c] = j
        
        return lmax

if __name__ == '__main__':
    sln = Solution2()
    res = sln.lengthOfLongestSubstring("pwwkew")
    print(res)