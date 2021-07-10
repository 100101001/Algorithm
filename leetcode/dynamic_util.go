package main

func fillTwoArray(m, n, val int) [][]int {
	var dp = make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = val
		}
	}
	return dp
}

func fillOneArray(m, val int) []int {
	var dp = make([]int, m)
	for i := range dp {
		dp[i] = val
	}
	return dp
}
