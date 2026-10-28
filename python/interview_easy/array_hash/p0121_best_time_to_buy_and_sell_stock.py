# https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

class Solution:
    def maxProfit(self, prices):
        # one pass: track the lowest price so far, sell today if better
        low, best = float('inf'), 0
        for p in prices:
            low = min(low, p)
            best = max(best, p - low)
        return best

if __name__ == "__main__":
    sln = Solution()
    assert sln.maxProfit([7, 1, 5, 3, 6, 4]) == 5
    assert sln.maxProfit([7, 6, 4, 3, 1]) == 0
    assert sln.maxProfit([2]) == 0
    print("ok")
