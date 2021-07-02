package main

import "math"

var maxSUM = 0

func maxSumBST(root *TreeNode) int {
	maxSUM = math.MinInt64
	if root == nil {
		return maxSUM
	}
	maxSumBSTHelper(root)
	return maxSUM
}

type maxSumReturn struct {
	Valid bool
	Min   int
	Max   int
	Sum   int
}

// 题意表述：是一颗二叉搜索树，且和最大
// 后序遍历：返回以root中的最大二叉搜索树
// 函数定义获取子树的最大键值和
// 后序遍历前：
// 1. 递归子树的终止条件：遇到空节点
// 后序遍历后：
// 1. 是否是二叉搜索树： 左子树是，右子树是，且判断当前根，和左子树中的最大值，以及右子树的最小值是否满足约束
// 2. 如果是二叉搜索树：左子树和+根+右子树和，迭代子树的最大，最小值
func maxSumBSTHelper(root *TreeNode) *maxSumReturn {
	ret := &maxSumReturn{}

	if root == nil {
		ret.Valid = true
		ret.Max = math.MinInt64  // (左)子树最大值
		ret.Min = math.MaxInt64 // (右)子树最小值
		ret.Sum = 0
		return ret
	}

	left := maxSumBSTHelper(root.Left)
	right := maxSumBSTHelper(root.Right)

	rootValid := root.Val < right.Min && root.Val > left.Max
	ret.Valid = left.Valid && right.Valid && rootValid

	if ret.Valid {
		ret.Sum = root.Val + right.Sum + left.Sum
		if ret.Sum > maxSUM {
			maxSUM = ret.Sum
		}
		ret.Min = minInt(root.Val, left.Min) // 满足
		ret.Max = maxInt(root.Val, right.Max)
	}

	return ret
}