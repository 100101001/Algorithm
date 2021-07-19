package main

// 有序数组，考虑双指针
func twoSum(numbers []int, target int) []int {
	return twoSumLR(numbers, target)
}

func twoSumLR(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1

	for left < right {
		cur := numbers[left] + numbers[right]
		if cur == target {
			return []int{left + 1, right + 1}
		} else if cur < target {
			left++
		} else if cur > target {
			right--
		}
	}

	return []int{-1, -1}
}
