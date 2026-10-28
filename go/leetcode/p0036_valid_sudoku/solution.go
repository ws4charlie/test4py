// https://leetcode.com/problems/valid-sudoku/
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
