// https://leetcode.com/problems/rotate-image/
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
