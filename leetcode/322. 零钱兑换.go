package main

func coinChange(coins []int, amount int) int {

	// dp[i]的意思是凑成i最少需要多少类coins
	dp := make([]int, 0, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}

	//
	dp[0] = 0
	for i := 0; i < amount+1; i++ {
		for _, c := range coins {
			if i-c >= 0 {
				if dp[i-c]+1 < dp[i] {
					dp[i] = dp[i-c] + 1
				}
			}
		}
	}

	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}
