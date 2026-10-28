# https://leetcode.com/problems/valid-sudoku/

class Solution:
    def isValidSudoku(self, board) -> bool:
        row_sets = [set() for i in range(9)]
        col_sets = [set() for i in range(9)]
        cub_sets = [set() for i in range(9)]
        
        for row in range(9):
            for col in range(9):
                cub = self.cub_index(row, col)
                cell = board[row][col]
                if cell != '.':
                    if cell not in row_sets[row]:
                        row_sets[row].add(cell)
                    else:
                        return False
                    if cell not in col_sets[col]:
                        col_sets[col].add(cell)
                    else:
                        return False
                    if cell not in cub_sets[cub]:
                        cub_sets[cub].add(cell)
                    else:
                        return False
        return True
        
    @staticmethod
    def cub_index(row, col):
        index = row // 3 * 3 + col // 3
        return index


board = [["8","3",".",".","7",".",".",".","."],
         ["6",".",".","1","9","5",".",".","."],
         [".","9","8",".",".",".",".","6","."],
         ["8",".",".",".","6",".",".",".","3"],
         ["4",".",".","8",".","3",".",".","1"],
         ["7",".",".",".","2",".",".",".","6"],
         [".","6",".",".",".",".","2","8","."],
         [".",".",".","4","1","9",".",".","5"],
         [".",".",".",".","8",".",".","7","9"]]

if __name__ == '__main__':
    sln = Solution()
    valid = sln.isValidSudoku(board)