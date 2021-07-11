package main

func searchRange(nums []int, target int) []int {

	if len(nums) == 0 {
		return []int{-1, -1}
	}

	left, right := 0, len(nums) // 左闭右开
	for left < right {          // [left, left) => 区间为空
		mid := left + (right-left)/2
		if nums[mid] == target {
			right = mid
		} else if nums[mid] > target { // right 左移
			right = mid
		} else if nums[mid] < target { // left 右移
			left = mid + 1
		}
	}

	leftI := left
	// target 比所有数都要大
	if leftI >= len(nums) || nums[leftI] != target {
		leftI = -1
	}

	left, right = 0, len(nums) // 左闭右开
	for left < right {         // [left, left) => 区间为空
		mid := left + (right-left)/2
		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] > target { // right 左移
			right = mid
		} else if nums[mid] < target { // left 右移
			left = mid + 1
		}
	}

	rightI := left - 1
	// target比所有数都要小
	if rightI < 0 || nums[rightI] != target {
		rightI = -1
	}

	return []int{leftI, rightI}
}
