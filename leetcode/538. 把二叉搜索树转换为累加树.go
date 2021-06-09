package main

var (
	SUM = 0
)

func convertBST(root *TreeNode) *TreeNode {
	SUM = 0
	convertBTHelper(root)
	return root
}

// 中序遍历的时候，先遍历右子树再遍历根节点再遍历左子树，就可以使得所有比根节点大的节点在根节点之前被访问到
func convertBTHelper(root *TreeNode) {
	if root == nil {
		return
	}
	convertBTHelper(root.Right)
	SUM += root.Val
	root.Val = SUM
	convertBTHelper(root.Left)
}
