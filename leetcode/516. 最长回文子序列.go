package main

// dp[i][j] s[i..j]的最大回文子序列长度
func longestPalindromeSubSeq(s string) int {

	// 备忘录空字符串=0, 单个字符=1
	var dp = fillTwoArray(len(s), len(s), 0)
	for i := 0; i < len(dp); i++ {
		dp[i][i] = 1
	}

	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				// 两字符不相等则不可能同时出现在 s[i..j]的最长回文子序列中
				dp[i][j] = maxIntV2(dp[i+1][j], dp[i][j-1])
			}
		}
	}

	return dp[0][len(s)-1]
}
