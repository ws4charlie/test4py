"""
3. Longest Substring Without Repeating Characters  (Medium)
https://leetcode.com/problems/longest-substring-without-repeating-characters/

Given a string s, return the length of the longest substring (a
contiguous run of characters) in which no character appears twice.

Example 1:
    Input:  s = "abcabcbb"
    Output: 3
    Explanation: "abc" (also "bca", "cab")

Example 2:
    Input:  s = "bbbbb"
    Output: 1
    Explanation: "b"

Example 3:
    Input:  s = "pwwkew"
    Output: 3
    Explanation: "wke". "pwke" doesn't count: it is a subsequence, not a
                 substring.

Constraints:
    0 <= s.length <= 5 * 10^4
    s consists of English letters, digits, symbols and spaces.
"""

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
    sln = Solution()
    res = sln.lengthOfLongestSubstring("pwwkew")
    print(res)