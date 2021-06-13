package main

import "math"

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lh := 1
	for left := root.Left; left != nil; left = left.Left {
		lh++
	}

	rh := 1
	for right := root.Right; right != nil; right = right.Right {
		rh++
	}

	if lh == rh {
		count := math.Pow(2, float64(rh))
		return int(count) - 1
	}

	return 1 + countNodes(root.Left) + countNodes(root.Right)
}
