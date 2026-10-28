
# https://leetcode.com/problems/container-with-most-water/

class Solution:
    def maxArea(self, height) -> int:
        l, r = 0, len(height) - 1
        
        res = 0
        while l < r:
            hl = height[l]
            hr = height[r]
            res = max(res, (r-l) * min(hl, hr))
            
            if hl < hr:
                l += 1
            else:
                r -= 1
        
        return res

if __name__ == "__main__":
    sln = Solution()
    height = [1,8,6,2,5,4,8,3,7]
    area = sln.maxArea(height)
    print(area)