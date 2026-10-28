"""
15. 3Sum  (Medium)
https://leetcode.com/problems/3sum/

Return every distinct triplet [nums[i], nums[j], nums[k]] with i, j, k
all different and nums[i] + nums[j] + nums[k] == 0. The result must not
contain the same triplet twice; order doesn't matter.

Example 1:
    Input:  nums = [-1,0,1,2,-1,-4]
    Output: [[-1,-1,2],[-1,0,1]]
    Explanation: -1 + 0 + 1 = 0 and -1 + -1 + 2 = 0. The triplet
                 [-1,0,1] can be formed twice but appears once.

Example 2:
    Input:  nums = [0,1,1]
    Output: []

Example 3:
    Input:  nums = [0,0,0]
    Output: [[0,0,0]]

Constraints:
    3 <= nums.length <= 3000
    -10^5 <= nums[i] <= 10^5
"""

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