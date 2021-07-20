package main

func change(amount int, coins []int) int {
	return changeDPCompress(amount, coins)
}

// 1： status,choice: 前i个硬币，amount；选择就是第i个币是否选择
// 2: base case&dp：dp[i][j], 前i个硬币能凑出j的组合数量；dp[..][0] = 1, dp[0][..] = 0 除非j==0
// 3: transform：
// 如果用了第i个硬币，组合数就等于 dp[i][j-coin[i-1]]（前i个硬币凑用完后的前的组合数）；
// 如果不用第i个硬币，组合数就等于 dp[i-1][j] （前i-1个硬币凑j的组合数）
// 4: compress
func changeDP(amount int, coins []int) int {
	var dp = fillTwoArrayInt(len(coins)+1, amount+1, 0)
	for i := 0; i <= len(coins); i++ {
		dp[i][0] = 1
	}

	for i := 1; i <= len(coins); i++ {
		for j := 1; j <= amount; j++ {
			dp[i][j] = dp[i-1][j]
			if left := j - coins[i-1]; left >= 0 {
				dp[i][j] += dp[i][left]
			}
		}
	}

	return dp[len(coins)][amount]
}

func changeDPCompress(amount int, coins []int) int {
	var dp = fillOneArrayInt(amount+1, 0)
	dp[0] = 1

	for i := 1; i <= len(coins); i++ {
		for j := 1; j <= amount; j++ {
			if left := j - coins[i-1]; left >= 0 {
				dp[j] += dp[left]
			}
		}
	}
	return dp[amount]
}
