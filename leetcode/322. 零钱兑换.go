package main

import "math"

func coinChange(coins []int, amount int) int {
	//return coinChangeDP(coins, amount)
	return coinChangeDFS(coins, amount)
}

// 时间复杂度： O(len(coins) ^ amount* len(coins))
// 选择硬币：+1
// 状态：凑的金额
// base case: 凑的金额 == 0
func coinChangeDP(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	// 初始化备忘录最大值
	var dp = make([]int, amount+1)
	for i := 0; i <= amount; i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0

	// 自底向上填充备忘录
	for i := 1; i <= amount; i++ {
		// 最优子结构，求最小值, 1+ min(sub1, sub2, sub3) 就等于循环最小值
		for _, coin := range coins {
			if i-coin >= 0 {
				dp[i] = minInt(dp[i], dp[i-coin]+1)
			}
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}

	return dp[amount]
}

var coinChangeMem []int

func coinChangeDFS(coins []int, amount int) int {
	coinChangeMem = make([]int, amount+1)
	for i := 0; i <= amount; i++ {
		coinChangeMem[i] = -1
	}
	if res := coinChangeHelper(coins, amount); res == math.MaxInt16 {
		return -1
	} else {
		return res
	}
}

func coinChangeHelper(coins []int, amount int) int {
	var res = math.MaxInt16
	if amount < 0 {
		return res
	}

	if coinChangeMem[amount] >= 0 {
		return coinChangeMem[amount]
	}

	if amount == 0 {
		return 0
	}

	for _, v := range coins {
		sub := coinChangeHelper(coins, amount-v)
		res = minInt(res, sub+1)
	}
	coinChangeMem[amount] = res
	return res
}
