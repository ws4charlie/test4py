"""
121. Best Time to Buy and Sell Stock  (Easy)
https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

prices[i] is a stock's price on day i. Buy on one day and sell on a later
day. Return the maximum profit, or 0 if no trade makes money.

Example 1:
    Input:  prices = [7,1,5,3,6,4]
    Output: 5
    Explanation: buy on day 1 (price 1), sell on day 4 (price 6).
                 Selling before buying isn't allowed.

Example 2:
    Input:  prices = [7,6,4,3,1]
    Output: 0
    Explanation: prices only fall, so don't trade.

Constraints:
    1 <= prices.length <= 10^5
    0 <= prices[i] <= 10^4
"""

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
