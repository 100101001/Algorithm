package main

var ComposeRES [][]int

func combine(n, k int) [][]int {
	ComposeRES = [][]int{}
	combineHelper(n, k, 1, []int{})
	return ComposeRES
}

// 获取组合
func combineHelper(n, k, start int, path []int) {
	// 终止
	if len(path) == k {
		ComposeRES = append(ComposeRES, path)
	}

	// 选择列表
	for i := start; i <= n; i++ {
		// 做选择
		path = append(path, i)
		combineHelper(n, k, i+1, path)
		// 撤销选择
		path = path[0 : len(path)-1 : len(path)-1]
	}
}
