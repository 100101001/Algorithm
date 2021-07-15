package main

import "fmt"

//
func twoSum(numbers []int, target int) []int {
	return twoSumLeftRight(numbers, target)
}

func twoSumLeftRight(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1

	for left < right {
		cur := numbers[left] + numbers[right]
		if numbers[left]+numbers[right] == target {

		}
		fmt.Println(cur)
	}

	return nil
}
