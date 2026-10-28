// Package p0048 solves LeetCode 48. Rotate Image (Medium).
// https://leetcode.com/problems/rotate-image/
//
// Rotate an n x n matrix 90 degrees clockwise, in place: modify matrix
// directly instead of building a new one.
//
// Example 1:
//
//	Input:  matrix = [[1,2,3],[4,5,6],[7,8,9]]
//	Output: [[7,4,1],[8,5,2],[9,6,3]]
//
//	    1 2 3         7 4 1
//	    4 5 6   -->   8 5 2
//	    7 8 9         9 6 3
//
// Example 2:
//
//	Input:  matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
//	Output: [[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]
//
//	     5  1  9 11         15 13  2  5
//	     2  4  8 10   -->   14  3  4  1
//	    13  3  6  7         12  6  8  9
//	    15 14 12 16         16  7 10 11
//
// Constraints:
//
//	n == matrix.length == matrix[i].length
//	1 <= n <= 20
//	-1000 <= matrix[i][j] <= 1000
package p0048

// Rotate turns the matrix 90° clockwise in place: transpose, then reverse
// each row.
func Rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	for _, row := range matrix {
		for l, r := 0, n-1; l < r; l, r = l+1, r-1 {
			row[l], row[r] = row[r], row[l]
		}
	}
}

// RotateLayers rotates ring by ring, moving four cells at a time
// (left -> top -> right -> bottom -> left).
func RotateLayers(matrix [][]int) {
	n := len(matrix)
	for layer := 0; layer < n/2; layer++ {
		first, last := layer, n-1-layer
		for i := first; i < last; i++ {
			k := i - first
			top := matrix[first][i]
			matrix[first][i] = matrix[last-k][first]
			matrix[last-k][first] = matrix[last][last-k]
			matrix[last][last-k] = matrix[i][last]
			matrix[i][last] = top
		}
	}
}
