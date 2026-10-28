// Package p0017 solves LeetCode 17. Letter Combinations of a Phone Number (Medium).
// https://leetcode.com/problems/letter-combinations-of-a-phone-number/
//
// Given a string of digits 2-9, return every letter string they could spell
// on a phone keypad, in any order. Digit 1 maps to no letters.
//
//	+-------+-------+-------+
//	|   1   |   2   |   3   |
//	|       |  abc  |  def  |
//	+-------+-------+-------+
//	|   4   |   5   |   6   |
//	|  ghi  |  jkl  |  mno  |
//	+-------+-------+-------+
//	|   7   |   8   |   9   |
//	| pqrs  |  tuv  | wxyz  |
//	+-------+-------+-------+
//	|   *   |   0   |   #   |
//	+-------+-------+-------+
//
// Example 1:
//
//	Input:  digits = "23"
//	Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
//
// Example 2:
//
//	Input:  digits = "2"
//	Output: ["a","b","c"]
//
// Constraints:
//
//	digits.length <= 4
//	digits[i] is a digit in '2'-'9'.
package p0017

var letters = map[byte]string{
	'2': "abc", '3': "def", '4': "ghi", '5': "jkl",
	'6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz",
}

// LetterCombinations backtracks over digits, reusing one byte buffer.
func LetterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}
	var res []string
	comb := make([]byte, len(digits))
	var backtrack func(i int)
	backtrack = func(i int) {
		if i == len(digits) {
			res = append(res, string(comb))
			return
		}
		for _, c := range []byte(letters[digits[i]]) {
			comb[i] = c
			backtrack(i + 1)
		}
	}
	backtrack(0)
	return res
}

// LetterCombinationsIterative grows the result one digit at a time (BFS).
func LetterCombinationsIterative(digits string) []string {
	if digits == "" {
		return nil
	}
	res := []string{""}
	for i := 0; i < len(digits); i++ {
		var next []string
		for _, prefix := range res {
			for _, c := range letters[digits[i]] {
				next = append(next, prefix+string(c))
			}
		}
		res = next
	}
	return res
}
