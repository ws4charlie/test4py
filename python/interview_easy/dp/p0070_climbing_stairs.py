"""
70. Climbing Stairs  (Easy)
https://leetcode.com/problems/climbing-stairs/

A staircase has n steps and you climb 1 or 2 steps at a time. Return the
number of distinct ways to reach the top.

Example 1:
    Input:  n = 2
    Output: 2
    Explanation: 1 + 1, or 2

Example 2:
    Input:  n = 3
    Output: 3
    Explanation: 1 + 1 + 1, 1 + 2, or 2 + 1

Constraints:
    1 <= n <= 45
"""

class Solution:
    def climbStairs(self, n):
        # f(n) = f(n-1) + f(n-2): the last step is either 1 or 2 stairs
        a, b = 1, 1  # f(0), f(1)
        for _ in range(n - 1):
            a, b = b, a + b
        return b

if __name__ == "__main__":
    sln = Solution()
    assert [sln.climbStairs(n) for n in range(1, 8)] == [1, 2, 3, 5, 8, 13, 21]
    assert sln.climbStairs(45) == 1836311903
    print("ok")
