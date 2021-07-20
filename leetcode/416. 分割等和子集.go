package main

// 分割等和子集
func canPartition(nums []int) bool {
	return canPartitionDP(nums)
}

// 0-1 背包问题框架：
// 1：状态和选择
// 2：定义dp和base case
// 3：状态转移
// 4: 状态压缩：倒着压不然前面的小w会被覆盖
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
	var dp = fillOneArrayBool(target+1, true)
	dp[0] = true

	for i := 0; i < len(nums); i++ {
		for curW := target; curW >= 1; curW-- {
			if curW >= nums[i] {
				// 放或不放的选择下：两者只要有一种ok就ok
				dp[curW] = dp[curW-nums[i]] || dp[curW]
			}
		}
	}

	return dp[target]
}
