package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 后序遍历，
// 函数定义：拉直子树
// 将左子变为右子，将右子变为左子最后一个被遍历到的节点的孩子（通过后序遍历拉直，此时就是只找右子即可）
func flatten(root *TreeNode) *TreeNode {
	helper(root)
	return root
}

// 将以n为根的树拉成右直链
func helper(n *TreeNode) {
	if n == nil {
		return
	}
	// 将左子树为根的树拉成右直链，将右子树为根的树拉成右直链
	helper(n.Left)
	helper(n.Right)

	//
	left := n.Left
	right := n.Right

	n.Left = nil
	n.Right = left

	var r = n
	for ; r.Right != nil; r = r.Right {
	}
	r.Right = right
}
