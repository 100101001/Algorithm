package main

func deleteNode(root *TreeNode, key int) *TreeNode {
	return deleteNodeHelper(root, key)
}

func deleteNodeHelper(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		if root.Left == nil {
			// 情况一：左子为空
			return root.Right
		} else if root.Right == nil {
			// 情况一：右子为空
			return root.Left
		} else {
			// 两个节点都不为空
			// 右子树最小或者左子树最大来替换自己
			minInRightSon := getMin(root.Right)
			root.Val = minInRightSon.Val
			root.Right = deleteNodeHelper(root.Right, minInRightSon.Val)
		}
	} else if root.Val > key {
		root.Right = deleteNodeHelper(root.Right, key)
	} else if root.Val < key {
		root.Left = deleteNodeHelper(root.Left, key)
	}
	return root
}

func getMin(root *TreeNode) *TreeNode {
	var final *TreeNode
	for r := root; r.Left != nil; r = r.Left {
		final = r.Left
	}
	return final
}
