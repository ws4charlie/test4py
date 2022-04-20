
# https://leetcode.com/problems/spiral-matrix/

class Solution:
    def back_trace(self, res, matrix, i, j, m, n):
        if m == 0 or n == 0:
            return
        if m == 1:
            for col in range(j, j+n):
                res.append(matrix[i][col])
            return
        if n == 1:
            for row in range(i, i+m):
                res.append(matrix[row][j])
            return
        
        for col in range(j, j+n-1):
            res.append(matrix[i][col])
        for row in range(i, i+m-1):
            res.append(matrix[row][j+n-1])
        for col in range(j+n-1, j, -1):
            res.append(matrix[i+m-1][col])
        for row in range(i+m-1, i, -1):
            res.append(matrix[row][j])
        
        self.back_trace(res, matrix, i+1, j+1, m-2, n-2)
    
    def spiralOrder(self, matrix: List[List[int]]) -> List[int]:
        if not matrix:
            return []
        
        res = []
        m = len(matrix)
        n = len(matrix[0])
        self.back_trace(res, matrix, 0, 0, m, n)
        return res

if __name__ == "__main__":
    sln = Solution()
    matrix1 = [[ 1, 2, 3 ],
               [ 4, 5, 6 ],
               [ 7, 8, 9 ]]
    matrix2 = [[1, 2, 3, 4],
              [5, 6, 7, 8],
              [9,10,11,12]]
    res1 = sln.spiralOrder(matrix1)
    res2 = sln.spiralOrder(matrix2)
