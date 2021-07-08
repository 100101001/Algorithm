package main

var RES [][]int

func subsets(nums []int) [][]int {
	RES = [][]int{}
	subsetsHelper(nums, 0, []int{})
	return RES
}

// 暴力穷举：回溯
func subsetsHelper(nums []int, start int, path []int) {
	RES = append(RES, path)
	for i := start; i < len(nums); i++ {
		path = append(path, nums[i])
		subsetsHelper(nums, i+1, path)
		path = path[0 : len(path)-1 : len(path)-1]
	}
}

func subsets2(nums []int) [][]int {
	return subsets2Helper(len(nums)-1, nums, []int{})
}

// 归纳法: 后序
func subsets2Helper(n int, nums []int, sets []int) [][]int {
	if n == -1 {
		return [][]int{{}}
	}

	subsets := subsets2Helper(n-1, nums, sets)

	var res [][]int
	for _, subset := range subsets {
		var subsetN []int
		for _, v := range subset {
			subsetN = append(subsetN, v)
		}
		res = append(res, append(subsetN, nums[n]))
		res = append(res, subset)
	}
	return res
}
