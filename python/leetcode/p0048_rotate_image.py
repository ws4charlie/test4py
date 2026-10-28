"""
48. Rotate Image  (Medium)
https://leetcode.com/problems/rotate-image/

Rotate an n x n matrix 90 degrees clockwise, in place: modify matrix
directly instead of building a new one.

Example 1:
    Input:  matrix = [[1,2,3],[4,5,6],[7,8,9]]
    Output: [[7,4,1],[8,5,2],[9,6,3]]

        1 2 3         7 4 1
        4 5 6   -->   8 5 2
        7 8 9         9 6 3

Example 2:
    Input:  matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
    Output: [[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]

         5  1  9 11         15 13  2  5
         2  4  8 10   -->   14  3  4  1
        13  3  6  7         12  6  8  9
        15 14 12 16         16  7 10 11

Constraints:
    n == matrix.length == matrix[i].length
    1 <= n <= 20
    -1000 <= matrix[i][j] <= 1000
"""

from typing import List

class Solution_inplace:
    def backtrace(self, matrix, i, j, n):
        if n <= 1:
            return
        
        for k in range(n-1):
            temp = matrix[i][j+k]
            matrix[i][j+k] = matrix[i+n-1-k][j]
            matrix[i+n-1-k][j] = matrix[i+n-1][j+n-1-k]
            matrix[i+n-1][j+n-1-k] = matrix[i+k][j+n-1]
            matrix[i+k][j+n-1] = temp
        
        self.backtrace(matrix, i+1, j+1, n-2)
        
        
    def rotate(self, matrix: 'List[List[int]]') -> None:
        """
        Do not return anything, modify matrix in-place instead.
        """
        
        if len(matrix) <= 1:
            return
        
        n = len(matrix)
        self.backtrace(matrix, 0, 0, n)

class Solution_space_N:     
    def rotate(self, matrix: List[List[int]]) -> None:
        """
        Do not return anything, modify matrix in-place instead.
        """
        
        if len(matrix) <= 1:
            return
        
        n = len(matrix)
        res = [[] for i in range(n)]
        for i in range(n):
            row = res[i]
            for j in range(n, 0, -1):
                row.append(matrix[j-1][i])
                
        matrix[:] = res[:]

if __name__ == "__main__":
    matrix = [
        [1,2,3],
        [4,5,6],
        [7,8,9]]
    sln = Solution_inplace()
    sln.rotate(matrix)
    print(matrix)