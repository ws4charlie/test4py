// Package secondlargest finds the second largest value, ported from
// python/basics/find_second_largest.py.
package secondlargest

// Find returns the second largest value in one pass, counting duplicates
// ([5, 5] -> 5). ok is false when there are fewer than two values.
func Find(nums []int) (second int, ok bool) {
	if len(nums) < 2 {
		return 0, false
	}
	first, second := max(nums[0], nums[1]), min(nums[0], nums[1])
	for _, v := range nums[2:] {
		if v > first {
			first, second = v, first
		} else if v > second {
			second = v
		}
	}
	return second, true
}
