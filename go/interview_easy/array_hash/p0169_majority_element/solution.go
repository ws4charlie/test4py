// https://leetcode.com/problems/majority-element/
package p0169

// MajorityElement is the Boyer-Moore vote: the majority survives pairing
// off against every other value. O(n) time, O(1) space.
func MajorityElement(nums []int) int {
	cand, count := 0, 0
	for _, v := range nums {
		if count == 0 {
			cand = v
		}
		if v == cand {
			count++
		} else {
			count--
		}
	}
	return cand
}
