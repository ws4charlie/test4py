
# https://leetcode.com/problems/3sum/

class Solution:
    def threeSum(self, nums) -> list:
        if len(nums) < 3:
            return []
        
        nums.sort()
        hs = set()
        for i in range(len(nums) - 2):
            vi = nums[i]
            if vi > 0:
                break
                
            j = i + 1
            k = len(nums) - 1            
            while j < k:
                vj = nums[j]
                vk = nums[k]
                sum = vi + vj + vk
                
                if sum < 0:
                    j += 1
                elif sum > 0:
                    k -= 1
                else:
                    j += 1
                    k -= 1
                    hs.add((vi, vj, vk))
        
        return list(hs)

if __name__ == "__main__":
    sln = Solution()
    nums = [-1,0,1,2,-1,-4]
    res = sln.threeSum(nums)
    print(res)