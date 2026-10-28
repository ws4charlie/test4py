"""
1306. Jump Game III  (Medium)
https://leetcode.com/problems/jump-game-iii/

arr is an array of non-negative integers and you start at index start.
From index i you can jump to i + arr[i] or i - arr[i], but never outside
the array. Return True if you can reach any index whose value is 0.

Example 1:
    Input:  arr = [4,2,3,0,3,1,2], start = 5
    Output: True
    Explanation: index 5 -> 4 -> 1 -> 3 (value 0)

Example 2:
    Input:  arr = [4,2,3,0,3,1,2], start = 0
    Output: True
    Explanation: index 0 -> 4 -> 1 -> 3

Example 3:
    Input:  arr = [3,0,2,1,2], start = 2
    Output: False
    Explanation: from index 2 you can only ever reach indices 0, 3 and
                 4; the only 0 is at index 1.

Constraints:
    1 <= arr.length <= 5 * 10^4
    0 <= arr[i] < arr.length
    0 <= start < arr.length
"""

from typing import List

class Solution_BFS:
    # i is the root
    def reach(self, arr, i, seen):
        if arr[i] == 0:
            return True
        else:
            seen.add(i)
        
        left = i - arr[i]
        if left >= 0 and left not in seen:
            if self.reach(arr, left, seen):
                return True
            
        right = i + arr[i]
        if right < len(arr) and right not in seen:
            if self.reach(arr, right, seen):
                return True
    
    # Treats start as the root a tree flatted in arr
    def canReach(self, arr: List[int], start: int) -> bool:
        seen = set()
        return self.reach(arr, start, seen)

class Solution_Topdown:
    # i is the root
    def reach(self, arr, i, mem, seen):
        if i in mem:
            return mem[i]
        
        seen.add(i)
        
        left = i - arr[i]
        if left >= 0 and left not in seen:
            if self.reach(arr, left, mem, seen):
                mem[i] = True
                return True
        
        right = i + arr[i]
        if right < len(arr) and right not in seen:
            if self.reach(arr, right, mem, seen):
                mem[i] = True
                return True
        
        mem[i] = False
        return False

    def canReach(self, arr: List[int], start: int) -> bool:
        if not arr:
            return False
        
        # init mem
        mem = {}
        for i, v in enumerate(arr):
            if v == 0:
                mem[i] = True
                
        # to avoid cycle search, keep posotions we have seen already
        seen = set()
        
        # start
        return self.reach(arr, start, mem, seen)
   
if __name__ == "__main__":
    sln = Solution_Topdown()
    arr = [4,2,3,0,3,1,2]
    res = sln.canReach(arr, 5)
