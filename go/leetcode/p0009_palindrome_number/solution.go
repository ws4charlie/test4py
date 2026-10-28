// https://leetcode.com/problems/palindrome-number/
package p0009

// IsPalindrome reverses only the second half of the digits, without
// converting to a string. Negatives and numbers ending in 0 (except 0)
// can't be palindromes.
func IsPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	half := 0
	for x > half {
		half = half*10 + x%10
		x /= 10
	}
	return x == half || x == half/10 // half/10 drops the middle digit
}
