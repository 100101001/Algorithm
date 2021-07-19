package main

func fillTwoArrayInt(m, n, val int) [][]int {
	var dp = make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = val
		}
	}
	return dp
}

func fillTwoArrayBool(m, n int, val bool) [][]bool {
	var dp = make([][]bool, m)
	for i := range dp {
		dp[i] = make([]bool, n)
		for j := range dp[i] {
			dp[i][j] = val
		}
	}
	return dp
}

func fillOneArrayInt(m, val int) []int {
	var dp = make([]int, m)
	for i := range dp {
		dp[i] = val
	}
	return dp
}
