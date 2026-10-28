// https://leetcode.com/problems/two-sum/
package p0001

// TwoSum returns the indices of the two numbers that add up to target.
func TwoSum(nums []int, target int) []int {
	seen := make(map[int]int) // value -> index
	for i, v := range nums {
		if j, ok := seen[target-v]; ok {
			return []int{j, i}
		}
		seen[v] = i
	}
	return nil
}
