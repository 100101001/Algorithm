package main

// 翻转一棵二叉树。
// 先序遍历：
// 先序遍历终止条件：root为空
// 先序遍历前：对换左右孩子
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
