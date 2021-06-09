package main

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	return insertBSTHelper(root, val)
}

func insertBSTHelper(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	//if root.Val == val {
	//} else
	if val < root.Val {
		root.Left = insertBSTHelper(root.Left, val)
	} else {
		root.Right = insertBSTHelper(root.Right, val)
	}
	return root
}
