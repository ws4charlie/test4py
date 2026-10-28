"""
36. Valid Sudoku  (Medium)
https://leetcode.com/problems/valid-sudoku/

Decide whether a partly filled 9x9 Sudoku board is valid. Only the filled
cells are checked:
    - each row has no repeated digit 1-9,
    - each column has no repeated digit 1-9,
    - each of the nine 3x3 boxes has no repeated digit 1-9.
The board doesn't have to be solvable.

Example 1:
    Input:  the board below
    Output: True

        5 3 . | . 7 . | . . .
        6 . . | 1 9 5 | . . .
        . 9 8 | . . . | . 6 .
        ------+-------+------
        8 . . | . 6 . | . . 3
        4 . . | 8 . 3 | . . 1
        7 . . | . 2 . | . . 6
        ------+-------+------
        . 6 . | . . . | 2 8 .
        . . . | 4 1 9 | . . 5
        . . . | . 8 . | . 7 9

Example 2:
    Input:  the same board, but with the top-left 5 changed to 8
    Output: False
    Explanation: the top-left box (and column 0) now has two 8s.

Constraints:
    board.length == 9, board[i].length == 9
    board[i][j] is a digit '1'-'9' or '.'.
"""

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