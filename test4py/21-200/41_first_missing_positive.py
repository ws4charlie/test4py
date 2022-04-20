# https://leetcode.com/problems/first-missing-positive/

class Solution_1:
    def firstMissingPositive(self, nums: List[int]) -> int:
        if not nums:
            return 1
        
        nums.sort()
        
        res = 1
        for v in nums:
            if v < res:
                continue
            elif v == res:
                res += 1
            else:
                break
        return res

class Solution_2:
    def firstMissingPositive(self, nums: List[int]) -> int:
        if not nums:
            return 1
        
        heapq.heapify(nums)
        
        res = 1
        while nums:
            v = heapq.heappop(nums)
            if v < res:
                continue
            elif v == res:
                res += 1
            else:
                break
        return res

if __name__ == "__main__":
    sln = Solution()
    nums = [0,2,2,1,1]
    res = sln.firstMissingPositive(nums)
    print(res)
