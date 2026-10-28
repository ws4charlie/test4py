// https://leetcode.com/problems/missing-number/
package p0268

// MissingNumber XORs every index 0..n with every value; all but the
// missing number pair up and cancel.
func MissingNumber(nums []int) int {
	res := len(nums)
	for i, v := range nums {
		res ^= i ^ v
	}
	return res
}

// MissingNumberSum subtracts the actual sum from 0+1+...+n.
func MissingNumberSum(nums []int) int {
	n := len(nums)
	sum := n * (n + 1) / 2
	for _, v := range nums {
		sum -= v
	}
	return sum
}
