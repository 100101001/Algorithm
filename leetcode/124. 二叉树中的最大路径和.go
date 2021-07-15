package main

import "math"

var RET int

func maxPathSum(root *TreeNode) int {
	RET = math.MinInt64
	maxPathSumHelper(root)
	return RET
}

// 递归函数：包含当前根的最大向下的路径和
func maxPathSumHelper(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxIntV2(maxPathSumHelper(root.Left), 0)
	right := maxIntV2(maxPathSumHelper(root.Right), 0)
	ret := root.Val + maxIntV2(left, right)
	if root.Val+left+right > RET {
		RET = root.Val + left + right
	}
	return ret
}
