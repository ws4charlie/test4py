// Package p0054 solves LeetCode 54. Spiral Matrix (Medium).
// https://leetcode.com/problems/spiral-matrix/
//
// Return all elements of an m x n matrix in spiral order: clockwise from the
// top-left corner, moving inward.
//
// Example 1:
//
//	Input:  matrix = [[1,2,3],[4,5,6],[7,8,9]]
//	Output: [1,2,3,6,9,8,7,4,5]
//
//	    1 -> 2 -> 3
//	              |
//	              v
//	    4 -> 5    6
//	    ^         |
//	    |         v
//	    7 <- 8 <- 9
//
// Example 2:
//
//	Input:  matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
//	Output: [1,2,3,4,8,12,11,10,9,5,6,7]
//
//	     1 ->  2 ->  3 ->  4
//	                       |
//	                       v
//	     5 ->  6 ->  7     8
//	     ^                 |
//	     |                 v
//	     9 <- 10 <- 11 <- 12
//
// Constraints:
//
//	m == matrix.length, n == matrix[i].length
//	1 <= m, n <= 10
//	-100 <= matrix[i][j] <= 100
package p0054

// SpiralOrder walks the outer ring and shrinks the four bounds after each
// side. The checks before the bottom and left sides stop a single
// remaining row or column from being read twice.
func SpiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	top, bottom := 0, len(matrix)-1
	left, right := 0, len(matrix[0])-1
	res := make([]int, 0, len(matrix)*len(matrix[0]))
	for top <= bottom && left <= right {
		for c := left; c <= right; c++ {
			res = append(res, matrix[top][c])
		}
		top++
		for r := top; r <= bottom; r++ {
			res = append(res, matrix[r][right])
		}
		right--
		if top <= bottom {
			for c := right; c >= left; c-- {
				res = append(res, matrix[bottom][c])
			}
			bottom--
		}
		if left <= right {
			for r := bottom; r >= top; r-- {
				res = append(res, matrix[r][left])
			}
			left++
		}
	}
	return res
}
