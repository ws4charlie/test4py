// https://leetcode.com/problems/valid-palindrome/
package p0125

// IsPalindrome skips non-alphanumeric bytes from both ends and compares
// the rest case-insensitively (the input is ASCII).
func IsPalindrome(s string) bool {
	for l, r := 0, len(s)-1; l < r; {
		switch {
		case !isAlnum(s[l]):
			l++
		case !isAlnum(s[r]):
			r--
		case lower(s[l]) != lower(s[r]):
			return false
		default:
			l, r = l+1, r-1
		}
	}
	return true
}

func isAlnum(c byte) bool {
	return 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || '0' <= c && c <= '9'
}

func lower(c byte) byte {
	if 'A' <= c && c <= 'Z' {
		return c + 'a' - 'A'
	}
	return c
}
