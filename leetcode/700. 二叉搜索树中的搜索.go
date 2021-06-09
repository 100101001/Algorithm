package main

func isInBST(root *TreeNode, target int) bool {
	return isInBSTHelper(root, target)
}

// 判断是否在BST
func isInBSTHelper(root *TreeNode, target int) bool {
	if root == nil {
		return false
	}
	if target == root.Val {
		return true
	}
	if target < root.Val {
		return isInBSTHelper(root.Left, target)
	} else {
		return isInBSTHelper(root.Right, target)
	}
}
