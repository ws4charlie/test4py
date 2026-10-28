// https://leetcode.com/problems/contains-duplicate/
package p0217

// ContainsDuplicate uses a set (map to empty struct).
func ContainsDuplicate(nums []int) bool {
	seen := make(map[int]struct{}, len(nums))
	for _, v := range nums {
		if _, ok := seen[v]; ok {
			return true
		}
		seen[v] = struct{}{}
	}
	return false
}
