// https://leetcode.com/problems/single-number/
package p0136

// SingleNumber XORs everything: x^x == 0 and x^0 == x, so pairs cancel.
func SingleNumber(nums []int) int {
	res := 0
	for _, v := range nums {
		res ^= v
	}
	return res
}
