package main

import "sort"

var TARGET int

func canPartitionKSubsets(nums []int, k int) bool {
	if k > len(nums) {
		return false
	}
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%k != 0 {
		return false
	}

	TARGET = sum / k

	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	return canPartitionKSubsetsBT(nums, 0, make([]int, k, k))
}

// 路径：bucket
// 选择：把nums的第i个元素放到哪个bucket中
// 回溯条件：i == len(nums)
// 递归穷举nums中每个数字放到第x个桶中
// 尽量命中剪枝条件可以优化时间：大的元素放前面，可以尽早剪枝
func canPartitionKSubsetsBT(nums []int, idx int, bucket []int) bool {
	if idx == len(nums) {
		for _, v := range bucket {
			if v != TARGET {
				return false
			}
		}
		return true
	}

	for i := 0; i < len(bucket); i++ {
		// 剪枝
		if !validChoice(bucket, i, nums[idx]) {
			continue
		}
		bucket[i] += nums[idx]
		if canPartitionKSubsetsBT(nums, idx+1, bucket) {
			return true
		}
		bucket[i] -= nums[idx]
	}
	return false
}

func validChoice(bucket []int, choice, num int) bool {
	return bucket[choice]+num <= TARGET
}

func canPartitionKSubsets2(nums []int, k int) bool {
	if k > len(nums) {
		return false
	}
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%k != 0 {
		return false
	}

	TARGET = sum / k

	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	return canPartitionKSubsetsBT2(k, 0, nums, make([]bool, len(nums), len(nums)))
}


// 路径是
func canPartitionKSubsetsBT2(k, bucket int, nums []int, used []bool) bool {
	if k == 0 {
		return true
	}

	// 装满第k个桶了，看装第k-1个桶
	if bucket == TARGET {
		return canPartitionKSubsetsBT2(k-1, 0, nums, used)
	}

	for i := 0; i < len(nums); i++ {
		// 剪枝
		if used[i] || !canPartitionKSubsetsBT2Valid(nums[i], bucket) {
			continue
		}
		used[i] = true
		if canPartitionKSubsetsBT2(k, bucket+nums[i], nums, used) {
			return true
		}
		used[i] = false
	}
	return false
}

func canPartitionKSubsetsBT2Valid(n, bucket int) bool {
	return bucket+n <= TARGET
}
