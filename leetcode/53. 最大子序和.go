package main

import "math"

func maxSubArray(nums []int) int {
	return maxSubArrayDP(nums)
}

// dp[i] 的定义是以i结尾最大子串，注意和子序列的区别导致无需遍历j<i去算
func maxSubArrayDP(nums []int) int {

	var dp = make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		// 自己OR接上子串
		dp[i] = maxIntV2(dp[i-1]+nums[i], nums[i])
	}

	// 返回结果
	var max = math.MinInt64
	for i := range dp {
		if max < dp[i] {
			max = dp[i]
		}
	}
	return max
}
