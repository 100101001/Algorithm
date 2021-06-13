package main

func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}
	col := len(board[0])
	row := len(board)

	// 找到四边的O以及DFS找到连通用#标记
	// 首尾列
	for i := 0; i < row; i++ {
		if board[i][0] == 'O' {
			dfsSolve(board, i, 0)
		}
		if board[i][col-1] == 'O' {
			dfsSolve(board, i, col-1)
		}
	}
	// 首尾行
	for i := 0; i < col; i++ {
		if board[0][i] == 'O' {
			dfsSolve(board, 0, i)
		}
		if board[row-1][i] == 'O' {
			dfsSolve(board, row-1, i)
		}
	}

	// 剩余O标记为X，#标记为O
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == '#' {
				board[i][j] = 'O'
			}
		}
	}
}

func dfsSolve(board [][]byte, i int, j int) {
	col := len(board[0])
	row := len(board)

	if i > row-1 || i < 0 || j > col-1 || j < 0 {
		return
	}
	if board[i][j] != 'O' {
		return
	}

	// 与边上的O连通就标记#
	board[i][j] = '#'

	d := [][]int{
		{1, 0}, {-1, 0}, {0, 1}, {0, -1},
	}
	for _, v := range d {
		dfsSolve(board, i+v[0], j+v[1])
	}
}

func solve1(board [][]byte) {
	if len(board) == 0 {
		return
	}

	row := len(board)
	col := len(board[0])

	// 初始化 union - find
	uf := newUF(row*col + 1)
	dummy := row * col

	//
	// 首尾行
	for i := 0; i < col; i++ {
		if board[0][i] == 'O' {
			uf.union(0*col+i, dummy)
		}
		if board[row-1][i] == 'O' {
			uf.union((row-1)*col+i, dummy)
		}
	}
	// 首尾列
	for i := 0; i < row; i++ {
		if board[i][0] == 'O' {
			uf.union(i*col+0, dummy)
		}
		if board[i][col-1] == 'O' {
			uf.union(i*col+col-1, dummy)
		}
	}

	//
	d := [][]int{
		{1, 0}, {0, 1}, {-1, 0}, {0, -1},
	}

	for i := 1; i < row-1; i++ {
		for j := 1; j < col-1; j++ {
			// 自身是O
			if board[i][j] == 'O' {
				uf.union(i*col+j, dummy)
			}
			// 上下左右连通
			for _, vv := range d {
				x := i + vv[0]
				y := j + vv[1]
				// x,y不会超过
				if board[x][y] == 'O' {
					uf.union(x*col+y, i*col+j)
				}
			}
		}
	}

	// 所有不和 dummy 连通的 O，都要被替换
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if board[i][j] == 'O' && !uf.connected(i*col+j, dummy) {
				board[i][j] = 'X'
			}
		}
	}
}
