#Given an array of integers, return indices of the two numbers such that they add up to a specific target.

#You may assume that each input would have exactly one solution, and you may not use the same element twice.

#Example:
#Given nums = [2, 7, 11, 15], target = 9,

#Because nums[0] + nums[1] = 2 + 7 = 9,
#return [0, 1].

# https://leetcode.com/problems/two-sum/
# fullstack
# backend
class Solution:
    def twoSum(nums, target):
        res = []#list
        seen = {}#map
        
        #for (int i = 0; i < nums.size(); i++) {
        #    int v = nums[i];
        #
        #}
        #for (int v :nums)
        #  {cout << v;}
        #for v in nums:
        #  print(v)
        for i, v in enumerate(nums):
            other = target - v
            if other in seen:#seen.find(other) != seen.end()
                res = [seen[other], i]
                break
            else:
                seen[v] = i
        return res

if __name__ == '__main__':
    sln = Solution()
    ary = [2, 7, 11, 15]
    tar = 9
    res = sln.twoSum(ary, tar)
    print(res)
