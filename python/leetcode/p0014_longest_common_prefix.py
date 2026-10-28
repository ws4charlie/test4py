# https://leetcode.com/problems/longest-common-prefix/

class Solution:
    def longestCommonPrefix(self, strs) -> str:
        if len(strs) == 0:
            return ''
        
        lmin = len(strs[0])
        for s in strs:
            lmin = min(lmin, len(s))
            
        res = ''
        for i in range(lmin):
            c = ''
            for s in strs:
                if c == '':
                    c = s[i]
                elif s[i] != c:
                    return res
            res += c
        return res

if __name__ == '__main__':
    sln = Solution()
    res = sln.longestCommonPrefix(["flower","flow","flight"])
    print(res)