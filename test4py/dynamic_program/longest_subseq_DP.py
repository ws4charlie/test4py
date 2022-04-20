
# 最长子序列
# longest increasing sub-sequence
# i is a position, f(i) is length of longest sub-inc-seq starting from i
# f(i) = max{f(p) + 1}, where nums[p] >= f(i)
class Solution1_Top_Bottom:
    def back_trace(self, nums, i, mem):
        if i in mem:
            return mem[i]
        if i == len(nums) - 1:
            return 1

        # next positions to start with
        g_max = 1
        for j in range(i+1, len(nums)):
            if nums[j] >= nums[i]:
                c_max = self.back_trace(nums, j, mem)
                g_max = max(g_max, c_max+1)

        # keep result in mem and return it
        mem[i] = g_max
        return g_max

    def longestIncSeq(self, nums: 'List[int]') -> int:
        if not nums:
            return 0

        mem = {}
        res = self.back_trace(nums, 0, mem)
        return res

# longest increasing sub-sequence
# i is a position, f(i) is length of longest sub-inc-seq ending at i
# f(i) = max{f(p) + 1}, where nums[p] <= f(i)
class Solution2:
    def back_trace(self, nums, i):
        if i == 0:
            return 1

        # next positions to end with
        g_max = 1
        for j in range(i-1, -1, -1):
            if nums[j] < nums[i]:
                c_max = self.back_trace(nums, j)
                g_max = max(g_max, c_max + 1)
        return g_max

    def longestIncSeq(self, nums: 'List[int]') -> int:
        if not nums:
            return 0

        res = self.back_trace(nums, len(nums)-1)
        return res

if __name__ == "__main__":
    sln1 = Solution1()
    sln2 = Solution2()
    sln3 = Solution1_memo()
    nums = [1,5,3,4,6,9,7,8]
    res1 = sln1.longestIncSeq(nums)
    res2 = sln2.longestIncSeq(nums)
    res3 = sln3.longestIncSeq(nums)
    print(res)