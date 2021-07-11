package main

// 最长递增子序列
func lengthOfLIS(nums []int) int {
	return lengthOfLISDP(nums)
}

// dp[i] 是指以 nums[i]结尾的递增子序列最大长度
func lengthOfLISDP(nums []int) int {

	var dp = make([]int, len(nums))

	// 填充 dp
	for i := 0; i < len(nums); i++ {
		dp[i] = 1 // 注意备忘录初始值是1
		for j := i - 1; j >= 0; j-- {
			if nums[j] < nums[i] {
				dp[i] = maxIntV2(dp[i], dp[j]+1)
			}
		}
	}

	// 获取最大值
	var maxI = 1
	for _, v := range dp {
		if maxI < v {
			maxI = v
		}
	}
	return maxI
}
