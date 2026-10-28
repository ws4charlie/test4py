// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
package p0121

import "math"

// MaxProfit tracks the lowest price so far and the best sale today.
func MaxProfit(prices []int) int {
	low, best := math.MaxInt, 0
	for _, p := range prices {
		low = min(low, p)
		best = max(best, p-low)
	}
	return best
}
