package main

// 分割等和子集
func canPartition(nums []int) bool {
	return canPartitionDP(nums)
}

// 0-1 背包问题框架：
// 1：状态和选择
// 2：定义dp和base case
// 3：状态转移
func canPartitionDP(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum&1 == 1 {
		return false
	}

	target := sum / 2

	// 前i个物品能否挑出物品刚好装满w
	var dp = fillTwoArrayBool(len(nums), target+1, false)

	for i := 0; i < len(nums); i++ {
		for curW := 0; curW <= target; curW++ {
			if curW == 0 {
				dp[i][curW] = true
			} else if i-1 < 0 {
				// 只有一个物品就只能看本物品能否塞满
				dp[i][curW] = curW == nums[i]
			} else if curW < nums[i] {
				dp[i][curW] = dp[i-1][curW]
			} else {
				// 放或不放的选择下：两者只要有一种ok就ok
				dp[i][curW] = dp[i-1][curW-nums[i]] || dp[i-1][curW]
			}
		}
	}

	return dp[len(nums)-1][target]
}
