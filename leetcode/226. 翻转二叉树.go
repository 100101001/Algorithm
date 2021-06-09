package main

func invertTree(root *TreeNode) *TreeNode {
	invertHelper(root)
	return root
}

// 对换root的左右子树
func invertHelper(root *TreeNode) {
	if root == nil {
		return
	}

	tmp := root.Right
	root.Right = root.Left
	root.Left = tmp

	invertHelper(root.Left)
	invertHelper(root.Right)
	return
}
