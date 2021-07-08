package main

var AllPath [][]int

func permute(nums []int) [][]int {
	AllPath = [][]int{}
	backtrack(nums, []int{})
	return AllPath
}

// 复杂度:N!
// 路径：path，
// 选择列表通过传入nums可做判断
// 回溯条件：路径长度等于排列原数据长度
func backtrack(nums []int, path []int) {
	if len(path) == len(nums) {
		AllPath = append(AllPath, path)
		return
	}

	for i := 0; i < len(nums); i++ {
		if permuteContains(nums[i], path) {
			continue
		}
		path = append(path, nums[i])
		backtrack(nums, path)
		path = path[0 : len(path)-1 : len(path)-1]
	}
}

func permuteContains(i int, path []int) bool {
	for _, v := range path {
		if v == i {
			return true
		}
	}
	return false
}
