package main

import "math"

func minFallingPathSum(matrix [][]int) int {
	return minFallingPathSumDFS(matrix)
}

var minFallingPathSumDFSMEM [][]int

func minFallingPathSumDFS(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])

	// 初始化备忘录
	minFallingPathSumDFSMEM = fillTwoArray(m, n, math.MaxInt64)

	// 最小值是遍历最后一行所有终点的最小值
	var resInt = math.MaxInt64
	for i := 0; i < n; i++ {
		resInt = minIntV2(minFallingPathSumDFSHelper(matrix, m-1, i), resInt)
	}

	return resInt
}

func minFallingPathSumDFSHelper(matrix [][]int, i, j int) int {

	if i < 0 {
		return 0
	}
	if j < 0 || j >= len(matrix) {
		return math.MaxInt64
	}

	if minFallingPathSumDFSMEM[i][j] != math.MaxInt64 {
		return minFallingPathSumDFSMEM[i][j]
	}

	minFallingPathSumDFSMEM[i][j] = minIntV2(
		minFallingPathSumDFSHelper(matrix, i-1, j),
		minFallingPathSumDFSHelper(matrix, i-1, j-1),
		minFallingPathSumDFSHelper(matrix, i-1, j+1),
	) + matrix[i][j]

	return minFallingPathSumDFSMEM[i][j]
}

//
func minFallingPathSumDP(matrix [][]int) int {
	var dp = fillTwoArray(len(matrix), len(matrix[0]), -1)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			var (
				dp1 int
				dp2 int
				dp3 int
			)
			if i-1 >= 0 {
				if j-1 < 0 {
					dp1 = math.MaxInt64
				} else {
					dp1 = dp[i-1][j-1]
				}
				if j+1 >= len(matrix[0]) {
					dp3 = math.MaxInt64
				} else {
					dp3 = dp[i-1][j+1]
				}
				dp2 = dp[i-1][j]
			}
			dp[i][j] = minIntV2(dp1, dp2, dp3) + matrix[i][j]
		}
	}

	var resMin = math.MaxInt64
	for _, v := range dp[len(matrix)-1] {
		resMin = minIntV2(resMin, v)
	}
	return resMin
}
