
class Solution_Topdown:
    def minJump(self, nums, i, mem):
        if i in mem:
            return mem[i]
        if i >= len(nums):
            return 0
        
        minjump = -1 # -1 means can't jump to end from position i
        for j in range(1, nums[i]+1):
            jump = self.minJump(nums, i+j, mem)
            if jump != -1:
                if minjump == -1: # if not yet initialized
                    minjump = jump+1
                else:
                    minjump = min(minjump, jump+1)
        
        mem[i] = minjump
        return minjump
    
    # f(i) = min{f(j)} + 1, i < j <= i+nums[i]
    def jump(self, nums: List[int]) -> int:
        mem = {}
        mem[len(nums)-1] = 0
        
        return self.minJump(nums, 0, mem)
   
if __name__ == "__main__":
    sln = Solution_Topdown()
    nums = [2,3,1,1,4]
    res = sln.minJump(nums)
