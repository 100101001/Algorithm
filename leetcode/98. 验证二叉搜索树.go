package main

func isValidBST(root *TreeNode) bool {
	return isValidBSTHelper(root, nil, nil)
}

//      5
//   1    6
//      3   7  这棵树告知不能简单通过比较左右子树根和树根自身，而是应该遵循定义
// root 必须处于 lower和upper限定的开区间内
func isValidBSTHelper(root, lower, upper *TreeNode) bool {
	if root == nil {
		return true
	}
	// 根节点大于左子树最大值
	if lower != nil && root.Val <= lower.Val {
		return false
	}
	// 根节点大于左子树最大值
	if upper != nil && root.Val >= upper.Val {
		return false
	}
	// 父亲决定了左子树的开区间的upper, 右子树的开区间的lower
	return isValidBSTHelper(root.Left, lower, root) &&
		isValidBSTHelper(root.Right, root, upper)
}
