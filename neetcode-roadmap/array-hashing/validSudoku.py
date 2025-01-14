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

        # 3*3 checking
        d2 = defaultdict(bool)
        for r in range(9):
            for c in range(9):
                if board[r][c] == ".":
                    continue
                pos = (r // 3, c // 3)
                num = board[r][c]
                key = (pos, num)
                if key in d2:
                    return False 
                else:
                    d2[key] = True

        return True

