// https://leetcode.com/problems/roman-to-integer/
package p0013

var value = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

// RomanToInt subtracts a symbol when a larger one follows it (IV, CM, ...).
func RomanToInt(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		v := value[s[i]]
		if i+1 < len(s) && v < value[s[i+1]] {
			res -= v
		} else {
			res += v
		}
	}
	return res
}
