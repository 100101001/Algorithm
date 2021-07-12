package main

// 最长递增子序列
func lengthOfLIS(nums []int) int {
	return lengthOfLIPBinarySearchLeft(nums)
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

// patience sort，top内是最长递增子序列，二分左边界，piles是有效长度变量
func lengthOfLIPBinarySearchLeft(nums []int) int {
	top := make([]int, len(nums)+1)
	var piles = 0 // 记录了当前nums有效长度
	for i := 0; i < len(nums); i++ {
		//
		left, right := 0, piles
		for left < right {
			mid := left + (right-left)/2
			if top[mid] == nums[i] {
				right = mid
			} else if top[mid] < nums[i] {
				left = mid + 1
			} else if top[mid] > nums[i] {
				right = mid
			}
		}

		//
		if left == piles {
			piles++
		}
		// 更新数组内容
		top[left] = nums[i]
	}
	return piles
}
