
class Solution:
    def bs(self, nums, i, j, target) -> int:
        mid = (i+j)//2
        if mid == i:
            if nums[i] == target:
                return i
            if nums[j] == target:
                return j
            return -1
        
        pv = nums[0]
        mv = nums[mid]
        if target >= pv:
            if mv >= pv:
                if target <= mv:
                    return self.bs(nums, i, mid, target)
                else:
                    return self.bs(nums, mid, j, target)
            else:
                return self.bs(nums, i, mid, target)
        else:
            if mv < pv:
                if target <= mv:
                    return self.bs(nums, i, mid, target)
                else:
                    return self.bs(nums, mid, j, target)
            else:
                return self.bs(nums, mid, j, target);
        
    def search(self, nums, target: int) -> int:
        if len(nums) == 0:
            return -1
        if len(nums) == 1:
            if nums[0] == target:
                return 0
            else:
                return -1
            
        index = self.bs(nums, 0, len(nums)-1, target)
        return index

if __name__ == "__main__":
    nums = [4,5,6,7,0,1,2]
    target = 0
    sln = Solution()
    res = sln.search(nums, target)
