import itertools


class Solution:
    def isValidSudoku(self, board: list[list[str]]) -> bool:

        def is_unit_valid(unit: list[str]) -> bool:
            unit = [i for i in unit if i != '.']
            return len(set(unit)) == len(unit)

        def is_square_valid(board: list[list[str]]) -> bool:
            for i, j in itertools.product(range(0, 9, 3), range(0, 9, 3)):
                square = [board[x][y] for x in range(i, i + 3) for y in range(j, j + 3)]
                if not is_unit_valid(square):
                    return False
            return True

        def is_row_valid(board: list[list[str]]) -> bool:
            return all(is_unit_valid(row) for row in board)

        def is_col_valid(board: list[list[str]]) -> bool:
            return all(is_unit_valid(list(col)) for col in zip(*board))

        return is_row_valid(board) and is_col_valid(board) and is_square_valid(board)
