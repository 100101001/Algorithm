package main

import "strings"

var AllPaths [][]string

// 回溯本质是暴力搜索
func solveNQueens(n int) [][]string {

	var board [][]string
	for i := 0; i < n; i++ {
		var row []string
		for j := 0; j < n; j++ {
			row = append(row, ".")
		}
		board = append(board, row)
	}
	solveNQueensHelper(0, board)
	return AllPaths
}

func solveNQueensHelper(row int, board [][]string) {
	if len(board) == row {
		var res []string
		for _, row := range board {
			res = append(res, strings.Join(row, ""))
		}
		AllPaths = append(AllPaths, res)
		return
	}

	n := len(board)
	for col := 0; col < n; col++ {
		// 剪枝
		if !isNQueenValid(row, col, board) {
			continue
		}
		board[row][col] = "Q"
		solveNQueensHelper(row+1, board)
		board[row][col] = "."
	}
}

// 路径：board 中小于 row 的那些行都已经成功放置了皇后
// 选择列表：第 row 行的所有列都是放置皇后的选择
// 结束条件：row 超过 board 的最后一行
func isNQueenValid(row, col int, board [][]string) bool {
	for i := row - 1; i >= 0; i-- {
		if board[i][col] == "Q" {
			return false
		}
	}
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == "Q" {
			return false
		}
	}
	for i, j := row-1, col+1; i >= 0 && j < len(board); i, j = i-1, j+1 {
		if board[i][j] == "Q" {
			return false
		}
	}
	return true
}

func isNQueenChoiceValid(i int, path []int) bool {
	row := len(path)
	for j := row - 1; j >= 0; j-- {
		pos := path[j]
		if pos == i {
			return false
		}
		rowDiff := nQAbs(row, j)
		colDiff := nQAbs(pos, i)
		if rowDiff == colDiff {
			return false
		}
	}
	return true
}

func nQAbs(i, j int) int {
	if i > j {
		return i - j
	}
	return j - i
}

func nQInt2str(intP, n int) string {
	var res string
	for i := 0; i < n; i++ {
		if i != intP {
			res += "."
		} else {
			res += "Q"
		}
	}
	return res
}
