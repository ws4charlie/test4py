"""
38. Count and Say  (Medium)
https://leetcode.com/problems/count-and-say/

The count-and-say sequence starts with countAndSay(1) = "1". Each later
term describes the previous one: read it left to right and replace every
run of the same digit with the run length followed by the digit.

    "3322251" -> two 3s, three 2s, one 5, one 1 -> "23321511"

Return the n-th term.

Example 1:
    Input:  n = 4
    Output: "1211"
    Explanation:
        countAndSay(1) = "1"
        countAndSay(2) = "11"     (one 1)
        countAndSay(3) = "21"     (two 1s)
        countAndSay(4) = "1211"   (one 2, one 1)

Example 2:
    Input:  n = 1
    Output: "1"

Constraints:
    1 <= n <= 30

Follow-up: can you solve it iteratively?
"""

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
