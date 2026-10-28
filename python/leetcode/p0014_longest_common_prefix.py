"""
14. Longest Common Prefix  (Easy)
https://leetcode.com/problems/longest-common-prefix/

Return the longest string that every string in strs starts with, or ""
if there is none.

Example 1:
    Input:  strs = ["flower","flow","flight"]
    Output: "fl"

Example 2:
    Input:  strs = ["dog","racecar","car"]
    Output: ""

Constraints:
    1 <= strs.length <= 200
    0 <= strs[i].length <= 200
    strs[i] consists of lowercase English letters (if non-empty).
"""

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