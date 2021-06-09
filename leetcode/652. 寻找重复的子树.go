package main

import "fmt"

var (
	res   []*TreeNode
	dedup map[string]int
)

// 找到重复子树
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	res = make([]*TreeNode, 0)
	dedup = map[string]int{}
	findDupSubTreeHelper(root)
	return res
}

// 1. 需要知道自己为根的树长什么样，因此是后序框架
func findDupSubTreeHelper(root *TreeNode) string {
	if root == nil {
		return "#"
	}

	left := findDupSubTreeHelper(root.Left)
	right := findDupSubTreeHelper(root.Right)

	treeStr := fmt.Sprintf("%s+%s-%d", left, right, root.Val)
	dedup[treeStr]++
	if dedup[treeStr] == 2 {
		res = append(res, root)
	}
	return treeStr
}
