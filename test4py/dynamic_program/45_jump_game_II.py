
# https://leetcode.com/problems/jump-game-ii/

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
