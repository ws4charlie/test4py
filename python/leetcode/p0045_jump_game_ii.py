"""
45. Jump Game II  (Medium)
https://leetcode.com/problems/jump-game-ii/

You start at index 0. nums[i] is the longest jump you can make from index
i (any length from 0 to nums[i]). Return the minimum number of jumps to
reach the last index. The last index is always reachable.

Example 1:
    Input:  nums = [2,3,1,1,4]
    Output: 2
    Explanation: jump 1 step to index 1, then 3 steps to the last index.

Example 2:
    Input:  nums = [2,3,0,1,4]
    Output: 2

Constraints:
    1 <= nums.length <= 10^4
    0 <= nums[i] <= 1000
    The last index is reachable from index 0.
"""

from typing import List

class Solution_greedy:
    def jump(self, nums: List[int]) -> int:
        n = len(nums)
        if n == 1:
            return 0
        
        cur_max = nums[0]
        next_max = 0 # whatever
        jumps = 1
        for i in range(1, n):
            # scan elems in [i+1, pre_max]
            next_max = max(next_max, i+nums[i])
            # after scanning the last
            if i == cur_max and i < n-1:
                jumps += 1
                cur_max = next_max
        
        return jumps
       
if __name__ == "__main__":
    sln = Solution_greedy()
    nums = [2,3,1,1,4]
    res = sln.jump(nums)
