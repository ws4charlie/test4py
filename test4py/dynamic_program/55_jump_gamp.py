
# https://leetcode.com/problems/jump-game/

class Solution_brute_force:
    def back_trace(self, nums, pos):
        if pos >= len(nums) - 1:
            return True

      # we can optimize by making the biggest jump first at each position
      # for i in range(nums[pos], 0, -1):
        for i in range(1, nums[pos]+1):
            canJump = self.back_trace(nums, pos+i)
            if canJump:
                return True
        return False
    
    def canJump(self, nums: List[int]) -> bool:
        if len(nums) == 1:
            return True
        
        return self.back_trace(nums, 0)

class Solution_Greedy:
    """
    # i is a position, define f(i) as the furthest position we can reach if we start from i
    # f(0) = 0+nums[0]
    # f(1) = max(f(0), 1+nums[1]),  condition f(0) >= 1
    # f(2) = max(f(1), 2+nums[2]),  condigion f(1) >= 2
    # ...
    # f(i) = max(f(i-1), i+nums[i]) condition f(i-1) >= i
    # The problem then can be defined as: Can we satisfy all above conditions and reach to len(nums)-1?
    """
    def canJump(self, nums: List[int]) -> bool:
        if len(nums) == 1:
            return True
        
        maxpos = nums[0]
        pos = 1
        while pos <= maxpos and pos < len(nums):
            maxpos = max(maxpos, pos+nums[pos])
            pos += 1
            
        return maxpos >= len(nums)-1

# problem definition
# i is a position
# f(i) returns True if we can jump from i to the end and returns False otherwise
# f(i) == True, only if any of {f(j)} returns True, i < j < i+nums[i]
class Solution_Topdown:
    def can_jump(self, nums, pos, mem):
        if pos in mem:
            return mem[pos]
        if pos >= len(nums):
            return True
        
        canjump = False
        for i in range(1, nums[pos]+1):
            if (self.can_jump(nums, pos+i, mem)):
                canjump = True
                break

        mem[pos] = canjump
        return canjump
    
    def canJump(self, nums: List[int]) -> bool:
        if len(nums) == 1:
            return True
        
        # init mem
        mem = {}
        mem[len(nums)-1] = True

        return self.can_jump(nums, 0, mem)

class Solution_Bottomup:
    def canJump(self, nums: List[int]) -> bool:
        if len(nums) == 1:
            return True
        
        # init mem
        mem = {}
        mem[len(nums)-1] = True
        
        # start
        for i in range(len(nums)-2, -1, -1):
            canjump = False
            longest = min(len(nums)-1, i+nums[i])
            for j in range(i+1, longest+1):
                if mem[j]:
                    canjump = True
                    break
            mem[i] = canjump
            
        return mem[0]

class Solution_backtrace_jump_over_zeros:
    def back_trace(self, nums, beg):
        # find first 0 in [p0, len(nums))
        p0 = beg
        while p0 < len(nums) and nums[p0] != 0:
            p0 += 1
        # no more 0 or single 0 at the end
        if p0 >= len(nums)-1:
            return True
        
        # see if we can jump over this 0
        can_jump_over = False
        for j in range(p0, -1, -1):
            if nums[j] > p0 - j:
                can_jump_over = True
                break
        
        if not can_jump_over:
            return False
        else:
            return self.back_trace(nums, p0+1)
    
    def canJump(self, nums: List[int]) -> bool:
        if len(nums) == 1:
            return True
        
        # blockers are the zeros in nums. See if we can jump over all of them
        return self.back_trace(nums, 0)
   
if __name__ == "__main__":
    sln1 = Solution_1();
    sln2 = Solution_2();
    nums = [2,3,1,1,4]
    res1 = sln1.canJump(nums)
    res2 = sln2.canJump(nums)
