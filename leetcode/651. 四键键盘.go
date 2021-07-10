package main

// 基于的信息是最优问题结尾一定是拷贝
func maxA(n int) int {
	return maxADP(n)
}

// 状态只有1维
func maxADP(n int) int {
	// n次操作的最大值
	var dp = make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		var pressA = dp[i-1] + 1
		dp[i] = pressA
		for j := 2; j < i; j++ {
			var copied = dp[j-2]
			var copiedTimes = i - j
			dp[i] = maxInt(dp[i], copied+copied*copiedTimes)
		}
	}
	return dp[n]
}
