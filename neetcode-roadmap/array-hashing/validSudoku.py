class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        d = defaultdict(bool)
       
       # row checking
        for _, inner_list in enumerate(board):
            d.clear()
            for _, n in enumerate(inner_list):
                if n.isdigit():
                    if n in d:
                        return False
                    else:
                        d[n] = True

        # column checking
        rowsize = len(board[0])
        for i in range(0, rowsize):
            d.clear()
            for row in board:
                if row[i].isdigit():
                    if row[i] in d:
                        return False
                    else:
                        d[row[i]] = True

