# https://leetcode.com/problems/rotate-image/

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