
# https://leetcode.com/problems/longest-palindromic-substring/

#Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

#Example 1:
#Input: "babad"
#Output: "bab"
#Note: "aba" is also a valid answer.

#Example 2:
#Input: "cbbd"
#Output: "bb"


class Solution:
    def longestPalindrome(self, s: str) -> str:
        if len(s) <= 1:
            return s
        
        res = ''
        for i in range(len(s) - 1):
            sub1 = self.maxPalindrome(s, i, i)
            sub2 = self.maxPalindrome(s, i, i + 1)
            if len(sub1) > len(res):
                res = sub1
            if len(sub2) > len(res):
                res = sub2
        
        return res
    
    def maxPalindrome(self, s, i, j):
        while i >=0 and j < len(s) and s[i] == s[j]:
            i, j = i-1, j+1
        return s[i+1:j]
   
if __name__ == '__main__':
    sln = Solution()
    print("")