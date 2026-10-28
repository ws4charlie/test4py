"""
28. Find the Index of the First Occurrence in a String  (Easy)
https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/
(formerly "Implement strStr()")

Return the index where needle first appears in haystack, or -1 if it
doesn't appear.

Example 1:
    Input:  haystack = "sadbutsad", needle = "sad"
    Output: 0
    Explanation: "sad" appears at 0 and 6; the first is 0.

Example 2:
    Input:  haystack = "leetcode", needle = "leeto"
    Output: -1

Constraints:
    1 <= haystack.length, needle.length <= 10^4
    Both consist of lowercase English letters only.
"""

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