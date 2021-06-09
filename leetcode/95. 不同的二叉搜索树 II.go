package main

func generateTrees(n int) []*TreeNode {
	return generateTreesHelper(1, n)
}

// 生成以lo,hi为区间的树
func generateTreesHelper(lo, hi int) []*TreeNode {
	var res []*TreeNode
	if lo > hi {
		res = append(res, nil)
		return res
	}

	for i := lo; i <= hi; i++ {
		lefts := generateTreesHelper(lo, i-1)
		rights := generateTreesHelper(i+1, hi)
		for _, left := range lefts {
			for _, right := range rights {
				root := &TreeNode{Val: i}
				root.Left = left
				root.Right = right
				res = append(res, root)
			}
		}
	}
	return res
}
