package main

// 解数独
func solveSudoku(board [][]byte) {
	solveSudokuHelper(board, 0, 0)
}

// 下一个?：先移动列，再移动行
// 选择列表：1~9
// 路径：棋盘
// 返回值：是否结束不再
func solveSudokuHelper(board [][]byte, i, j int) bool {

	if j == len(board[0]) {
		return solveSudokuHelper(board, i+1, 0)
	}

	if i == len(board) {
		return true
	}

	// 如果当前格子非空则直接跳过，枚举下一个格子
	if board[i][j] != byte('.') {
		return solveSudokuHelper(board, i, j+1)
	}

	for ch := '1'; ch <= '9'; ch++ {
		// 做选择
		if !isSudokuValid(board, i, j, byte(ch)) {
			continue
		}
		board[i][j] = byte(ch)
		if solveSudokuHelper(board, i, j+1) {
			return true
		}
		board[i][j] = byte('.')
	}
	return false
}

func isSudokuValid(board [][]byte, i, j int, choice byte) bool {

	//for t := 0; t < 9; t++ {
	//	if board[i][t] == choice || board[t][j] == choice {
	//		return false
	//	}
	//	if board[(i/3)*3+t/3][(j/3)*3+t%3] == choice {
	//		return false
	//	}
	//}

	// 检查每行（该列没一样的）
	for ii := 0; ii < len(board); ii++ {
		if board[ii][j] == choice {
			return false
		}
	}

	// 检查该行每列（该行没一样的）
	for jj := 0; jj < len(board[0]); jj++ {
		if board[i][jj] == choice {
			return false
		}
	}

	// 得到中心格子
	r := (i/3)*3 + 1 //0,1,2 | 3,4,5 | 6,7,8
	c := (j/3)*3 + 1
	// 相对中心格子进行移动，判断九宫格是否有一样的
	moves := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 0}, {0, 1}, {1, -1}, {1, 0}, {1, 1},
	}
	for _, move := range moves {
		rr := r + move[0]
		cc := c + move[1]
		if board[rr][cc] == choice {
			return false
		}
	}

	return true
}
