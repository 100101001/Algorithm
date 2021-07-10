package main

func longestCommonSubsequence(text1 string, text2 string) int {
	return longestCommonSubsequenceDP(text1, text2)
}

var LCSMem [][]int

// 最长公共子串『自顶向下』
func longestCommonSubsequenceDFS(text1 string, text2 string) int {
	LCSMem = make([][]int, len(text1))
	for i := range LCSMem {
		LCSMem[i] = make([]int, len(text2))
		for j := range LCSMem[i] {
			LCSMem[i][j] = -1
		}
	}

	return longestCommonSubsequenceDFSHelper(text1, text2, 0, 0)
}

// s1[i...] 与 s2[j...] 的最长公共子串
// s1[i] =?= s2[j] 进行分类讨论
func longestCommonSubsequenceDFSHelper(text1, text2 string, i, j int) int {
	if i == len(text1) || j == len(text2) {
		return 0
	}

	if LCSMem[i][j] >= 0 {
		return LCSMem[i][j]
	}

	if text1[i] == text2[j] {
		LCSMem[i][j] = longestCommonSubsequenceDFSHelper(text1, text2, i+1, j+1) + 1
	} else {
		LCSMem[i][j] = maxInt(
			// 意思是
			longestCommonSubsequenceDFSHelper(text1, text2, i, j+1),
			longestCommonSubsequenceDFSHelper(text1, text2, i+1, j))
	}
	return LCSMem[i][j]
}

// 最长公共子串『自底向上』
func longestCommonSubsequenceDP(text1, text2 string) int {
	var dp = make([][]int, len(text1))
	for i := range dp {
		dp[i] = make([]int, len(text2))
	}

	for i := len(text1) - 1; i >= 0; i-- {
		for j := len(text2) - 1; j >= 0; j-- {
			if text1[i] == text2[j] {
				var dp1 = 0
				if i+1 < len(text1) && j+1 < len(text2) {
					dp1 = dp[i+1][j+1]
				}
				dp[i][j] = dp1 + 1
			} else {
				var (
					dp1 = 0
					dp2 = 0
				)
				if i+1 < len(text1) {
					dp1 = dp[i+1][j]
				}
				if j+1 < len(text2) {
					dp2 = dp[i][j+1]
				}
				dp[i][j] = maxInt(dp1, dp2)
			}
		}
	}
	return dp[0][0]
}
