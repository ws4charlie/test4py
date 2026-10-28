// https://leetcode.com/problems/excel-sheet-column-title/
package p0168

// ConvertToTitle is base 26 with digits 1..26 instead of 0..25, so shift
// down by one before each % and /.
func ConvertToTitle(n int) string {
	var buf []byte
	for n > 0 {
		n--
		buf = append(buf, byte('A'+n%26))
		n /= 26
	}
	for l, r := 0, len(buf)-1; l < r; l, r = l+1, r-1 {
		buf[l], buf[r] = buf[r], buf[l]
	}
	return string(buf)
}
