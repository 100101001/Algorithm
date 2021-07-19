package main


// 0-1 背包问题
func KBackpack(w int, val []int) int {

	var dp = fillTwoArrayInt(len(val), w+1, 0)

	// 前i个物品，放入重量为w的背包
	for i := 0; i < len(val); i++ {
		for curW := 0; curW <= w; curW++ {
			if curW-val[i] < 0 {
				dp[i][curW] = 0
			} else if i-1 < 0 {
				dp[i][curW] = val[i]
			} else {
				// 前i-1个物品放curW的最大值 | 前i-1放 curW-val[i]的最大值加上val[i]
				dp[i][curW] = maxIntV2(dp[i-1][curW-val[i]]+val[i], dp[i-1][curW])
			}
		}
	}

	return dp[len(val)-1][w]
}
