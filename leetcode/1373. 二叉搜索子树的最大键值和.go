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

func maxSumBSTHelper(root *TreeNode) *maxSumReturn {
	ret := &maxSumReturn{}

	if root == nil {
		ret.Valid = true
		ret.Max = math.MinInt64
		ret.Min = math.MaxInt64
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
		ret.Min = minInt(root.Val, left.Min)
		ret.Max = maxInt(root.Val, right.Max)
	}

	return ret
}
