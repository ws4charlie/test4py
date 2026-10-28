// Package p0036 solves LeetCode 36. Valid Sudoku (Medium).
// https://leetcode.com/problems/valid-sudoku/
//
// Decide whether a partly filled 9x9 Sudoku board is valid. Only the filled
// cells are checked:
//   - each row has no repeated digit 1-9,
//   - each column has no repeated digit 1-9,
//   - each of the nine 3x3 boxes has no repeated digit 1-9.
//
// The board doesn't have to be solvable.
//
// Example 1:
//
//	Input:  the board below
//	Output: true
//
//	    5 3 . | . 7 . | . . .
//	    6 . . | 1 9 5 | . . .
//	    . 9 8 | . . . | . 6 .
//	    ------+-------+------
//	    8 . . | . 6 . | . . 3
//	    4 . . | 8 . 3 | . . 1
//	    7 . . | . 2 . | . . 6
//	    ------+-------+------
//	    . 6 . | . . . | 2 8 .
//	    . . . | 4 1 9 | . . 5
//	    . . . | . 8 . | . 7 9
//
// Example 2:
//
//	Input:  the same board, but with the top-left 5 changed to 8
//	Output: false
//	Explanation: the top-left box (and column 0) now has two 8s.
//
// Constraints:
//
//	board.length == 9, board[i].length == 9
//	board[i][j] is a digit '1'-'9' or '.'.
package p0036

// IsValidSudoku records each digit seen per row, column and 3x3 box.
func IsValidSudoku(board [][]byte) bool {
	var rows, cols, boxes [9][9]bool
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] == '.' {
				continue
			}
			d := board[r][c] - '1'
			b := r/3*3 + c/3
			if rows[r][d] || cols[c][d] || boxes[b][d] {
				return false
			}
			rows[r][d], cols[c][d], boxes[b][d] = true, true, true
		}
	}
	return true
}
